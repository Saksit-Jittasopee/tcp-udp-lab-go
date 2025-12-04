package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const LISTENING_PORT = "12345"

func main() {
	// สร้าง listener บน port
	listener, err := net.Listen("tcp", ":"+LISTENING_PORT)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on port", LISTENING_PORT)

	// รอรับการเชื่อมต่อจาก client
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		return
	}
	defer conn.Close()

	fmt.Println("A new connection from:", conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)
	input := bufio.NewReader(os.Stdin)

	for {
		// อ่านข้อความจาก client
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected.")
			break
		}
		message = strings.TrimSpace(message)
		fmt.Println("client:", message)

		// อ่านข้อความจากฝั่ง server (console)
		fmt.Print(">> ")
		reply, _ := input.ReadString('\n')
		reply = strings.TrimSpace(reply)

		// ส่งกลับไปยัง client
		_, err = conn.Write([]byte(reply + "\n"))
		if err != nil {
			fmt.Println("Error sending message:", err)
			break
		}
		
	}
}
