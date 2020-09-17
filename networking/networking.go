package networking

import (
	"fmt"
	"net"
	"time"
)

/*
* TODO: Add in error codes to see if proces was succesful
 */



func SendMessage(ipa net.Addr,message string) {
	pc, err := net.ListenPacket("udp4", ":8829")
	if err != nil {
		panic(err)
	}
	defer pc.Close()


	addr, err := net.ResolveUDPAddr("udp4", "192.168.2.255:8829")
	if err != nil {
		panic(err)
	}

	_, err = pc.WriteTo([]byte(message), addr)
	if err != nil {
		panic(err)
	}
}

func Listener() (net.Addr, int, []byte) {
	pc, err := net.ListenPacket("udp4", ":8829")
	if err != nil {
		panic(err)
	}
	defer pc.Close()

	buf := make([]byte, 1024)
	n, addr, err := pc.ReadFrom(buf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s sent: %s\n", addr, buf[:n])
	return addr, n, buf
}

// Collects responding users in array
func UserCollect() ([]net.Addr) {


	// Necessary setup
	var users []net.Addr
	c1 := make(chan net.Addr, 1)

	addr, err := net.ResolveUDPAddr("udp4", "255.255.255.255:8829")
	if err != nil {
		panic(err)
	}

	// Run a broadcast
	SendMessage(addr, "data")

	// Run Listener for one second to see
	// if users are found, if not do a timeout
	go func () {
		ipa, _, _ := Listener()
		c1 <- ipa
	}()

	// Timeout setup
	select {
	case res := <-c1:
		users = append(users, res)
	case <-time.After(1 * time.Second):
		break
	}

	// Return is in an array of IP addresses (net.Addr)
	return users
}
