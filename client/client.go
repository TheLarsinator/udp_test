package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func getIP() net.IP {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			return ipv4
		}
	}
	return nil
}

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	ip := getIP()
	if ip != nil {
		msg := "Hi client, I am: 1-" + ip.String()
		_, err := conn.WriteToUDP([]byte(msg), addr)
		if err != nil {
			fmt.Printf("Couldn't send response %v", err)
		}
	}
	fmt.Printf("Couldn't find IP to broadcast")
}

func main() {

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
		log.Println(n, "bytes read from", peer, "saying", string(b[0:n]))
		/*if _, err := c.WriteTo(b[:n], peer); err != nil {
			log.Fatal(err)
		}*/
	}
}
