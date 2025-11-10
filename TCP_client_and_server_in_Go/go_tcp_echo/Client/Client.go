package main

import (
	"fmt"
	"net"
    "time"
)

func main() {
    conn, _ := net.Dial("tcp", "localhost:8080")
    for i := 1; i <= 10; i++ {
    message := fmt.Sprintf("Hello from the client!: %d", i)
    conn.Write([]byte(message))
    time.Sleep(10 * time.Millisecond) // Slow down slightly to observe

    conn.SetReadDeadline(time.Now().Add(50 * time.Millisecond)) 
    Buffer := make([]byte, 1024)
     _, err := conn.Read(Buffer)
    if err != nil {
        fmt.Printf("ERROR: %s\n", message)
        conn.Write([]byte(message)) // Simple retransmission
    } else {
        fmt.Printf("%s\n", message)
    }
    }

}
