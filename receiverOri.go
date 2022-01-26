package main

import (
	"log"
	"net"
)

func main() {
	net.InterfaceAddrs()
	laddr := net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0), //写局域网下分配IP，0.0.0.0可以用来测试
		Port: 9999,
	}
	conn, err := net.ListenUDP("udp", &laddr)
	if err != nil {
		log.Panicln(err.Error())
	}
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Panicln(err.Error())
		}
		log.Println(string(buf[:n]))
	}
}

