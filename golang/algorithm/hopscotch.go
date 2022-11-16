package main

import (
	"bytes"
	"encoding/binary"
	"hash/fnv"
	"log"
	"math/rand"
	"time"
)

type Elem struct {
	Key   int
	Value interface{}
}

type HashEntry struct {
	Dist uint64
	Elem *Elem
}

type HopScotchHashTable struct {
	arr     []*HashEntry
	maxDist int
	sz      int
}

func NewHopScotchHashTable(size, maxDist int) *HopScotchHashTable {
	arr := make([]*HashEntry, size)
	for i := 0; i < len(arr); i++ {
		arr[i] = &HashEntry{}
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
		if h.arr[i].Elem != nil {
			log.Println(h.arr[i].Dist, *h.arr[i].Elem)
		} else {
			log.Println(h.arr[i].Dist, "nil")
		}
	}
}

func (h *HopScotchHashTable) Range(f func(key int, value interface{})) {
	for i := 0; i < len(h.arr); i++ {
		if h.arr[i].Elem != nil {
			f(h.arr[i].Elem.Key, h.arr[i].Elem.Value)
		}
	}
}

func (h *HopScotchHashTable) findPos(key int) int {
	pos := h.hasher(key)
	for i := 0; i < h.maxDist; i++ {
		if (h.arr[pos].Dist>>i)%2 == 1 {
			if h.arr[(pos+h.maxDist-1-i)%len(h.arr)].Elem.Key == key {
				//if h.arr[pos+h.maxDist-1-i].Elem.Key == key {
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
	return h.arr[pos%len(h.arr)].Elem.Value, true
	//return h.arr[pos].Elem.Value, true
}

func (h *HopScotchHashTable) Delete(key int) {
	pos := h.findPos(key)
	if pos != -1 {
		hash := h.hasher(key)
		h.arr[pos%len(h.arr)].Elem = nil
		h.arr[hash].Dist = h.arr[hash].Dist - (1 << (h.maxDist - 1 + hash - pos))
	}
}

func (h *HopScotchHashTable) Set(key int, value interface{}) {
	if h.sz >= len(h.arr) {
		h.rehash()
	}
	h.set(key, value)
}

func (h *HopScotchHashTable) set2(startPos, pos, key int, value interface{}) bool {
	realPos := pos
	n := len(h.arr)
	pos = pos + n

	if pos <= startPos+h.maxDist-1 {
		h.arr[realPos].Elem = &Elem{Key: key, Value: value}
		h.arr[startPos].Dist = h.arr[startPos].Dist + (1 << (h.maxDist - 1 + startPos - pos)) // 1 << (h.maxDist - 1 + 领主位置 - 领子位置)
		h.sz++
		//log.Println("insert3:", key, value)
		return true
	}

	for {
		isNotDist := false

		for i := h.maxDist - 1; i > 0; i-- {
			for j := h.maxDist - 1; j > h.maxDist-1-i; j-- {
				if (h.arr[(pos-i)%n].Dist>>j)%2 == 1 {
					item := h.arr[(pos-i+h.maxDist-1-j)%n]
					h.arr[realPos].Elem = item.Elem
					item.Elem = nil
					// 领主位置: pos-i
					// 旧位置: pos-i+h.maxDist-1-j
					// 新位置: pos
					// 从领域摘除,再重新设置新位置
					h.arr[(pos-i)%n].Dist = h.arr[(pos-i)%n].Dist - (1 << j) + (1 << (h.maxDist - 1 - i))

					// pos新位置,相当于pos向上移动
					pos = pos - i + h.maxDist - 1 - j
					realPos = pos % n

					log.Printf("key: %d, value: %d, startPos: %d, realPos: %d, pos: %d, cap: %d", key, value, startPos, realPos, pos, len(h.arr))

					if pos <= startPos+h.maxDist-1 {
						h.arr[realPos].Elem = &Elem{Key: key, Value: value}
						h.arr[startPos].Dist = h.arr[startPos].Dist + (1 << (h.maxDist - 1 + startPos - pos))
						h.sz++
						log.Println("insert4:", key, value)
						return true
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

	return false
}

func (h *HopScotchHashTable) set(key int, value interface{}) {
	for {
		pos := h.hasher(key)
		startPos := pos
		flag := false
		for h.arr[pos].Elem != nil {
			pos++

			if pos >= len(h.arr) {
				pos = 0
				flag = true
			}
		}

		if flag {
			if h.set2(startPos, pos, key, value) {
				return
			}
			h.rehash()
			continue
		}

		if pos <= startPos+h.maxDist-1 {
			h.arr[pos].Elem = &Elem{Key: key, Value: value}
			h.arr[startPos].Dist = h.arr[startPos].Dist + (1 << (h.maxDist - 1 + startPos - pos)) // 1 << (h.maxDist - 1 + 领主位置 - 领子位置)
			h.sz++
			//log.Println("insert1:", key, value)
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

						if pos <= startPos+h.maxDist-1 {
							h.arr[pos].Elem = &Elem{Key: key, Value: value}
							h.arr[startPos].Dist = h.arr[startPos].Dist + (1 << (h.maxDist - 1 + startPos - pos))
							h.sz++
							//log.Println("insert2:", key, value)
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
	}
}

func (h *HopScotchHashTable) rehash() {
	oldArr := h.arr
	newArr := make([]*HashEntry, 2*len(oldArr))
	for i := 0; i < len(newArr); i++ {
		newArr[i] = &HashEntry{}
	}
	h.sz = 0
	h.arr = newArr

	//log.Printf("rehash start, cap: %d", len(h.arr))
	for i := 0; i < len(oldArr); i++ {
		if elem := oldArr[i].Elem; elem != nil {
			h.set(elem.Key, elem.Value)
		}
	}
	//log.Printf("rehash end, cap: %d", len(h.arr))
}

func (h *HopScotchHashTable) hasher(key int) int {
	return key % len(h.arr)
}

func (h *HopScotchHashTable) hasher2(key int) int {
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
	//log.Printf("key: %d, hashcode: %d", key, int(hashCode))

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

	n := 10000
	for i := 0; i < n; i++ {
		h.Set(i, i+1000)
	}

	for i := 0; i < n; i++ {
		v, ok := h.Get(i)
		if !ok {
			log.Fatalf("key: %d", i)
		}

		if value, ok2 := v.(int); !ok2 || value != i+1000 {
			log.Fatalf("key: %d", i)
		}
	}
}

func do3() {
	h := NewHopScotchHashTable(16, 4)

	var keys []int
	count := 13
	for i := 0; i < count; i++ {
		h.Set(i, i+1000)
		keys = append(keys, i)
	}
	h.Set(14, 14+1000)
	h.Set(15, 15+1000)
	h.Set(31, 31+1000) // insert4
	keys = append(keys, 14)
	keys = append(keys, 15)
	keys = append(keys, 31)

	for i := 0; i < len(keys); i++ {
		v, ok := h.Get(keys[i])
		if !ok {
			log.Fatalf("key %d not existed", keys[i])
		}

		if value, ok2 := v.(int); !ok2 || value != keys[i]+1000 {
			log.Fatalf("key %d value: %d error", keys[i], value)
		}
	}

	h.Range(func(key int, value interface{}) {
		val, ok := value.(int)
		if !ok {
			log.Fatal("type error")
		}
		if val != key+1000 {
			log.Fatal("value error")
		}
	})

	h.Print()
}

func do4() {
	h := NewHopScotchHashTable(64, 8)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	var keys []int
	n := 100000
	for i := 0; i < n; i++ {
		k := r.Int()
		h.Set(k, k+1000)
		keys = append(keys, k)
	}

	for i := 0; i < len(keys); i++ {
		v, ok := h.Get(keys[i])
		if !ok {
			log.Fatalf("key %d not existed", keys[i])
		}

		if value, ok2 := v.(int); !ok2 || value != keys[i]+1000 {
			log.Fatalf("key %d value: %d error", keys[i], value)
		}
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do3()
}
