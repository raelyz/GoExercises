package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func StartTCPServer() {
	myListener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer myListener.Close()
	for {
		conn, err := myListener.Accept()
		fmt.Println(conn.RemoteAddr().String())
		fmt.Println(conn.(net.Conn))
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	fmt.Println(conn.RemoteAddr().String())
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("Received:", string(data))
		retMsg := string(data) + "\n"
		conn.Write([]byte(retMsg))
	}
	conn.Close()
}
