package main

import (
	"bytes"
	"encoding/binary"
	"hash/fnv"
	"log"
)

type Elem struct {
	Key   int
	Value interface{}
}

type HashItem struct {
	Dist uint64
	Elem *Elem
}

type HopScotchHashTable struct {
	arr     []*HashItem
	maxDist int
	sz      int
}

func NewHopScotchHashTable(size, maxDist int) *HopScotchHashTable {
	arr := make([]*HashItem, size)
	for i := 0; i < len(arr); i++ {
		arr[i] = &HashItem{}
	}

	if maxDist > 64 {
		maxDist = 64
	}

	return &HopScotchHashTable{
		arr:     arr,
		maxDist: maxDist,
	}
}

func (h *HopScotchHashTable) Print() {
	for i := 0; i < len(h.arr); i++ {
		log.Println(*h.arr[i])
	}
}

func (h *HopScotchHashTable) findPos(key int) int {
	pos := h.hasher(key)
	for i := 0; i < h.maxDist; i++ {
		if (h.arr[pos].Dist>>i)%2 == 1 {
			if h.arr[pos+h.maxDist-1-i].Elem.Key == key {
				return pos + h.maxDist - 1 - i
			}
		}
	}
	return -1
}

func (h *HopScotchHashTable) Get(key int) (interface{}, bool) {
	pos := h.findPos(key)
	if pos == -1 {
		return nil, false
	}
	return h.arr[pos].Elem.Value, true
}

func (h *HopScotchHashTable) Delete(key int) {
	pos := h.findPos(key)
	if pos != -1 {
		hash := h.hasher(key)
		h.arr[pos].Elem = nil
		h.arr[hash].Dist = h.arr[hash].Dist - (1 << (h.maxDist - 1 + hash - pos))
	}
}

func (h *HopScotchHashTable) Set(key int, value interface{}) {
	for {
		pos := h.hasher(key)
		temp := pos
		for pos < len(h.arr) && h.arr[pos].Elem != nil {
			pos++
		}

		// TODO
		if pos >= len(h.arr) {
			h.rehash()
			continue
		}

		if pos <= temp+h.maxDist-1 {
			h.arr[pos].Elem = &Elem{Key: key, Value: value}
			h.arr[temp].Dist = h.arr[temp].Dist + (1 << (h.maxDist - 1 + temp - pos)) // 1 << (h.maxDist - 1 + 领主位置 - 领子位置)
			h.sz++
			log.Println("insert1:", key, value)
			return
		}

		for {
			isNotDist := false

			for i := h.maxDist - 1; i > 0; i-- {
				for j := h.maxDist - 1; j > h.maxDist-1-i; j-- {
					if (h.arr[pos-i].Dist>>j)%2 == 1 {
						item := h.arr[pos-i+h.maxDist-1-j]
						h.arr[pos].Elem = item.Elem
						item.Elem = nil
						// 领主位置: pos-i
						// 旧位置: pos-i+h.maxDist-1-j
						// 新位置: pos
						// 从领域摘除,再重新设置新位置
						h.arr[pos-i].Dist = h.arr[pos-i].Dist - (1 << j) + (1 << (h.maxDist - 1 - i))

						// pos新位置,相当于pos向上移动
						pos = pos - i + h.maxDist - 1 - j

						if pos <= temp+3 {
							h.arr[pos].Elem = &Elem{Key: key, Value: value}
							h.arr[temp].Dist = h.arr[temp].Dist + (1 << (temp + h.maxDist - 1 - pos))
							h.sz++
							log.Println("insert2:", key, value)
							return
						} else {
							isNotDist = true
							break
						}
					}
				}

				if isNotDist {
					break
				}
			}

			if !isNotDist {
				break
			}
		}

		h.rehash()
		//log.Println("insert fail, value:", value)
		//return
	}
}

func (h *HopScotchHashTable) rehash() {
	oldArr := h.arr
	newArr := make([]*HashItem, 2*len(oldArr))
	for i := 0; i < len(newArr); i++ {
		newArr[i] = &HashItem{}
	}
	h.sz = 0
	h.arr = newArr
	for i := 0; i < len(oldArr); i++ {
		if elem := oldArr[i].Elem; elem != nil {
			h.Set(elem.Key, elem.Value)
		}
	}
	log.Println("rehash")
}

func (h *HopScotchHashTable) hasher(key int) int {
	buf := bytes.NewBuffer(nil)
	if err := binary.Write(buf, binary.LittleEndian, uint64(key)); err != nil {
		log.Fatal(err)
	}
	f := fnv.New64()
	if _, err := f.Write(buf.Bytes()); err != nil {
		log.Fatal(err)
	}
	total := len(h.arr)
	hashCode := f.Sum64() % uint64(total)
	log.Println(int(hashCode))
	return int(hashCode)
}

func do1() {
	h := NewHopScotchHashTable(8, 4)
	h.Set(1, 23)
	h.Set(2, 23)
	h.Set(10, 23)
	h.Set(3, 23)
	h.Set(4, 23)
	h.Set(9, 23)
	h.Print()
	log.Println(h.findPos(3))
}

func do2() {
	h := NewHopScotchHashTable(8, 4)
	h.Set(1, 23)
	h.Set(2, 23)
	h.Set(10, 23)
	//h.Set(3, 23)
	//h.Set(4, 23)
	//h.Set(9, 23)
	h.Print()
	//log.Println(h.findPos(3))
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do2()
}
