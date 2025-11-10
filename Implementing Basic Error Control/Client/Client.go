// client.go (Snippet - use a loop to send 10 sequential messages)
// ...
package main

import (
	"fmt"
	"net"
    "time"
)

func main() {
    conn, _ := net.DialUDP("udp", nil, &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8080})
    for i := 1; i <= 10; i++ {
    message := fmt.Sprintf("UDP Segment #%d", i)
    conn.Write([]byte(message))
    time.Sleep(10 * time.Millisecond) // Slow down slightly to observe

    conn.SetReadDeadline(time.Now().Add(50 * time.Millisecond)) 
    ackBuffer := make([]byte, 1024)
     _, _, err := conn.ReadFromUDP(ackBuffer)
 
    if err != nil {
        fmt.Printf("TIMEOUT: Retransmitting %s\n", message)
        conn.Write([]byte(message)) // Simple retransmission
    } else {
        fmt.Printf("ACK received for %s\n", message)
    }
    }

}

// ...