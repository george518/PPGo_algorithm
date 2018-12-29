/************************************************************
** @Description: _46_2lrucache
** @Author: george hao
** @Date:   2018-12-20 16:50
** @Last Modified by:  george hao
** @Last Modified time: 2018-12-20 16:50
*************************************************************/
package main

import (
	"container/list"
	"fmt"
)

func main() {

	lc := Constructor(10)

	//s := []string{"LRUCache", "put", "put", "put", "put", "put", "get", "put", "get", "get", "put", "get", "put", "put", "put", "get", "put", "get", "get", "get", "get", "put", "put", "get", "get", "get", "put", "put", "get", "put", "get", "put", "get", "get", "get", "put", "put", "put", "get", "put", "get", "get", "put", "put", "get", "put", "put", "put", "put", "get", "put", "put", "get", "put", "put", "get", "put", "put", "put", "put", "put", "get", "put", "put", "get", "put", "get", "get", "get", "put", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "get", "get", "get", "put", "put", "put", "get", "put", "put", "put", "get", "put", "put", "put", "get", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "put", "put", "put"}
	//nums := [][]int{[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]}
	//[10,13],[3,17],[6,11],[10,5],[9,10]

	lc.Put(10, 13)
	lc.Put(3, 17)
	lc.Put(6, 11)
	lc.Put(10, 5)
	lc.Put(9, 10)

	lc.Get(13)

	lc.Put(2, 19)

	lc.Get(2)
	lc.Get(3)

	lc.Put(5, 25)

	lc.Get(8)

	lc.Put(9, 22)
	lc.Put(5, 5)
	lc.Put(1, 30)

	lc.Get(11)

	lc.Put(9, 12)

	lc.Get(7)
	lc.Get(5)
	lc.Get(8)
	lc.Get(9)

	lc.Put(4, 30)
	lc.Put(9, 3)

	lc.Get(9)
	lc.Get(10)
	lc.Get(10)
	lc.Put(6, 14)
	lc.Put(3, 1)
	lc.Get(3)
	lc.Put(10, 11)
	lc.Get(8)
	lc.Put(2, 14)
	lc.Get(1)
	lc.Get(5)
	lc.Get(4)
	lc.Put(11, 4)
	lc.Put(12, 24)
	lc.Put(5, 18)
	lc.Get(13)
	lc.Put(7, 23)
	lc.Get(8)
	lc.Get(12)
	lc.Put(3, 27)
	lc.Put(2, 12)
	lc.Get(5)
	lc.Put(2, 9)
	lc.Put(13, 4)
	lc.Put(8, 18)
	lc.Put(1, 7)
	lc.Get(6)
	lc.Put(9, 29)
	lc.Put(8, 21)
	lc.Get(5)
	lc.Put(6, 30)
	lc.Put(1, 12)
	lc.Get(10)
	lc.Put(4, 15)
	lc.Put(7, 22)
	lc.Put(11, 26)
	lc.Put(8, 17)
	lc.Put(9, 29)
	lc.Get(5)
	lc.Put(3, 4)
	lc.Put(11, 30)
	lc.Get(12)
	lc.Put(4, 29)
	lc.Get(3)
	lc.Get(9)
	lc.Get(6)
	lc.Put(3, 4)
	lc.Get(1)
	lc.Get(10)
	lc.Put(3, 29)
	lc.Put(10, 28)
	lc.Put(1, 20)
	lc.Put(11, 13)
	lc.Get(3)
	lc.Put(3, 12)
	lc.Put(3, 8)
	lc.Put(10, 9)
	lc.Put(3, 26)
	fmt.Println(lc.Get(8),
		lc.Get(7),
		lc.Get(5))

	fmt.Println()
	lc.Put(13, 17)
	lc.Put(2, 27)
	lc.Put(11, 15)
	lc.Get(12)
	lc.Put(9, 19)
	lc.Put(2, 15)
	lc.Put(3, 16)
	lc.Get(1)
	lc.Put(12, 17)
	lc.Put(9, 1)
	lc.Put(6, 19)
	lc.Get(4)
	lc.Get(5)
	lc.Get(5)
	lc.Put(8, 1)
	lc.Put(11, 7)
	lc.Put(5, 2)
	lc.Put(9, 28)
	lc.Get(1)
	lc.Put(2, 2)
	lc.Put(7, 4)
	lc.Put(4, 22)
	lc.Put(7, 24)
	lc.Put(9, 26)
	lc.Put(13, 28)
	lc.Put(11, 26)

}

type CacheNode struct {
	Key, Value int
}

type LRUCache struct {
	Capacity  int
	CacheList *list.List
	cacheMap  map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	lc := new(LRUCache)
	lc.Capacity = capacity
	lc.CacheList = list.New()
	lc.cacheMap = make(map[int]*list.Element)
	return *lc
}

func (this *LRUCache) Get(key int) int {
	if this.cacheMap == nil {
		return -1
	}

	if pElem, ok := this.cacheMap[key]; ok {
		this.CacheList.MoveToFront(pElem)
		return pElem.Value.(*CacheNode).Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.CacheList == nil {
		return
	}

	if pElem, ok := this.cacheMap[key]; ok {
		this.CacheList.MoveToFront(pElem)
		pElem.Value.(*CacheNode).Value = value
		return
	}

	newElem := this.CacheList.PushFront(&CacheNode{key, value})
	this.cacheMap[key] = newElem

	if this.CacheList.Len() > this.Capacity {
		lastElem := this.CacheList.Back()

		if lastElem == nil {
			return
		}

		cacheNode := lastElem.Value.(*CacheNode)
		delete(this.cacheMap, cacheNode.Key)
		this.CacheList.Remove(lastElem)
	}

	return

}
