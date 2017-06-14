package main

import (
	"fmt"
	"net"
)

func Chat() {
	host := "127.0.0.1:5056"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", host)
	ChkErr(err, "tcpAddr")

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	ChkErr(err, "DialTCP")

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	ChkErr(err, "Read")
	fmt.Println(string(buf[:n]))
}
