package main

import (
	"container/heap"
	"log"
)

type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func do1() {
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.update(item, item.value, 5)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		log.Printf("%.2d:%s ", item.priority, item.value)
	}
}

type myHeap []int

/* 实现排序 */
func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// 最小堆实现
func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

/* 实现往堆中添加元素 */
func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

/* 实现删除堆中元素 */
func (h *myHeap) Pop() (v interface{}) {
	*h, v = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return
}

// 按层来遍历和打印堆数据，第一行只有一个元素，即堆顶元素
func (h myHeap) printHeap() {
	n := 1
	levelCount := 1
	for n <= h.Len() {
		log.Println(h[n-1 : n-1+levelCount])
		n += levelCount
		levelCount *= 2
	}
}

func do2() {
	data := [7]int{13, 12, 45, 23, 11, 9, 20}
	aHeap := new(myHeap)
	for i := 0; i < len(data); i++ {
		aHeap.Push(data[i])
	}
	aHeap.printHeap()

	// 堆排序处理
	heap.Init(aHeap)
	aHeap.printHeap()
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do2()
}
