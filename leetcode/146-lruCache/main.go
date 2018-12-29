/************************************************************
** @Description: lruCache
** @Author: george hao
** @Date:   2018-12-19 16:43
** @Last Modified by:  george hao
** @Last Modified time: 2018-12-19 16:43
*************************************************************/
package main

import "fmt"

//146
//运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
//
//获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
//写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。
//
//进阶:
//
//你是否可以在 O(1) 时间复杂度内完成这两种操作？

func main() {
	lc := Constructor(2)

	lc.Put(1, 1)
	lc.Put(2, 2)

	//fmt.Println(lc, lc.Head.Key, lc.Tail.Key)
	lc.Get(1)
	//fmt.Println(lc, lc.Head.Key, lc.Tail.Key)~
	lc.Put(3, 3)
	lc.Get(2)
	lc.Put(4, 4)

	fmt.Println(lc)
	lc.Get(1)

	lc.Get(3)
	lc.Get(4)
}

type CacheNode struct {
	Prev  *CacheNode
	Next  *CacheNode
	Key   int
	Value int
}

type LRUCache struct {
	Indexes  map[int]*CacheNode
	Capacity int //容量
	Length   int //长度
	Head     *CacheNode
	Tail     *CacheNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		Indexes:  make(map[int]*CacheNode, capacity),
	}
}

func (this *LRUCache) Get(key int) int {

	//查找map里是否有
	if node, ok := this.Indexes[key]; ok {
		//直接放到双链表队首
		this.MoveToHead(node)
		return node.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {

	node, ok := this.Indexes[key]
	if ok {
		//更新并放到队首
		node.Value = value
		this.MoveToHead(node)

	} else {
		if this.Capacity == this.Length {
			//需要删除最后一个节点
			tailKey := this.Tail.Key
			delete(this.Indexes, tailKey)
			this.RemoveTail()
		} else {
			this.Length += 1
		}
		node = new(CacheNode)
		node.Key = key
		node.Value = value
		this.Indexes[key] = node
		this.InsertToHead(node)
	}

}

//替换头部节点
func (this *LRUCache) MoveToHead(node *CacheNode) {

	switch node {
	case this.Head:
		return
	case this.Tail:
		this.RemoveTail()
	default:
		//成全左右两个节点
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	}

	this.InsertToHead(node)
}

func (this *LRUCache) RemoveTail() {

	if this.Tail.Prev != nil {
		this.Tail.Prev.Next = nil
	} else {
		this.Head = nil
	}

	this.Tail = this.Tail.Prev
}

func (this *LRUCache) InsertToHead(node *CacheNode) {

	if this.Tail == nil {
		this.Tail = node
	} else {
		node.Next = this.Head
		this.Head.Prev = node
	}

	this.Head = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
