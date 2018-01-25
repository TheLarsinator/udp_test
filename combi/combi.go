package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

type elevator_addr struct {
	id int
	ip string
}

func broadcast_ip(id int) {
	c, err := net.ListenPacket("udp", ":0")

	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	dst, err := net.ResolveUDPAddr("udp", "255.255.255.255:8032")
	if err != nil {
		log.Fatal(err)
	}
	for {
		if _, err := c.WriteTo([]byte(string(id)), dst); err != nil {
			log.Fatal(err)
		}
		time.Sleep(1000 * time.Millisecond)
	}
}

func listenUDP() {
	c, err := net.ListenPacket("udp", ":8032")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	for {
		b := make([]byte, 512)
		n, peer, err := c.ReadFrom(b)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(n, "bytes read from", peer, "saying", b[0])
		//TODO: Send the peer addr and data to our handler, and register the elevator
	}
}

func main() {

	idPtr := flag.Int("id", -1, "Elevator ID")
	flag.Parse()
	fmt.Println(*idPtr)

	time.Sleep(1000 * time.Millisecond)
	go broadcast_ip(*idPtr)
	go listenUDP()

	time.Sleep(100000 * time.Millisecond)
}
