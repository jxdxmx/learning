package main

import (
	"bufio"
	"log"
	"net"
	"time"
)

func f202112128() {
	go tcp()
	go udp()
	time.Sleep(20000 * time.Second)
}

func udp() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{Port: 20000})
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	defer listen.Close()
	for {
		var buf [1024]byte
		n, addr, err := listen.ReadFromUDP(buf[:])
		if err != nil {
			log.Printf("read udp error: %v\n", err)
			continue
		}
		data := append([]byte("hello "), buf[:n]...)
		_, _ = listen.WriteToUDP(data, addr)
	}
}
func tcp() {
	listen, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept error: %v\n", err)
			continue
		}
		// 开始goroutine监听连接
		go handleConn(conn)
	}
}
func handleConn(conn net.Conn) {
	defer conn.Close()
	// 读写缓冲区
	rd := bufio.NewReader(conn)
	wr := bufio.NewWriter(conn)
	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			log.Printf("read error: %v\n", err)
			return
		}
		_, _ = wr.WriteString("hello ")
		_, _ = wr.Write(line)
		_, _ = wr.WriteString("\n")
		_ = wr.Flush() // 一次性syscall
	}
}
