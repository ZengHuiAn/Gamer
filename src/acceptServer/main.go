package main

import (
	"acceptServer/network"
	"errors"
	"log"
	"os"
	"os/signal"
	"sync"
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

type Chanel struct {
	lock   sync.Mutex  // 锁
	ch     chan []byte // 管道
	status bool
}

func (self *Chanel) createChanel() {
	self.ch = make(chan []byte, 256)
}

func (self *Chanel) Lock() {
	self.lock.Lock()
}
func (self *Chanel) UnLock() {
	self.lock.Unlock()
}

func (self *Chanel) Append(data []byte) {
	self.Lock()
	defer self.UnLock()
	if self.ch == nil {
		self.createChanel()
	}
	self.ch <- data
}

func (self *Chanel) Read() (data []byte, err error) {
	self.Lock()
	defer self.UnLock()
	if len(self.ch) == 0 {
		err = errors.New("管道可读大小为 0 ")
		return nil, err
	}
	data = <-self.ch
	return data, nil
}
