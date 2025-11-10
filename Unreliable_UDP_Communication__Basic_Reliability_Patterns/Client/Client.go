package main

import (
	"fmt"
	"net"
	"time"
	"strings"
)

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	buffer := make([]byte, 1024)

	for seq := 1; seq <= 5; seq++ {
		message := fmt.Sprintf("%d:Hello", seq)

		for {
			fmt.Println("Sending:", message)
			_, err := conn.Write([]byte(message))
			if err != nil {
				fmt.Println("Error writing:", err)
				return
			}

			conn.SetReadDeadline(time.Now().Add(2 * time.Second))
			n, _, err := conn.ReadFromUDP(buffer)
			if err != nil {
				fmt.Println("Timeout or read error, retrying...")
				continue
			}

			ack := string(buffer[:n])
			fmt.Println("Received:", ack)

			if strings.TrimSpace(ack) == fmt.Sprintf("ACK:%d", seq) {
				fmt.Printf("ACK %d received\n", seq)
				break
			}
		}
		time.Sleep(1 * time.Second)
	}
}