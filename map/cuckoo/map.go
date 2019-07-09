package cuckoo

import (
	"fmt"

	"github.com/ami-GS/AlgorithmNotes/map/utils"
)

type elementFlag byte

const (
	NO_VALUE elementFlag = iota
	HAS_VALUE
	INVALID
)

const MAX_LOOP = 128

type node struct {
	key  int
	val  int
	flag elementFlag
}

type Map struct {
	lists [2][]node
}

func NewMap(sizeHint int) *Map {
	// TODO: half
	return &Map{
		lists: [2][]node{
			make([]node, int(utils.GetClosestPrime(uint32(sizeHint)))),
			make([]node, int(utils.GetClosestPrime(uint32(sizeHint)))), // TODO: next prime
		},
	}
}

func (m *Map) Insert(key, val int) {
	origKey := key // because key is swapped
	for i := 0; i < MAX_LOOP; i++ {
		for j := 0; j < 2; j++ {
			list := m.lists[j]
			idx := (origKey + i) % len(list)

			if list[idx].flag == NO_VALUE || list[idx].flag == INVALID {
				list[idx].key = key
				list[idx].val = val
				list[idx].flag = HAS_VALUE
				return
			} else if list[idx].key == key {
				list[idx].val = val
				return
			}
			// swap
			key, val, list[idx].key, list[idx].val = list[idx].key, list[idx].val, key, val
		}
	}
	m.rehash()
	m.Insert(key, val)
}

func (m *Map) Remove(key int) {
	for i := 0; i < MAX_LOOP; i++ {
		for j := 0; j < 2; j++ {
			list := m.lists[j]
			idx := (key + i) % len(list)

			if list[idx].flag == NO_VALUE {
				panic("invalid key")
			} else if list[idx].key == key {
				list[idx].flag = INVALID
				return
			}
		}
	}
	// TODO rehash
	panic("rehash is required")
}

func (m *Map) Get(key int) int {
	for i := 0; i < MAX_LOOP; i++ {
		for j := 0; j < 2; j++ {
			list := m.lists[j]
			idx := (key + i) % len(list)

			if list[idx].flag == NO_VALUE {
				panic("invalid key")
			} else if list[idx].key == key {
				return list[idx].val
			}
		}
	}
	// TODO rehash
	panic("rehash is required")
}

func (m *Map) rehash() {
	// TODO: closer prime
	list1 := m.lists[0]
	list2 := m.lists[1]
	m.lists[0] = make([]node, len(m.lists[0])*2)
	m.lists[1] = make([]node, len(m.lists[1])*2)

	fmt.Println(list1)
	fmt.Println(list2)
	for _, v := range list1 {
		if v.flag != NO_VALUE {
			m.Insert(v.key, v.val)
		}
	}
	for _, v := range list2 {
		if v.flag != NO_VALUE {
			m.Insert(v.key, v.val)
		}
	}
}
