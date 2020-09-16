package main

import (
	"fmt"
	"net"
	"share/networking"
	"share/util"
)

func server() {
	broadcast, err := net.ResolveUDPAddr("upd4", "255.255.255.255")
	util.Check(err)
	networking.UserCollect

}

func client() {
	
}


func main() {
	fmt.Printf("1 for Scan, 2 for Listen:")
	answer, _ := fmt.Scanf("%s")
	if answer == 1 {
		// Run Scan func
	} else {
		// Run listen func
	}
}

