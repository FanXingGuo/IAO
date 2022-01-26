package main

import (
	"fmt"
	"net"
)

func main() {

	conn,err:=net.Dial("udp","127.0.0.1:9999")
	if err!=nil{
		panic(err)
	}
	defer conn.Close()

	conn.Write([]byte("hello"))

	buf:=make([]byte,4096)

	n,err:=conn.Read(buf)
	if err!=nil{
		panic(err)
	}
	fmt.Println(buf[:n])


	//laddr := net.UDPAddr{
	//	IP:   net.IPv4(0, 0, 0, 0),//写局域网下分配IP，0.0.0.0可以用来测试
	//	Port: 9999,
	//}
	//
	//raddr := net.UDPAddr{
	//	IP:   net.IPv4(255, 255, 255, 255), //局域网广播地址
	//	Port: 9999,
	//}
	//
	//conn, err := net.DialUDP("udp", &laddr, &raddr)
	//if err != nil {
	//	log.Panicln(err.Error())
	//}
	//
	//
	//
	//defer conn.Close()
	//scan := bufio.NewScanner(os.Stdin)
	//
	//for scan.Scan() {
	//	mess := scan.Text()
	//	if mess == "quit" {
	//		return
	//	}
	//	_, err := conn.Write([]byte(mess))
	//	if err != nil {
	//		log.Panicln(err.Error())
	//	}
	//}
}

