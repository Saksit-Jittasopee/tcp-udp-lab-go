package main

import (
	"fmt"
	"net"
)

func main() {
	listener, _ := net.Listen("tcp", ":8080");
	defer listener.Close();
	fmt.Println("TCP server listening on port 8080");

	for {
		conn, _ := listener.Accept();
		go handleConnection(conn);
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close();
	buffer := make([]byte, 1024);

	for {
		n, _ := conn.Read(buffer);

		message := string(buffer[:n]);
		fmt.Printf("Received: %s\n", message);

		_, _ = conn.Write([]byte(message));
	}
}