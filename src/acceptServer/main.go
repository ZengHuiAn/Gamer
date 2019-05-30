package main

import (
	"acceptServer/network"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	network.StartTransfer()
	s := wait(os.Interrupt, os.Kill, syscall.SIGTERM)
	log.Printf("Got signal `%s`", s.String())
}

func wait(signals ...os.Signal) os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, signals...)
	s := <-c
	return s
}
