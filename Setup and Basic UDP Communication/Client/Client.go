// client.go (Snippet - use a loop to send 10 sequential messages)
// ...
package main

import (
	"fmt"
	"net"
    "time"
)

func main() {
    conn, _ := net.Dial("udp", "127.0.0.1:8080")    
    for i := 1; i <= 10; i++ {
    message := fmt.Sprintf("UDP Segment #%d", i)
    conn.Write([]byte(message))
    time.Sleep(10 * time.Millisecond) // Slow down slightly to observe
    }

}

// ...