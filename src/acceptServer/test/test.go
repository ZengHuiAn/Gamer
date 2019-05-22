package test

import (
	"acceptServer/amf"
	"acceptServer/proto"
	"bytes"
	"encoding/binary"
	"fmt"
)




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
	header := proto.ClientPackageHeader{Length: len, Flag: 1, Cmd: 1}
	err = binary.Write(&headerBf, binary.BigEndian, &header)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bf.Bytes())
	fmt.Println(headerBf.Bytes())

	fmt.Println(append(headerBf.Bytes(), bf.Bytes()...))
}
