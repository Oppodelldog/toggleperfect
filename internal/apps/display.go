package apps

import (
	"context"
	"gitlab.com/Oppodelldog/toggleperfect/internal/display"
	"image"
	"image/jpeg"
	"log"
	"os"
)

func NewDevDisplayChannel(ctx context.Context) display.UpdateChannel {
	images := make(display.UpdateChannel)
	go func() {
		defer close(images)
		for {
			select {
			case displayImage := <-images:
				saveImage(displayImage)
			case <-ctx.Done():
				return
			}
		}
	}()

	return images
}

func saveImage(displayImage image.Image) {
	opt := jpeg.Options{
		Quality: 90,
	}
	f, err := os.OpenFile("dev-display.jpg", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0655)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Printf("error closing dev display image: %f", err)
		}
	}()
	err = jpeg.Encode(f, displayImage, &opt)
	if err != nil {
		panic(err)
	}

}
