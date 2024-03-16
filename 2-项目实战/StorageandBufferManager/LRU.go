package main

import "fmt"

type LRUEle struct {
	bcb        *BCB
	preLRUEle  *LRUEle
	postLRUEle *LRUEle
}

type LRU struct {
	lru *LRUEle
	mru *LRUEle
}

func NewLRU() *LRU {
	return &LRU{
		lru: nil,
		mru: nil,
	}
}
func (l *LRU) removeLRUEle(frameID int) {
	if l.lru != nil && l.lru.bcb.Frame_id == frameID {
		l.lru = l.lru.postLRUEle
		if l.lru != nil {
			l.lru.preLRUEle = nil
		}
	} else if l.mru != nil && l.mru.bcb.Frame_id == frameID {
		l.mru = l.mru.preLRUEle
		if l.mru != nil {
			l.mru.postLRUEle = nil
		}
	} else {
		p := l.lru
		for p != nil && p.bcb.Frame_id != frameID {
			p = p.postLRUEle
		}
		if p == nil {
			fmt.Println("removeLRUEle异常：未在LRU链表中找到相应的页帧")
		} else {
			p.preLRUEle.postLRUEle = p.postLRUEle
			p.postLRUEle.preLRUEle = p.preLRUEle
		}
	}
}

func (l *LRU) getLRUEle(frameID int) *LRUEle {
	p := l.mru
	for p != nil && p.bcb.Frame_id != frameID {
		p = p.preLRUEle
	}
	if p == nil {
		fmt.Println("getLRUEle异常：未在LRU链表中找到相应的页帧")
	}
	return p
}

func (l *LRU) addLRUEle(bcb *BCB) {
	newLRUEle := &LRUEle{
		bcb: bcb,
	}
	if l.lru == nil && l.mru == nil {
		l.lru = newLRUEle
		l.mru = newLRUEle
	} else {
		l.mru.postLRUEle = newLRUEle
		newLRUEle.preLRUEle = l.mru
		newLRUEle.postLRUEle = nil
		l.mru = newLRUEle
	}
}

func (l *LRU) moveToMru(lruEle *LRUEle) {
	if lruEle.postLRUEle == nil {
		return
	} else if lruEle.preLRUEle == nil {
		l.lru = lruEle.postLRUEle
		if l.lru != nil {
			l.lru.preLRUEle = nil
		}
		lruEle.postLRUEle = nil
		l.mru.postLRUEle = lruEle
		lruEle.preLRUEle = l.mru
		l.mru = lruEle
	} else {
		lruEle.preLRUEle.postLRUEle = lruEle.postLRUEle
		lruEle.postLRUEle.preLRUEle = lruEle.preLRUEle
		l.mru.postLRUEle = lruEle
		lruEle.preLRUEle = l.mru
		lruEle.postLRUEle = nil
		l.mru = lruEle
	}
}
