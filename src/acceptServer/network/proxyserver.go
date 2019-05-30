package network

import (
	"acceptServer/amf"
	"acceptServer/gateway"
	"acceptServer/proto"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
)

func StartTransfer() {
	tcp_addr, err := net.ResolveTCPAddr("tcp", ":8911")
	if err != nil {
		log.Fatal(err)
	}
	l, err := net.ListenTCP("tcp", tcp_addr)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	log.Printf("[TransferTcpStart]tcp listening on %+v\n", tcp_addr)
	for {
		c, err := l.AcceptTCP()
		if err != nil {
			log.Printf("[TransferTcpStart] AcceptTCP error %v\n", err)
			continue
		}
		go handleRequest(c)
	}
}

func handleRequest(conn net.Conn) {
	fmt.Println(conn.RemoteAddr())

	var header proto.ClientPackageHeader
	err := binary.Read(conn, binary.BigEndian, &header)

	if err != nil {
		log.Printf("[getTargetIp] error, binary fail to read, %v\n", err)
		return
	}

	//body
	bs := make([]byte, header.Length-12)
	_, err = io.ReadFull(conn, bs)
	if err != nil {
		log.Printf("[getTargetIp] error, io fail to read body\n", err)
		_ = conn.Close()
		return
	}

	//encode amf message
	cli_buf := bytes.NewBuffer(bs)
	amf_buf, err := amf.Decode(cli_buf)
	if err != nil {
		log.Printf("[getTargetIp] fail to decode bs %v\n", err)
		return
	}

	log.Println("client Info", header, amf_buf)

	gateway.OnMessage(conn, amf_buf)
}
