package remote

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"log"

	"github.com/Oppodelldog/toggleperfect/internal/display"
)

func startDisplay(display display.UpdateChannel, output chan Message) {
	go func() {
		for img := range display {
			imageMessage, ok := imageToMessage(img)
			if ok {
				output <- imageMessage
			}
		}
	}()
}

func imageToMessage(img image.Image) (Message, bool) {
	buf := bytes.NewBuffer([]byte{})
	err := png.Encode(buf, img)
	if err != nil {
		log.Printf("error encoding display image: %v", err)
		return Message{}, false
	}

	base64Image := base64.StdEncoding.EncodeToString(buf.Bytes())
	imageUrl := fmt.Sprintf("data:image/png;base64,%s", base64Image)

	return Message{
		Action: actionDisplay,
		Data:   imageUrl,
	}, true
}
