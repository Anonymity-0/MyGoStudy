package main

import (
	"fmt"
)

// FrameSize 是一个常量，表示 BFrame 中的 Field 字段的大小。
const FrameSize = 1024

// BFrame 代表一个缓冲帧，它包含了一个固定大小的 rune 数组。
type BFrame struct {
	Field [FrameSize]rune // Field 是一个固定大小的 rune 数组，用于存储数据。
}

// NewBFrame 创建一个新的 BFrame，并使用给定的 buffer 初始化它的 Field 字段。
// buffer 是一个字节切片，它的内容会被复制到新的 BFrame 的 Field 字段中。
func NewBFrame(buffer []byte) *BFrame {
	b := &BFrame{}
	copy(b.Field[:], []rune(fmt.Sprint(buffer)))
	return b
}

// GetField 返回 BFrame 的 Field 字段的内容，作为一个字节切片。
func (b *BFrame) GetField() []byte {
	return []byte(fmt.Sprint(b.Field))
}
