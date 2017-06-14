package chat

import (
	"../util"
	"bufio"
	"fmt"
	"net"
	"os"
)

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

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	util.ChkErr(err, "Read")
	fmt.Println(string(buf[:n]))
}
