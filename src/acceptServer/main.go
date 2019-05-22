package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/AnZenghui/goserver/src/acceptServer/amf"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type ClientPackageHeader struct {
	Length uint32
	Flag   uint32
	Cmd    uint32
}

func test() {
	arr := make([]interface{}, 9)
	arr[0] = 1
	arr[1] = nil
	arr[2] = true
	arr[3] = false
	arr[4] = 3.5
	arr[5] = "lskjdfksff"
	arr[6] = []byte{1, 2, 3, 4, 5}
	arr[7] = []interface{}{3, "xx"}
	arr[8] = []int{3, 4, 5}
	var bf bytes.Buffer
	_, err := amf.Encode(&bf, arr)
	if err != nil {
		fmt.Println(err)
		return
	}
	var headerBf bytes.Buffer
	len := uint32(12 + bf.Len())
	header := ClientPackageHeader{Length: len, Flag: 1, Cmd: 1}
	err = binary.Write(&headerBf, binary.BigEndian, &header)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bf.Bytes())
	fmt.Println(headerBf.Bytes())

	fmt.Println(append(headerBf.Bytes(), bf.Bytes()...))
}

func main() {
	//

	go test()

	s := wait(os.Interrupt, os.Kill, syscall.SIGTERM)
	log.Printf("Got signal `%s`", s.String())
	//chanel.Append(buffer.Bytes())
	//data, err := chanel.Read()
	//fmt.Println("管道读取数据", data, err)
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
