package main

import (
	"fmt"
	"io"
	"os"
)

// 你需要将这个值设置为 Constants.MAXPAGES 的实际值
const MAXPAGES = 50000

// DSMgr 代表一个磁盘空间管理器，它负责管理磁盘上的页面。
type DSMgr struct {
	currentFile *os.File      // currentFile 是当前打开的文件。
	numPages    int           // numPages 是当前文件中的页面数量。
	pages       [MAXPAGES]int // pages 是一个数组，用于记录每个页面的使用情况。
	ICounter    int           // ICounter 是读取操作的计数器。
	OCounter    int           // OCounter 是写入操作的计数器。
}

// NewDSMgr 创建一个新的 DSMgr，并初始化它的字段。
func NewDSMgr() *DSMgr {
	dsmgr := &DSMgr{
		currentFile: nil,
		numPages:    0,
		pages:       [MAXPAGES]int{},
		ICounter:    0,
		OCounter:    0,
	}
	return dsmgr
}

// openFile 打开一个文件，并将其设置为当前文件。
func (d *DSMgr) openFile(fileName string) int {
	var err error
	d.currentFile, err = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return 1
}

// closeFile 关闭当前文件。
func (d *DSMgr) closeFile() int {
	err := d.currentFile.Close()
	if err != nil {
		fmt.Println(err)
		return 0
	}
	d.currentFile = nil
	return 1
}

// readPage 从当前文件中读取一个页面，并返回一个 BFrame。
func (d *DSMgr) readPage(page_id int) *BFrame {
	buffer := make([]byte, FRAMESIZE)
	_, err := d.currentFile.ReadAt(buffer, int64(page_id*FRAMESIZE))
	if err != nil {
		if err == io.EOF {
			//fmt.Println("文件已读取到末尾")
		} else {
			fmt.Println("文件打开失败")
		}
	}
	d.ICounter++
	return &BFrame{}
}

// writePage 将一个 BFrame 写入到当前文件的一个页面中。
func (d *DSMgr) writePage(page_id int, frm *BFrame) int {
	// 这里需要根据实际情况将 BFrame 转换为 []byte
	_, err := d.currentFile.WriteAt([]byte{}, int64(page_id*FRAMESIZE))
	if err != nil {
		fmt.Println("文件写异常")
	}
	d.OCounter++
	return FRAMESIZE
}

// seek 将当前文件的读写位置移动到指定的位置。
func (d *DSMgr) seek(offset, pos int) int {
	_, err := d.currentFile.Seek(int64(pos+offset), 0)
	if err != nil {
		fmt.Println("文件seek异常")
		return 0
	}
	return 1
}

// incNumPages 增加当前文件的页面数量。
func (d *DSMgr) incNumPages() {
	d.numPages++
}

// setUse 设置一个页面的使用位。
func (d *DSMgr) setUse(page_id, use_bit int) {
	d.pages[page_id] = use_bit
}

// getUse 返回一个页面的使用位。
func (d *DSMgr) getUse(page_id int) int {
	return d.pages[page_id]
}

// getNumPages 返回当前文件的页面数量。
func (d *DSMgr) getNumPages() int {
	return d.numPages
}

// getPages 返回 pages 数组的副本。
func (d *DSMgr) getPages() [MAXPAGES]int {
	return d.pages
}
