package display

import (
	"context"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"os"
	"time"

	"github.com/MaxHalford/halfgone"
	"github.com/stianeikeland/go-rpio/v4"
	"golang.org/x/sys/unix"
)

const (
	EpdWidth       = 176
	EpdHeight      = 264
	deviceFilePath = "/dev/spidev0.0"
)

type UpdateChannel chan image.Image

func NewDisplayChannel(ctx context.Context) UpdateChannel {
	images := make(UpdateChannel)
	go func() {
		display := NewDisplay()
		defer display.Close()
		for {
			select {
			case displayImage := <-images:
				display.DisplayImage(displayImage)
			case <-ctx.Done():
				return
			}
		}
	}()

	return images
}

// Ensure rpio.Open() is called before using this
func NewDisplay() Display {
	display := Display{
		rstPin:  rpio.Pin(17),
		dcPin:   rpio.Pin(25),
		csPin:   rpio.Pin(8),
		busyPin: rpio.Pin(24),
	}
	display.rstPin.Output()
	display.dcPin.Output()
	display.csPin.Output()
	display.busyPin.Input()

	dtStart := time.Now()
	device, err := openDev()
	if err != nil {
		log.Fatalf("unable to open device: %#v", err)
	}
	display.reset()
	fmt.Printf("init device: %v\n", time.Since(dtStart))

	display.device = device

	dtStart = time.Now()
	display.sendCommand(PowerSetting)
	display.sendData(0x03) // VDS_EN, VDG_EN
	display.sendData(0x00) // VCOM_HV, VGHL_LV[1], VGHL_LV[0]
	display.sendData(0x2b) // VDH
	display.sendData(0x2b) // VDL
	display.sendData(0x09) // VDHR
	display.sendCommand(BoosterSoftStart)
	display.sendData(0x07)
	display.sendData(0x07)
	display.sendData(0x17)
	// Power optimization
	display.sendCommand(0xF8)
	display.sendData(0x60)
	display.sendData(0xA5)
	// Power optimization
	display.sendCommand(0xF8)
	display.sendData(0x89)
	display.sendData(0xA5)
	// Power optimization
	display.sendCommand(0xF8)
	display.sendData(0x90)
	display.sendData(0x00)
	// Power optimization
	display.sendCommand(0xF8)
	display.sendData(0x93)
	display.sendData(0x2A)
	// Power optimization
	display.sendCommand(0xF8)
	display.sendData(0xA0)
	display.sendData(0xA5)
	// Power optimization
	display.sendCommand(0xF8)
	display.sendData(0xA1)
	display.sendData(0x00)
	// Power optimization
	display.sendCommand(0xF8)
	display.sendData(0x73)
	display.sendData(0x41)
	display.sendCommand(PartialDisplayRefresh)
	display.sendData(0x00)
	display.sendCommand(PowerOn)

	display.waitUntilIdle()
	display.sendCommand(PanelSetting)
	display.sendData(0xAF) // KW-BF   KWR-AF    BWROTP 0f
	display.sendCommand(PllControl)
	display.sendData(0x3A) // 3A 100HZ   29 150Hz 39 200HZ    31 171HZ
	display.sendCommand(VcmDcSettingRegister)
	display.sendData(0x12)
	time.Sleep(2 * time.Millisecond)
	display.setLut()
	fmt.Printf("init device: %v\n", time.Since(dtStart))
	//  # EPD hardware init end
	return display
}

type Display struct {
	rstPin  rpio.Pin
	dcPin   rpio.Pin
	csPin   rpio.Pin
	busyPin rpio.Pin
	device  *os.File
}

func (d Display) DisplayImage(img image.Image) {
	grayImage := convertToGray(img)
	ditheredImage := halfgone.ThresholdDitherer{Threshold: 240}.Apply(grayImage)

	dtStart := time.Now()
	buf := newBuffer()
	writeImageToBuffer(ditheredImage, buf)
	fmt.Printf("writing to buffer: %v\n", time.Since(dtStart))

	dtStart = time.Now()
	d.displayFrame(buf)
	fmt.Printf("render to screen: %v\n", time.Since(dtStart))
}

func writeImageToBuffer(ditheredImage *image.Gray, buf []byte) {
	for y := 0; y < EpdHeight; y++ {
		for x := 0; x < EpdWidth; x++ {
			grayColor := ditheredImage.At(y, x).(color.Gray)
			// if grayColor.Y == 0 {
			if grayColor.Y > 0 {
				buf[(x+y*EpdWidth)/8] |= 0x80 >> (uint(x) % uint(8))
			}
		}
	}
}

func newBuffer() []byte {
	bufferLength := EpdWidth * EpdHeight / 8
	buf := make([]byte, bufferLength)
	for i := 0; i < bufferLength; i++ {
		buf[i] = 0x00
	}
	return buf
}

func openDev() (*os.File, error) {
	return os.OpenFile(deviceFilePath, unix.O_RDWR|unix.O_NOCTTY|unix.O_NONBLOCK, 0666)
}

func (d Display) reset() {
	d.rstPin.Low()
	time.Sleep(200 * time.Millisecond)
	d.rstPin.High()
	time.Sleep(200 * time.Millisecond)
}

func (d Display) sendCommand(b byte) {
	d.dcPin.Low()
	_, err := d.device.Write([]byte{b})
	if err != nil {
		log.Fatalf("failed to write command to device: %#v", err)
	}
}

func (d Display) sendData(b byte) {
	d.dcPin.High()
	_, err := d.device.Write([]byte{b})
	if err != nil {
		log.Fatalf("failed to write data to device: %#v", err)
	}
}

func (d Display) waitUntilIdle() {
	dtStart := time.Now()

	for {
		if d.busyPin.Read() == rpio.High {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}

	log.Printf("wait until idle: %v", time.Since(dtStart))
}

func (d Display) setLut() {
	d.sendCommand(LutForVcom) // vcom
	for count := 0; count < 44; count++ {
		d.sendData(lutVcomDc[count])
	}
	d.sendCommand(LutWhiteToWhite) // ww --
	for count := 0; count < 42; count++ {
		d.sendData(lutWw[count])
	}
	d.sendCommand(LutBlackToWhite) // bw r
	for count := 0; count < 42; count++ {
		d.sendData(lutBw[count])
	}
	d.sendCommand(LutWhiteToBlack) // wb w
	for count := 0; count < 42; count++ {
		d.sendData(lutBb[count])
	}
	d.sendCommand(LutBlackToBlack) // bb b
	for count := 0; count < 42; count++ {
		d.sendData(lutWb[count])
	}
}

func (d Display) displayFrame(b []byte) {
	dtStart := time.Now()
	size := len(b)

	d.sendCommand(DataStartTransmission1)

	time.Sleep(2 * time.Millisecond)
	log.Printf("send command 1: %v", time.Since(dtStart))
	dtStart = time.Now()

	for i := 0; i < size; i++ {
		d.sendData(0xFF)
	}
	time.Sleep(2 * time.Millisecond)
	log.Printf("send white image data: %v", time.Since(dtStart))

	dtStart = time.Now()
	d.sendCommand(DataStartTransmission2)
	time.Sleep(2 * time.Millisecond)
	log.Printf("send command 2: %v", time.Since(dtStart))

	dtStart = time.Now()
	for i := 0; i < size; i++ {
		d.sendData(b[i])
	}
	time.Sleep(2 * time.Millisecond)
	log.Printf("send image data: %v", time.Since(dtStart))

	dtStart = time.Now()
	d.sendCommand(Refresh)
	log.Printf("send display REFRESH: %v", time.Since(dtStart))

	d.waitUntilIdle()
}

func (d Display) Close() {
	err := d.device.Close()
	if err != nil {
		log.Printf("error closing SPI device: %v", err)
	}
}
