package chat

import (
	"../util"
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

var running = true

func sender(conn *net.TCPConn, name string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(name + "> ")
		input, _, _ := reader.ReadLine()
		if string(input) == "\\q" {
			_, _ = conn.Write([]byte("quit"))
			running = false
			break
		}
		_, _ = conn.Write([]byte(input))
	}
}

func receiver(conn *net.TCPConn) {

}

func Chat() {
	host := "127.0.0.1:5056"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", host)
	util.ChkErr(err, "tcpAddr")

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	util.ChkErr(err, "DialTCP")
	//defer conn.Close()

	fmt.Print("Please input your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _, err := reader.ReadLine()
	_, err = conn.Write(name[0 : len(name)-1])
	util.ChkErr(err, "Write name")

	go receiver(conn)
	go sender(conn, string(name))

	for running {
		time.Sleep(1 * 1e9)
	}
}
