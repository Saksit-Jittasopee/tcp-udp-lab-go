// server.go

package main

import (
    "fmt"
    "net"
)

func main() {
    addr, _ := net.ResolveUDPAddr("udp", ":8080")
    conn, _ := net.ListenUDP("udp", addr)
    defer conn.Close()
    fmt.Println("UDP Server listening on :8080")
    buffer := make([]byte, 1024)    
	for {
		n, remoteAddr, _ := conn.ReadFromUDP(buffer)
        fmt.Printf("Received: %s from %s\n", string(buffer[:n]), remoteAddr.String())
}
} 