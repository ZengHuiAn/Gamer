package framework

import (
	"github.com/pkg/errors"
	"sync"
)

var errLen = errors.New("管道可读大小为 0 ")

type Chanel struct {
	lock sync.Mutex  // 锁
	ch   chan []byte // 管道
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
		return nil, errLen
	}
	data = <-self.ch
	return data, nil
}
