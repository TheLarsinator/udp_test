package main

import (
	"log"
	"net"
	"time"
)
func main() {
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
			if _, err := c.WriteTo([]byte("hello"), dst); err != nil {
			log.Fatal(err)
			}
			time.Sleep(1000 * time.Millisecond)
		}
	}
