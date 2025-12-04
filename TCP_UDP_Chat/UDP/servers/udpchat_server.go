package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

const LISTENING_PORT = 12345

func main() {
	type Message struct { 
		data []byte
		addr *net.UDPAddr
	}
	var messages = make(chan Message, 100) 
	var clients = make(map[string]*net.UDPAddr) 
	var clientsMutex sync.Mutex 
	addr := net.UDPAddr{
		Port: LISTENING_PORT,
		IP:   net.ParseIP("0.0.0.0"),
	}
	conn, err := net.ListenUDP("udp", &addr) 
	if err != nil {
		log.Fatal("เปิด Port ไม่ได้:", err)
	}
	defer conn.Close()
	fmt.Println("UDP Chat Server เปิดแล้วที่ Port", LISTENING_PORT)

	go func() {
		buffer := make([]byte, 1024) 
		for {
			n, remoteAddr, err := conn.ReadFromUDP(buffer) 
			if err != nil {
				log.Println("Error reading:", err)
				continue
			}

			dataCopy := make([]byte, n) 
			copy(dataCopy, buffer[:n]) 
			messages <- Message{data: dataCopy, addr: remoteAddr} 
		}
	}() 
	go func() {
		for msg := range messages { 
			fmt.Printf("RECV: %s (from %s)\n", string(msg.data), msg.addr) 
			clientsMutex.Lock() 
			addrString := msg.addr.String() 
			if _, ok := clients[addrString]; !ok { 
				clients[addrString] = msg.addr 
				fmt.Println("Client ใหม่:", addrString) 
			}
			clientsMutex.Unlock() 

			var broadcastData []byte 
			msgString := string(msg.data) 

			if strings.HasPrefix(msgString, "!SIGNUP:") { 
				name := strings.SplitN(msgString, ":", 2)[1] //
				broadcastData = []byte(fmt.Sprintf("SERVER_MSG: %s ได้เข้าร่วมห้อง!", name)) 
			} else {
				broadcastData = msg.data 
			}

			var clientsToRemove []string 

			clientsMutex.Lock() 
			for strAddr, clientAddr := range clients { 
				_, err := conn.WriteToUDP(broadcastData, clientAddr) 
				if err != nil {
					log.Println("ส่งไม่สำเร็จให้:", clientAddr, err) 
					clientsToRemove = append(clientsToRemove, strAddr) 
				}
			}
			clientsMutex.Unlock() 

			if len(clientsToRemove) > 0 { 
				clientsMutex.Lock() 
				for _, addrToRemove := range clientsToRemove { 
					fmt.Println("Client", addrToRemove, "ถูกลบ (ส่งไม่สำเร็จ)") 
					delete(clients, addrToRemove) 
				}
				clientsMutex.Unlock() 
			}
		}
	}() 

	select {}
}