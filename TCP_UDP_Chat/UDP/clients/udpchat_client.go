package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const SERVER_IP = "127.0.0.1" 
const SERVER_PORT = 12345

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", SERVER_IP, SERVER_PORT))
	if err != nil {
		log.Fatal("หา Server ไม่เจอ:", err)
	}

	conn, err := net.ListenUDP("udp", nil)
	if err != nil {
		log.Fatal("เปิด Port ไม่ได้:", err)
	}
	defer conn.Close()


	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Print("กรุณาใส่ชื่อของคุณ: ")
	name, _ := consoleReader.ReadString('\n')
	name = strings.TrimSpace(name)

	go receive(conn)

	signupMsg := fmt.Sprintf("!SIGNUP:%s", name)
	conn.WriteToUDP([]byte(signupMsg), serverAddr)

	for {
		msg, _ := consoleReader.ReadString('\n')
		msg = strings.TrimSpace(msg)

		if msg == "!QUIT" {
			fmt.Println("กำลังออกจากโปรแกรม...")
			os.Exit(0) 
		}

		fullMsg := fmt.Sprintf("%s: %s", name, msg)
		_, err = conn.WriteToUDP([]byte(fullMsg), serverAddr)
		if err != nil {
			log.Println("ส่งข้อความไม่สำเร็จ:", err)
		}
	
	}
	
}

func receive(conn *net.UDPConn) {
	buffer := make([]byte, 2048)
	for {
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println("รับข้อความล้มเหลว:", err)
			return 
		}

		fmt.Println(string(buffer[:n]))
	}

}