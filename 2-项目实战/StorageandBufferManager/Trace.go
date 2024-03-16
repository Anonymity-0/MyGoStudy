package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
)

const FRAMESIZE = 4096

// Trace 结构体用于模拟数据库系统的追踪操作。
type Trace struct {
	BMgr      *BMgr          // BMgr 是缓冲区管理器，用于管理缓冲区中的页面。
	HitRate   float64        // HitRate 是缓冲区命中率。
	IOCounter int            // IOCounter 是输入/输出操作的计数器。
	wg        sync.WaitGroup // wg 是等待组，用于等待所有 goroutine 完成。
}

// NewTrace 创建一个新的 Trace，并初始化它的字段。
func NewTrace() *Trace {
	return &Trace{
		BMgr: NewBMgr(),
		wg:   sync.WaitGroup{},
	}
}

// CreateFile 创建一个新的文件，并填充初始数据。
func (t *Trace) CreateFile() {
	bu := make([]byte, FRAMESIZE)
	for i := 0; i < FRAMESIZE; i++ {
		bu[i] = '1'
	}
	f, _ := os.Create("data.dbf")
	defer f.Close()
	for j := 0; j < MAXPAGES; j++ {
		f.Write(bu)
	}
}

// Read 读取给定页面 ID 的页面。
// 这是一个异步操作，会启动一个新的 goroutine 来执行。
func (t *Trace) Read(page_id int) {
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		t.BMgr.latch.Lock()
		defer t.BMgr.latch.Unlock()
		t.BMgr.fixPage(page_id, 0)
		if t.BMgr.unFixPage(page_id) == -1 {
			fmt.Println("read异常: 页面释放错误")
		}
	}()
}

// Write 写入给定页面 ID 的页面。
// 这是一个异步操作，会启动一个新的 goroutine 来执行。
func (t *Trace) Write(page_id int) {
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		t.BMgr.latch.Lock()
		defer t.BMgr.latch.Unlock()
		t.BMgr.setDirty(t.BMgr.fixPage(page_id, 1))
		if t.BMgr.unFixPage(page_id) == -1 {
			fmt.Println("write异常: 页面释放错误")
		}
	}()
}

// GetStatistics 从测试文件中读取操作，并执行这些操作。
// 在所有操作执行完成后，计算缓冲区的命中率和输入/输出操作的数量。
func (t *Trace) GetStatistics() {
	var testArrayList []string
	data, err := ioutil.ReadFile("data-5w-50w-zipf.txt")
	if err != nil {
		fmt.Println("测试文件读取异常")
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			testArrayList = append(testArrayList, line)
		}
	}
	fmt.Println("---测试文件读取完成---")

	fmt.Println("---执行trace操作中---")
	for _, line := range testArrayList {
		temp_str2 := strings.Split(line, ",")
		operation, _ := strconv.Atoi(temp_str2[0])
		page_id, _ := strconv.Atoi(temp_str2[1])
		page_id--           // 文件中的页号从1开始
		if operation == 0 { // 读操作
			t.Read(page_id)
		} else { //写操作
			t.Write(page_id)
		}
	}
	t.Finish()
	t.HitRate = float64(t.BMgr.HitCounter) / float64(len(testArrayList))
	t.IOCounter = t.BMgr.DSMgr.ICounter + t.BMgr.DSMgr.OCounter
}

// Finish 等待所有操作完成，并将缓冲区中所有脏的页面写回磁盘。
func (t *Trace) Finish() {
	t.wg.Wait()
	t.BMgr.latch.Lock()
	defer t.BMgr.latch.Unlock()
	t.BMgr.writeDirtys()
}
