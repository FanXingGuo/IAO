package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	//"golang.design/x/clipboard"
	"github.com/atotto/clipboard"
	"log"
	"net"
	"path"
	"runtime"
	"strings"
	"sync"
	"time"
)

func getFilename()  string{
	_, fullFilename, _, _ := runtime.Caller(0)
	var filename string
	filename = path.Base(fullFilename)
	return strings.SplitN(filename,".",2)[0]
}
func EnCode(s string) string {
	key := "123456781234567812345678"
	filename:=getFilename()
	orig:=filename+":"+s
	encryptCode := AesEncrypt(orig, key)
	return encryptCode
}
func DeCode(s string)(string,bool)  {
	key := "123456781234567812345678"
	filename:=getFilename()
	decryptCode := AesDecrypt(s, key)
	spt:=strings.SplitN(decryptCode,":",2)
	if spt[0]==filename{
		return spt[1],true
	}else {
		return "",false
	}
}

func AesEncrypt(orig string, key string) string {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)

	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)

	return base64.StdEncoding.EncodeToString(cryted)

}

func AesDecrypt(cryted string, key string) string {
	// 转成字节数组
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	k := []byte(key)

	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return string(orig)
}

//补码
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func sender(){
	conn,err:=net.Dial("udp","255.255.255.255:9999")
	if err!=nil{
		panic(err)
	}
	defer conn.Close()
	contentA,_:=clipboard.ReadAll()

	for range time.Tick(500*time.Millisecond){
		content,_:=clipboard.ReadAll()
		if content!=contentA{
			//log.Println(content)
			s:=EnCode(content)
			conn.Write([]byte(s))
			contentA=content
		}
	}

}
func receiver(){
	serAddr,err:=net.ResolveUDPAddr("udp","0.0.0.0:9999")
	if err!=nil{
		panic(err)
	}
	udpConn,err:=net.ListenUDP("udp",serAddr)
	if err!=nil{
		panic(err)
	}
	defer udpConn.Close()

	buf:=make([]byte,32768)
	contentA,_:=clipboard.ReadAll()
	for{
		n,cltAddr,err:=udpConn.ReadFromUDP(buf)
		if err!=nil{
			panic(err)
		}
		data:=string(buf[:n])
		contentNow,isok:=DeCode(data)
		if isok{
			if contentNow!=contentA{
				clipboard.WriteAll(contentNow)
				//clipboard.WriteAll(string(buf[:n]))
				log.Println(cltAddr,contentNow)
				contentA=contentNow
			}
		}
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go receiver()
	go sender()
	wg.Wait()
}

