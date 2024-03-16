package main

import (
	"fmt"
	"sync"
)

// DEFBUFSIZE 是一个常量，表示 BMgr 中的缓冲区大小。
const DEFBUFSIZE = 1024

// BMgr 代表一个缓冲区管理器，它负责管理缓冲区中的页面。
type BMgr struct {
	Ftop       [DEFBUFSIZE]int     // Ftop 是一个数组，用于记录每个帧的页面 ID。
	Ptof       [DEFBUFSIZE]*BCB    // Ptof 是一个数组，用于记录每个页面的 BCB。
	Buf        [DEFBUFSIZE]*BFrame // Buf 是一个数组，用于存储缓冲区中的页面。
	DSMgr      *DSMgr              // DSMgr 是一个磁盘空间管理器，用于管理磁盘上的页面。
	LRU        *LRU                // LRU 是一个最近最少使用页面替换算法的实现。
	HitCounter int                 // HitCounter 是一个计数器，用于记录缓冲区命中的次数。
	latch      sync.RWMutex        // latch 是一个读写锁，用于并发控制。
}

// NewBMgr 创建一个新的 BMgr，并初始化它的字段。
func NewBMgr() *BMgr {
	dsmgr := NewDSMgr()
	bmgr := &BMgr{
		DSMgr:      dsmgr,
		LRU:        NewLRU(),
		Buf:        [DEFBUFSIZE]*BFrame{},
		Ftop:       [DEFBUFSIZE]int{},
		Ptof:       [DEFBUFSIZE]*BCB{},
		HitCounter: 0,
	}
	for i := 0; i < DEFBUFSIZE; i++ {
		bmgr.Ptof[i] = nil
		bmgr.Ftop[i] = -1
	}
	if dsmgr.openFile("data.dbf") == 0 {
		fmt.Println("文件打开失败")
	}
	return bmgr
}

// Close 关闭 BMgr，包括关闭当前文件和清空缓冲区。
func (b *BMgr) Close() {
	if b.DSMgr.closeFile() == 0 {
		fmt.Println("文件关闭失败")
	}
}

// / hash 是一个哈希函数，用于将页面 ID 映射到一个较小的范围。
func (b *BMgr) hash(page_id int) int {
	return page_id % DEFBUFSIZE
}

// fixPage 是一个核心方法，用于处理页面的请求。
// 如果请求的页面在缓冲区中，则直接返回该页面的帧号；
// 如果请求的页面不在缓冲区中，则需要从磁盘中读取该页面，并将其加载到缓冲区中。
// page_id 是请求的页面的 ID，prot 是请求的类型（0 表示读取，1 表示写入）。
// 方法返回请求的页面在缓冲区中的帧号。
func (b *BMgr) fixPage(page_id int, prot int) int {
	bcb := b.Ptof[b.hash(page_id)]
	for bcb != nil && bcb.page_id != page_id {
		bcb = bcb.next
	}

	if bcb == nil {
		vFrameId := b.selectVictim()
		newBCB := NewBCB(page_id, vFrameId, 1)
		if b.Ftop[vFrameId] != -1 {
			vBCB := b.Ptof[b.hash(b.Ftop[vFrameId])]
			for vBCB != nil && vBCB.Frame_id != vFrameId {
				vBCB = vBCB.next
			}
			if vBCB == nil {
				fmt.Println("fixPage异常: selectVictim未找到有效的页帧")
				return -1
			}
			b.removeBCB(vBCB, vBCB.page_id)
			b.removeLRUEle(vBCB.Frame_id)
			b.Ftop[vBCB.Frame_id] = -1
		}
		b.Ftop[newBCB.Frame_id] = newBCB.page_id
		temp := b.Ptof[b.hash(newBCB.page_id)]
		if temp == nil {
			b.Ptof[b.hash(newBCB.page_id)] = newBCB
		} else {
			for temp.next != nil {
				temp = temp.next
			}
			temp.next = newBCB
		}
		b.LRU.addLRUEle(newBCB)
		if prot == 0 {
			b.Buf[newBCB.Frame_id] = b.DSMgr.readPage(newBCB.page_id)
		} else {
			b.Buf[newBCB.Frame_id] = &BFrame{Field: [1024]rune{}}
		}
		return newBCB.Frame_id
	} else {

		b.HitCounter++
		p := b.getLRUEle(bcb.Frame_id)
		if p == nil {
			fmt.Println("fixPage异常: 未找到指定的LRU元素")
		} else {
			b.LRU.moveToMru(p)
		}
		bcb.count++
		return bcb.Frame_id
	}

}

// fixNewPage 为新的页面分配一个页面 ID，并将其标记为已使用。
// 如果磁盘已满，返回 -1。
func (b *BMgr) fixNewPage() int {
	// 如果磁盘已满
	if b.DSMgr.getNumPages() == len(b.DSMgr.getPages()) {
		return -1
	}
	for page_id, page := range b.DSMgr.getPages() {
		if page == 0 {
			b.DSMgr.setUse(page_id, MAXPAGES)
			b.DSMgr.incNumPages()
			return page_id
		}
	}
	return -1
}

// unFixPage 减少给定页面 ID 的引用计数，并返回该页面在缓冲区中的帧号。
// 如果页面 ID 不存在，返回 -1。
func (b *BMgr) unFixPage(page_id int) int {
	bcb := b.Ptof[b.hash(page_id)]
	for bcb != nil && bcb.page_id != page_id {
		bcb = bcb.next
	}
	if bcb == nil {
		return -1
	} else {
		bcb.count--
		return bcb.Frame_id
	}
}

// numFreeFrames 返回第一个可用的帧号。
// 如果没有可用的帧，返回 -1。
func (b *BMgr) numFreeFrames() int {
	i := 0
	for i < DEFBUFSIZE && b.Ftop[i] != -1 {
		i++
	}
	if i == DEFBUFSIZE {
		return -1
	} else {
		return i
	}
}

// selectVictim 使用 LRU 策略找到可以被替换的帧号。
// 注意此帧可能为空。
func (b *BMgr) selectVictim() int {
	vFrame_id := b.numFreeFrames()
	if vFrame_id != -1 {
		return vFrame_id
	} else {
		p := b.LRU.lru
		for p.bcb.count != 0 {
			p = p.postLRUEle
		}
		return p.bcb.Frame_id
	}
}

// removeBCB 移除给定的 BCB。
// 如果 BCB 是脏的，它会被写回磁盘。
func (b *BMgr) removeBCB(ptr *BCB, page_id int) {
	bcb := b.Ptof[b.hash(page_id)]
	if bcb == nil {
		return
	}
	if bcb == ptr {
		b.Ptof[b.hash(page_id)] = bcb.next
	} else {
		for bcb.next != nil && bcb.next != ptr {
			bcb = bcb.next
		}
		if bcb.next == nil {
			fmt.Println("removeBCB异常: 未找到指定的BCB")
		}
		bcb.next = ptr.next
	}
	ptr.next = nil
	// 如果是脏页，需要写回
	if ptr.dirty == 1 {
		if b.DSMgr.writePage(page_id, b.Buf[ptr.Frame_id]) != FRAMESIZE {
			fmt.Println("removeBCB异常: 页帧写入不完整")
		}
		b.unSetDirty(ptr.Frame_id)
	}
}

// removeLRUEle 移除给定帧号的 LRU 元素。
func (b *BMgr) removeLRUEle(frame_id int) {
	b.LRU.removeLRUEle(frame_id)
}

// getLRUEle 返回给定帧号的 LRU 元素。
func (b *BMgr) getLRUEle(frame_id int) *LRUEle {
	return b.LRU.getLRUEle(frame_id)
}

// setDirty 将给定帧号的页面标记为脏的。
func (b *BMgr) setDirty(frame_id int) {
	pid := b.Ftop[frame_id]
	fid := b.hash(pid)
	bcb := b.Ptof[fid]
	for bcb != nil && bcb.page_id != pid {
		bcb = bcb.next
	}
	if bcb != nil {
		bcb.dirty = 1
	}
}

// unSetDirty 将给定帧号的页面标记为不脏的。
func (b *BMgr) unSetDirty(frame_id int) {
	pid := b.Ftop[frame_id]
	fid := b.hash(pid)
	bcb := b.Ptof[fid]
	for bcb != nil && bcb.page_id != pid {
		bcb = bcb.next
	}
	if bcb != nil {
		bcb.dirty = 0
	}
}

// writeDirtys 将缓冲区中所有脏的页面写回磁盘。
func (b *BMgr) writeDirtys() {
	for _, bcb := range b.Ptof {
		for bcb != nil {
			if bcb.dirty == 1 {
				if b.DSMgr.writePage(bcb.page_id, b.Buf[bcb.Frame_id]) != FRAMESIZE {
					fmt.Println("writeDirtys异常: 页帧写入不完整")
				}
				b.unSetDirty(bcb.Frame_id)
			}
			bcb = bcb.next
		}
	}
}

// printFrame 打印给定帧号的页面的内容。
func (b *BMgr) printFrame(frame_id int) {
	fmt.Println(b.Buf[frame_id].GetField())
}
