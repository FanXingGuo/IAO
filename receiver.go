package main

import (
	"fmt"
	"net"
)

func main() {
	serAddr,err:=net.ResolveUDPAddr("udp","127.0.0.1:9999")
	if err!=nil{
		panic(err)
	}
	udpConn,err:=net.ListenUDP("udp",serAddr)
	if err!=nil{
		panic(err)
	}
	defer udpConn.Close()

	buf:=make([]byte,4096)
	n,cltAddr,err:=udpConn.ReadFromUDP(buf)
	if err!=nil{
		panic(err)
	}
	fmt.Println(cltAddr,string(buf[:n]))

	//net.InterfaceAddrs()
	//laddr := net.UDPAddr{
	//	IP:   net.IPv4(0, 0, 0, 0), //写局域网下分配IP，0.0.0.0可以用来测试
	//	Port: 9999,
	//}
	//connR, errR := net.ListenUDP("udp", &laddr)
	//if errR != nil {
	//	log.Panicln(errR.Error())
	//}
	//defer connR.Close()
	//buf := make([]byte, 1024)
	//for {
	//	n, err := connR.Read(buf)
	//	if err != nil {
	//		log.Panicln(err.Error())
	//	}
	//	log.Println(string(buf[:n]))
	//}
}