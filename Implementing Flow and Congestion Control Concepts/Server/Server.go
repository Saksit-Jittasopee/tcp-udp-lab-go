// server.go
package main

import (
	"fmt"
	"math/rand/v2"
	"net"
)

func main() {
	// Listen on UDP port 8080
	addr, _ := net.ResolveUDPAddr("udp", ":8080")
	conn, _ := net.ListenUDP("udp", addr)
	defer conn.Close()
	fmt.Println("UDP Server listening on :8080")
	buffer := make([]byte, 1024)
	for {
		// server.go (Modification)
		// ... inside the loop after receiving data ...
		n, remoteAddr, _ := conn.ReadFromUDP(buffer)
		// Simulate processing delay or potential loss
		if rand.IntN(10) < 2 { // 20% chance to 'lose' the ACK
			fmt.Printf("ERROR: Dropping ACK for %s\n", string(buffer[:n]))
		} else {
			ack := []byte("ACK:" + string(buffer[:n]))
			conn.WriteToUDP(ack, remoteAddr) // Send ACK back
		}
		// ...
	}
}
