package gateway

import (
	"acceptServer/amf"
	"acceptServer/proto"
	"bytes"
	"net"
)

func OnMessage(conn net.Conn, args interface{}) {
	//conn.Write()

	var header = proto.ClientPackageHeader{Length: 12, Flag: 1, Cmd: 2}

	var buffer bytes.Buffer
	count, err := amf.Encode(&buffer, header)
	if err != nil {
		println(err)
		_ = conn.Close()
		return
	}

	println(count)
	count, _ = conn.Write(buffer.Bytes())

	println("写入数据长度:", count)
}
