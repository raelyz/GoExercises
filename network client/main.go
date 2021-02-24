package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	StartTCPClient()

}

func StartTCPClient() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln("Connectionfails")
		return
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Keyinyourmessage:")
		message, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, message+"\n")
		// body := `<!DOCTYPEhtml><htmllang="en"><head><metacharet="UTF8"><title></title></head><body><strong>INDEX</strong><br><ahref="/">index</a><br><ahref="/about">about</a><br><ahref="/contact">contact</a><br><ahref="/apply">apply</a><br></body></html>`

		fmt.Fprintf(conn, reader)
		// recMessage, _ := bufio.NewReader(conn).ReadString('\n')
		// fmt.Println("Received:", recMessage)
	}
}
