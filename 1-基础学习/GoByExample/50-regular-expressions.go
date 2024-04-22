package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	// 使用 MatchString 函数检查字符串 "peach" 是否匹配正则表达式 "p([a-z]+)ch"
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println("match:", match)

	// 使用 Compile 函数编译正则表达式 "p([a-z]+)ch"
	r, _ := regexp.Compile("p([a-z]+)ch")
	// 使用编译后的正则表达式 r 检查字符串 "peach" 是否匹配
	fmt.Println("r.MatchString:", r.MatchString("peach"))

	// 使用 FindString 函数在字符串 "peach punch" 中查找第一个匹配的子串
	fmt.Println("r.FindString:", r.FindString("peach punch"))
	// 使用 FindStringIndex 函数在字符串 "peach punch" 中查找第一个匹配的子串的起始和结束位置
	fmt.Println("r.FindStringIndex:", r.FindStringIndex("peach punch"))
	// 使用 FindStringSubmatch 函数在字符串 "peach punch" 中查找第一个匹配的子串和它的子匹配
	fmt.Println("r.FindStringSubmatch:", r.FindStringSubmatch("peach punch"))
	// 使用 FindStringSubmatchIndex 函数在字符串 "peach punch" 中查找第一个匹配的子串和它的子匹配的起始和结束位置
	fmt.Println("r.FindStringSubmatchIndex:", r.FindStringSubmatchIndex("peach punch"))
	// 使用 FindAllString 函数在字符串 "peach punch pinch" 中查找所有匹配的子串，-1 表示没有限制
	fmt.Println("r.FindAllString:", r.FindAllString("peach punch pinch", -1))

	// 使用 FindAllStringSubmatchIndex 函数在字符串 "peach punch pinch" 中查找所有匹配的子串和它们的子匹配的起始和结束位置，-1 表示没有限制
	fmt.Println("all:", r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	// 使用 FindAllString 函数在字符串 "peach punch pinch" 中查找前两个匹配的子串
	fmt.Println("r.FindAllString2", r.FindAllString("peach punch pinch", 2))

	// 使用 Match 函数检查字节切片 "peach" 是否匹配正则表达式 r
	fmt.Println("r.Match:", r.Match([]byte("peach")))
	// 使用 MustCompile 函数编译正则表达式 "p([a-z]+)ch"，如果编译失败，它会抛出 panic
	r = regexp.MustCompile("p([a-z]+)ch")
	// 使用 ReplaceAllString 函数在字符串 "a peach" 中替换所有匹配的子串为 "<fruit>"
	fmt.Println("r.ReplaceAllString:", r.ReplaceAllString("a peach", "<fruit>"))
	in := []byte("a peach")
	// 使用 ReplaceAllFunc 函数在字节切片 in 中替换所有匹配的子串为它们的大写形式
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println("r.ReplaceAllFunc:", string(out))
}
