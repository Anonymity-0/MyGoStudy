package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	// 定义一个字符串 data
	data := "abc123!?$*&()'-=@~"

	// 使用 b64.StdEncoding 的 EncodeToString 方法将 data 的字节切片进行 Base64 编码
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	// 打印编码后的字符串
	fmt.Println(sEnc)

	// 使用 b64.StdEncoding 的 DecodeString 方法将编码后的字符串 sEnc 进行 Base64 解码
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	// 打印解码后的字符串
	fmt.Println(string(sDec))
	fmt.Println()

	// 使用 b64.URLEncoding 的 EncodeToString 方法将 data 的字节切片进行 URL 安全的 Base64 编码
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	// 打印编码后的字符串
	fmt.Println(uEnc)

	// 使用 b64.URLEncoding 的 DecodeString 方法将编码后的字符串 uEnc 进行 URL 安全的 Base64 解码
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	// 打印解码后的字符串
	fmt.Println(string(uDec))

	// 标准 Base64 编码和 URL 安全的 Base64 编码的编码字符串存在稍许不同（后缀为 + 和 -），
	// 但是两者都可以正确解码为原始字符串。
}
