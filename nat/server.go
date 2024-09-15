package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 9981})
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Printf("local addr: <%s> \n", listener.LocalAddr().String())
	peers := make([]net.UDPAddr, 0, 2)
	data := make([]byte, 1024)
	for {
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s", err)
		}
		log.Printf("<%s> %s\n", remoteAddr.String(), data[:n])
		peers = append(peers, *remoteAddr)
		if len(peers) == 2 {
			log.Printf("establishing %s <--> %s \n", peers[0].String(), peers[1].String())
			_, err := listener.WriteToUDP([]byte(peers[1].String()), &peers[0])
			if err != nil {
				log.Println("write to udp err:", err)
				return
			}
			_, err = listener.WriteToUDP([]byte(peers[0].String()), &peers[1])
			if err != nil {
				log.Println("write to udp err:", err)
				return
			}
			time.Sleep(time.Second * 8)
			log.Println("server exit, peers can still communicate with each other")
			return
		}
	}
}
