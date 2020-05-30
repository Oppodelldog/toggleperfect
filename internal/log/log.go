package log

import (
	"fmt"
	"log"
)

var receivers []chan string
var newReceiver chan chan string
var logdata chan string

func Init() {
	newReceiver = make(chan chan string)
	logdata = make(chan string)
	go func() {
		for {
			select {
			case msg := <-logdata:
				for _, r := range receivers {
					r <- msg
				}
			case newRecv := <-newReceiver:
				receivers = append(receivers, newRecv)
			}
		}
	}()
}

func AddReceiver(r chan string) {
	newReceiver <- r
}

func send(msg string) {
	if logdata != nil {
		logdata <- msg
	}
}

func Print(a ...interface{}) {
	msg := fmt.Sprint(a...)
	send(msg)

	log.Print(msg)
}

func Printf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	send(msg)

	log.Print(msg)
}

func Fatalf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a)
	send(msg)

	log.Fatal(msg)
}

func Fatal(a ...interface{}) {
	msg := fmt.Sprint(a...)
	send(msg)

	log.Fatal(msg)
}
