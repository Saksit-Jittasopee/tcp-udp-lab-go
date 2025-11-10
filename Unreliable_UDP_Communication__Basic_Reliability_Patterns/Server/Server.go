package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: 8080,
		IP:   net.ParseIP("0.0.0.0"),
	})
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer conn.Close()

	fmt.Println("UDP server listening on port 8080")

	buffer := make([]byte, 1024)

	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("ERROR: ", err)
			continue
		}

		msg := string(buffer[:n])
		fmt.Printf("Received from %v: %s\n", addr, msg)

		parts := strings.SplitN(msg, ":", 2)
		if len(parts) > 1 {
			seq := parts[0]
			ackMsg := fmt.Sprintf("ACK:%s", seq)
			conn.WriteToUDP([]byte(ackMsg), addr)
			fmt.Println("Sent:", ackMsg)
		}
	}
}