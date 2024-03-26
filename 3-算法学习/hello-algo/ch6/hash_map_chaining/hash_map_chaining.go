package main

import (
	"fmt"
	"strconv"
	"strings"
)

type pair struct {
	key int
	val string
}
type hashMapChaining struct {
	size          int
	capacity      int
	loadThreshold float64
	extendRadio   int
	buckets       [][]pair
}

// 创建哈希表
func newHashMapChaining() *hashMapChaining {
	buckets := make([][]pair, 4)
	for i := 0; i < 4; i++ {
		buckets[i] = make([]pair, 0)
	}
	return &hashMapChaining{
		size:          0,
		capacity:      4,
		loadThreshold: 2.0 / 3.0,
		buckets:       buckets,
	}
}

// 哈希函数
func (h *hashMapChaining) hashFunc(key int) int {
	return key % h.capacity

}

// 负载因子计算
func (h *hashMapChaining) loadFactor() float64 {
	return float64(h.size) / float64(h.capacity)
}

// 查询操作
func (h *hashMapChaining) get(key int) string {
	idx := h.hashFunc(key)
	bucket := h.buckets[idx]
	for _, pair := range bucket {
		if pair.key == key {
			return pair.val
		}
	}
	// 未找到，返回空字符串
	return ""
}

// 添加
func (h *hashMapChaining) put(key int, val string) {
	if h.loadFactor() > h.loadThreshold {
		h.extend()
	}
	idx := h.hashFunc(key)
	//遍历桶中的元素
	for i := range h.buckets[idx] {
		if h.buckets[idx][i].key == key {
			h.buckets[idx][i].val = val
			return
		}
	}
	// 未找到，添加新元素
	h.buckets[idx] = append(h.buckets[idx], pair{key, val})
	h.size++
}

// 删除
func (h *hashMapChaining) remove(key int) {
	idx := h.hashFunc(key)
	for i, p := range h.buckets[idx] {
		if p.key == key {
			//删除元素，将后面的元素向前移动
			h.buckets[idx] = append(h.buckets[idx][:i], h.buckets[idx][i+1:]...)
			h.size--
			break
		}
	}
}

func (h *hashMapChaining) extend() {
	tmpBuckets := make([][]pair, len(h.buckets))
	for i := 0; i < len(h.buckets); i++ {
		tmpBuckets[i] = make([]pair, len(h.buckets[i]))
		// 拷贝元素
		copy(tmpBuckets[i], h.buckets[i])
	}
	// 扩容
	h.capacity *= h.extendRadio
	// 重新分配桶
	h.buckets = make([][]pair, h.capacity)
	for i := 0; i < h.capacity; i++ {
		h.buckets[i] = make([]pair, 0)
	}
	m.size = 0
	// 重新插入元素
	for _, bucket := range tmpBuckets {
		for _, pair := range bucket {
			h.put(pair.key, pair.val)
		}
	}
}

// 打印哈希表
func (h *hashMapChaining) printHash() {
	var builder strings.Builder
	for _, bucket := range h.buckets {
		builder.WriteString("[")
		for _, p := range bucket {
			builder.WriteString(strconv.Itoa(p.key) + " -> " + p.val + " ")
		}
		builder.WriteString("]")
		fmt.Println(builder.String())
		builder.Reset()
	}
}
func main() {

}
