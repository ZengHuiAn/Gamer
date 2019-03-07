package Loger

import (
	"fmt"
	"io"
	"log"
)

func print(v ...interface{})  {
	log.Println("GOLANG:",v)
}

func AppendContent(w io.Writer, a ...interface{}) (n int, err error)  {
	//fmt.Fprint()
}