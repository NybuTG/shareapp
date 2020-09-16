package networking

import (
	"net"
	"time"
	"share/util"
)

var BROADCAST_ADDRESS, _ = net.ResolveUDPAddr("udp4", "255.255.255.255")

/*
* TODO: Add in error codes to see if proces was succesful
*/

func handler(pc net.PacketConn, err error) {
	util.Check(err)
	pc.Close()
}


func SendMessage(ipa net.Addr,message string) {
	pc, err := net.ListenPacket("udp4", ":8829")
	handler(pc, err)
	
	util.Check(err)

	_, err = pc.WriteTo([]byte(message), ipa)
	util.Check(err)
}

func Listener() (net.Addr, int, []byte) {
	pc, err := net.ListenPacket("udp4", ":8829")
	handler(pc, err)

	buf := make([]byte, 1024)
	n, addr, err := pc.ReadFrom(buf)
	if err != nil {
		panic(err)
	}
	return addr, n, buf
}

// Collects responding users in array
func UserCollect() ([]net.Addr) {

	// Necessary setup
	var users []net.Addr
	c1 := make(chan net.Addr, 1)

	// Run a broadcast
	SendMessage(BROADCAST_ADDRESS, "add_data_here")

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
