package closedhash

import (
	"github.com/ami-GS/AlgorithmNotes/map/utils"
)

type elementFlag byte

const (
	NO_VALUE elementFlag = iota
	HAS_VALUE
	INVALID
)

type node struct {
	key  int
	val  int
	flag elementFlag
}

type Map struct {
	list []node
}

func NewMap(sizeHint int) *Map {
	return &Map{
		list: make([]node,
			int(utils.GetClosestPrime(uint32(sizeHint)))),
	}
}

func (m *Map) Insert(key, val int) {
	idx := key % len(m.list)

	// or rehash
	// one skip by idx++
	for m.list[idx].flag == HAS_VALUE && m.list[idx].key != key {
		idx = (idx + 1) % len(m.list)
	}

	m.list[idx].val = val
	m.list[idx].key = key
	m.list[idx].flag = HAS_VALUE
}

func (m Map) findKey(key int) int {
	idx := key % len(m.list)
	idxMinusOne := idx - 1
	if idxMinusOne == -1 {
		idxMinusOne = len(m.list) - 1
	}

	for m.list[idx].flag == HAS_VALUE || m.list[idx].flag == INVALID {
		if m.list[idx].key == key {
			break
		} else if idx == idxMinusOne {
			return -1
		}
		idx = (idx + 1) % len(m.list)
	}
	return idx
}

func (m *Map) Remove(key int) {
	idx := m.findKey(key)
	if idx == -1 {
		panic("no entry found")
	}
	m.list[idx].flag = INVALID
}

func (m *Map) Get(key int) int {
	idx := m.findKey(key)
	if idx == -1 {
		panic("no entry found")
	}

	return m.list[idx].val
}
