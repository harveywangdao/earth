package main

import (
	"bytes"
	"encoding/binary"
	"hash/fnv"
	"log"
	"math/rand"
	"time"
)

type Entry struct {
	Key    int
	Value  interface{}
	Offset int
}

type RobinHoodHashTable struct {
	table      []*Entry
	size       int
	loadFactor float64
}

func NewRobinHoodHashTable(size int, loadFactor float64) *RobinHoodHashTable {
	arr := make([]*Entry, size)
	return &RobinHoodHashTable{
		table:      arr,
		size:       0,
		loadFactor: loadFactor,
	}
}

func (r *RobinHoodHashTable) Set(key int, value interface{}) {
	r.set(key, value)
}

func (r *RobinHoodHashTable) set(key int, value interface{}) {
	entry := &Entry{Key: key, Value: value}
	idx := r.hasher(key)

	for r.table[idx] != nil {
		if entry.Offset > r.table[idx].Offset {
			temp := r.table[idx]
			r.table[idx] = entry
			entry = temp
			idx = r.increment(idx)
			entry.Offset++
		} else if entry.Offset == r.table[idx].Offset {
			if entry.Key == r.table[idx].Key {
				r.table[idx].Value = entry.Value
				return
			} else {
				idx = r.increment(idx)
				entry.Offset++
			}
		} else {
			idx = r.increment(idx)
			entry.Offset++
		}
	}

	r.table[idx] = entry
	r.size++

	if float64(r.size) >= float64(len(r.table))*r.loadFactor {
		r.rehash(2 * len(r.table))
	}
}

func (r *RobinHoodHashTable) Print() {
	for _, e := range r.table {
		if e != nil {
			log.Println(*e)
		} else {
			log.Println("nil")
		}
	}
}

func (r *RobinHoodHashTable) rehash(newCap int) {
	oldTable := r.table
	newTable := make([]*Entry, newCap)
	r.size = 0
	r.table = newTable
	log.Printf("rehash start, cap: %d", newCap)
	for _, e := range oldTable {
		if e != nil {
			r.set(e.Key, e.Value)
		}
	}
}

func (r *RobinHoodHashTable) increment(idx int) int {
	idx++
	if idx >= len(r.table) {
		return 0
	}
	return idx
}

func (r *RobinHoodHashTable) decrement(idx int) int {
	idx--
	if idx < 0 {
		return len(r.table) - 1
	}
	return idx
}

func (r *RobinHoodHashTable) Get(key int) (interface{}, bool) {
	offset := 0
	idx := r.hasher(key)

	for r.table[idx] != nil {
		if offset > r.table[idx].Offset {
			return nil, false
		} else if offset == r.table[idx].Offset {
			if r.table[idx].Key == key {
				return r.table[idx].Value, true
			} else {
				offset++
				idx = r.increment(idx)
			}
		} else {
			offset++
			idx = r.increment(idx)
		}
	}

	return nil, false
}

func (r *RobinHoodHashTable) Delete(key int) {
	offset := 0
	idx := r.hasher(key)

	for r.table[idx] != nil {
		if offset > r.table[idx].Offset {
			return
		} else if offset == r.table[idx].Offset {
			if r.table[idx].Key == key {
				r.table[idx] = nil
				r.size--
				idx = r.increment(idx)
				for r.table[idx] != nil && r.table[idx].Offset > 0 {
					temp := r.table[idx]
					temp.Offset--
					r.table[r.decrement(idx)] = temp
					r.table[idx] = nil
					idx = r.increment(idx)
				}
				return
			} else {
				offset++
				idx = r.increment(idx)
			}
		} else {
			offset++
			idx = r.increment(idx)
		}
	}
}

func (r *RobinHoodHashTable) hasher(key int) int {
	return key % len(r.table)
}

func (r *RobinHoodHashTable) hasher2(key int) int {
	buf := bytes.NewBuffer(nil)
	if err := binary.Write(buf, binary.LittleEndian, uint64(key)); err != nil {
		log.Fatal(err)
	}
	f := fnv.New64()
	if _, err := f.Write(buf.Bytes()); err != nil {
		log.Fatal(err)
	}
	total := len(r.table)
	hashCode := f.Sum64() % uint64(total)
	//log.Printf("key: %d, hashcode: %d", key, int(hashCode))
	return int(hashCode)
}

// 冲突测试
func do1() {
	h := NewRobinHoodHashTable(16, 0.5)
	h.Set(1, 1000)
	h.Set(1+16, 1000)
	h.Set(1+2*16, 1000)
	h.Set(1+3*16, 1000)
	h.Set(1+4*16, 1000)
	h.Set(1+5*16, 1000)
	h.Set(1+6*16, 1000)
	h.Print()
}

// 交换测试
func do2() {
	h := NewRobinHoodHashTable(16, 0.5)
	h.Set(1, 1000)
	h.Set(2, 1000)
	h.Set(1+16, 1000)
	h.Set(1+2*16, 1000)
	h.Set(1+3*16, 1000)
	h.Set(1+4*16, 1000)
	h.Set(2+16, 1000)
	h.Print()
}

// 相同key
func do3() {
	h := NewRobinHoodHashTable(16, 0.5)
	h.Set(1, 1000)
	h.Set(1, 2000)
	h.Print()
}

// 删除key
func do4() {
	h := NewRobinHoodHashTable(16, 0.5)

	h.Set(1, 1000)
	h.Set(2, 1000)
	h.Set(1+16, 1000)
	h.Set(1+2*16, 1000)
	h.Set(1+3*16, 1000)
	h.Set(1+4*16, 1000)
	h.Set(2+16, 1000)
	h.Print()

	log.Println()
	h.Delete(33)
	h.Print()

	return
}

// 递增key
func do5() {
	h := NewRobinHoodHashTable(8, 0.5)

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

// 随机key
func do6() {
	h := NewRobinHoodHashTable(8, 0.5)
	r := rand.New(rand.NewSource(time.Now().Unix()))

	var keys []int
	n := 100000
	start := time.Now()
	for i := 0; i < n; i++ {
		k := r.Int()
		h.Set(k, k+1000)
		keys = append(keys, k)
	}
	log.Println("set cost:", time.Now().Sub(start))

	start = time.Now()
	for i := 0; i < len(keys); i++ {
		v, ok := h.Get(keys[i])
		if !ok {
			log.Fatalf("key %d not existed", keys[i])
		}

		if value, ok2 := v.(int); !ok2 || value != keys[i]+1000 {
			log.Fatalf("key %d value: %d error", keys[i], value)
		}
	}
	log.Println("get cost:", time.Now().Sub(start))
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do4()
}
