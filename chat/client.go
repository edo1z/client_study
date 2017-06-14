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
		_, err := conn.Write(input)
		util.ChkErr(err, "sender write")
	}
}

func receiver(conn *net.TCPConn, name string) {
	for running {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		util.ChkErr(err, "Receiver read")
		fmt.Println(string(buf[:n]))
		fmt.Println(name + "> ")
	}

}

func Chat() {
	fmt.Print("Please input your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _, err := reader.ReadLine()

	host := "127.0.0.1:7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", host)
	util.ChkErr(err, "tcpAddr")

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	util.ChkErr(err, "DialTCP")
	//defer conn.Close()

	_, err = conn.Write(name)
	util.ChkErr(err, "Write name")

	go receiver(conn, string(name))
	go sender(conn, string(name))

	for running {
		time.Sleep(1 * 1e9)
	}
}
