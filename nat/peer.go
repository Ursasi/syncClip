package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

var tag string

const HandShakeMsg = "udp hole"
const ServerIp = "xxx.xxx.xxx.xxx"

func main() {
	tag = "mac"
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 9983}
	dstAddr := &net.UDPAddr{IP: net.ParseIP(ServerIp), Port: 9981}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		fmt.Println(err)
	}
	if _, err = conn.Write([]byte("hello, I'm new peer:" + tag)); err != nil {
		log.Panic(err)
	}
	data := make([]byte, 1024)
	n, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Printf("error during read: %s", err)
	}
	conn.Close()
	anotherPeer := parseAddr(string(data[:n]))
	fmt.Printf("local:%s server:%s another:%s\n", srcAddr, remoteAddr, anotherPeer.String())
	bidirectionalHole(srcAddr, &anotherPeer)
}
func parseAddr(addr string) net.UDPAddr {
	t := strings.Split(addr, ":")
	port, _ := strconv.Atoi(t[1])
	return net.UDPAddr{
		IP:   net.ParseIP(t[0]),
		Port: port,
	}
}
func bidirectionalHole(srcAddr *net.UDPAddr, anotherAddr *net.UDPAddr) {
	conn, err := net.DialUDP("udp", srcAddr, anotherAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	if _, err = conn.Write([]byte(HandShakeMsg)); err != nil {
		log.Println("send handshake:", err)
	}
	go func() {
		for {
			time.Sleep(10 * time.Second)
			if _, err = conn.Write([]byte("from [" + tag + "]")); err != nil {
				log.Println("send msg fail", err)
			}
		}
	}()
	for {
		data := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(data)
		if err != nil {
			log.Printf("error during read: %s\n", err)
		} else {
			log.Printf("receive from %s data :%s\n", addr.String(), data[:n])
		}
	}
}
