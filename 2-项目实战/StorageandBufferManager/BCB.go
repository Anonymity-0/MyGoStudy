package main

// BCB 代表一个缓存控制块，它包含了一个页面的元数据。
type BCB struct {
	page_id  int  // 页面的 ID
	Frame_id int  // 页面在内存中的帧 ID
	latch    int  // 锁，用于并发控制
	count    int  // 页面的引用计数
	dirty    int  // 页面是否被修改的标志，0 表示未修改，1 表示已修改
	next     *BCB // 指向下一个 BCB 的指针
}

// NewBCB 创建一个新的 BCB，并初始化它的字段。
// page_id 是页面的 ID，Frame_id 是页面在内存中的帧 ID，count 是页面的引用计数。
func NewBCB(page_id int, Frame_id int, count int) *BCB {
	return &BCB{
		page_id:  page_id,
		Frame_id: Frame_id,
		latch:    0,     // 初始时，锁的状态是 0
		count:    count, // 初始时，页面的引用计数是 count
		dirty:    0,     // 初始时，页面未被修改
		next:     nil,   // 初始时，没有下一个 BCB
	}
}
