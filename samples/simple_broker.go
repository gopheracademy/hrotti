package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gopheracademy/hrotti/broker"
)

func main() {
	h := hrotti.NewHrotti(100)
	hrotti.INFO = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	h.AddListener("test", hrotti.NewListenerConfig("tcp://0.0.0.0:1883"))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	h.Stop()
}
