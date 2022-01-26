package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"path"
	"runtime"
	"strings"

)

func getFilename()  string{
	_, fullFilename, _, _ := runtime.Caller(0)
	fmt.Println(fullFilename)
	var filename string
	filename = path.Base(fullFilename)
	return filename
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
	}
	return "",false
}


func main() {



	addS:="IAO.exe"
	fmt.Println(addS)
	sp:=strings.SplitN(addS,".",2)
	fmt.Println(len(sp),sp[0])

	//contentA,err:=clipboard.ReadAll()
	//if err!=nil{
	//	panic(err)
	//}

	orig := "hello world"
	key := "123456781234567812345678"
	fmt.Println("原文：", orig)


	encryptCode := AesEncrypt(orig, key)
	fmt.Println("密文：" , encryptCode)



	decryptCode := AesDecrypt(encryptCode, key)


	fmt.Println("解密结果：", decryptCode)


	//for range time.Tick(1*time.Second){

		//fmt.Println(string(clipboard.Read(clipboard.FmtText)))
		//clipboard.Write(clipboard.FmtText, []byte("Hello, 世界"))
		//content,err:=clipboard.ReadAll()
		//if err!=nil{
		//	panic(err)
		//}
		//if content!=contentA{
		//	fmt.Println(content)
		//	contentA=content
		//}
	//}
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

