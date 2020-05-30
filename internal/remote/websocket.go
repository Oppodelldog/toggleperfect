package remote

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Oppodelldog/toggleperfect/internal/log"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Action string `json:"action"`
	Data   string `json:"data"`
}

func fromJson(data []byte) (Message, error) {
	var v Message
	err := json.NewDecoder(bytes.NewBuffer(data)).Decode(&v)
	if err != nil {
		return Message{}, fmt.Errorf("error deconding Message: %v", err)
	}
	return v, nil
}

func toJson(msg Message) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(buffer).Encode(msg)
	if err != nil {
		return nil, fmt.Errorf("error encoding Message: %v", err)
	}
	return buffer.Bytes(), nil
}

func NewWebsocketEndpoint(input chan Message, output chan Message) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mb := MessageBuffer{}
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer c.Close()

		go func() {
			for {
				mt, message, err := c.ReadMessage()
				if err != nil {
					log.Print("error reading from client:", err)
					break
				}
				if mt != websocket.TextMessage {
					log.Print("unexpected message type:", err)
					break
				}
				mb.buffer(string(message))
				if mb.hasMessages() {
					msg, err := fromJson([]byte(mb.next()))
					if err != nil {
						log.Print("could decode client message:", err)
						continue
					}
					input <- msg
				}
			}
		}()

		for {
			select {
			case message, ok := <-output:
				if !ok {
					log.Print("output closed, websocket handler done")
					return
				}
				jsonBytes, err := toJson(message)
				if err != nil {
					log.Print("error encoding message:", err)
					break
				}
				err = c.WriteMessage(websocket.TextMessage, jsonBytes)
				if err != nil {
					log.Print("error writing to client:", err)
					return
				}
			}
		}
	}
}

type MessageBuffer struct {
	messageBuffer string
	messages      []string
}

func (b *MessageBuffer) buffer(data string) {
	b.messageBuffer += data
	var buffer = ""
	for i := 0; i < len(b.messageBuffer); i++ {
		var char = b.messageBuffer[i]
		buffer += string(char)
		if char == '\n' {
			b.messages = append(b.messages, buffer)
			buffer = ""
		}
	}
	b.messageBuffer = buffer
}

func (b MessageBuffer) hasMessages() bool {
	return len(b.messages) > 0
}

func (b *MessageBuffer) next() string {
	msg := b.messages[0]
	b.messages = b.messages[1:]
	return msg
}
