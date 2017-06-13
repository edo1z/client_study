package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	host := "127.0.0.1:7777"
	msg := "hello"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", host)
	chkErr(err, "tcpAddr")
	fmt.Println(tcpAddr)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	chkErr(err, "DialTCP")
	fmt.Println(conn)

	_, err = conn.Write([]byte(msg))
	chkErr(err, "Write")
	fmt.Println("Wrote")

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	chkErr(err, "Read")
	fmt.Println(string(buf[:n]))
}

func chkErr(err error, place string) {
	if err != nil {
		fmt.Printf("(%s)", place)
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(0)
	}
}
