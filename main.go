package main

import (
	"fmt"
	"net"
	"time"
	"share/networking"
)

func broadcast() {
	pc, err := net.ListenPacket("udp4", ":8829")
	if err != nil {
		panic(err)
	}
	defer pc.Close()

	broadcast_addr, err := net.ResolveUDPAddr("udp4", "192.168.2.255:8829")
	if err != nil {
		panic(err)
	}

	_, err = pc.WriteTo([]byte("Hello World"), broadcast_addr)
	if err != nil {
		panic(err)
	}
}

func send(addr net.Addr, message []byte) {
	pc, err := net.ListenPacket("udp4", ":8829")
	if err != nil {
		panic(err)
	}
	defer pc.Close()

	_, err = pc.WriteTo(message, addr)
	if err != nil {
		panic(err)
	}
}

func listen() (net.Addr, int, []byte){
	pc, err := net.ListenPacket("udp4", ":8829")
	if err != nil {
		panic (err)
	}
	defer pc.Close()

	buf := make([]byte, 1024)
	n, addr, err := pc.ReadFrom(buf)
	if err != nil {
		panic (err)
	}
	return addr, n, buf
}

func gatherUsers() ([]net.Addr) {
	var users []net.Addr
	c1 := make(chan net.Addr, 1)
	go func() {
		split, _, _:= listen()
		c1 <- split 
	}()

	select {
	case res := <-c1:
		users = append(users, res)
	case <-time.After(5 * time.Second):
		fmt.Println("Timeout")
	}

	return users
}




func main() {
	
	var answer string
	fmt.Printf("1 for listen, 2 for broadcast: ")
	fmt.Scanln(&answer)
	fmt.Println("")
	if answer == "1" {
		response_addr, _, _ := listen()
		send(response_addr, []byte("Your message is received"))
	}

	if answer == "2" {
		fmt.Println("UPDATE: Sending Broadcast")
		broadcast()
		fmt.Println(gatherUsers())	
	}
}
