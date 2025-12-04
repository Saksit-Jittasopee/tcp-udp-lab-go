package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const SERVER_PORT = "12345"

func main() {
	// connect to server (localhost)
	host := "localhost:" + SERVER_PORT
	conn, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server on", host)

	reader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(conn)

	for {
		fmt.Print(">> ")
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		// send message to server
		_, err = conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Println("Error sending:", err)
			break
		}

		if strings.ToLower(message) == "bye" {
			fmt.Println("Connection closed.")
			break
		}

		// receive message from server
		reply, err := serverReader.ReadString('\n')
		if err != nil {
			fmt.Println("Server disconnected.")
			break
		}
		fmt.Println("server:", strings.TrimSpace(reply))
	}
}
