package openhash

import "github.com/ami-GS/AlgorithmNotes/map/utils"

type node struct {
	key int
	val int
	nxt *node
}

type Map struct {
	list []*node
}

func NewMap(sizeHint int) *Map {
	return &Map{
		list: make([]*node,
			int(utils.GetClosestPrime(uint32(sizeHint)))),
	}
}

func (m *Map) Insert(key, val int) {
	idx := key % len(m.list)

	if m.list[idx] == nil {
		m.list[idx] = &node{
			key: key,
			val: val,
			nxt: nil,
		}
	} else {
		np := m.list[idx]
		prv := np
		for np != nil && np.key != key {
			prv = np
			np = np.nxt
		}
		if np == nil {
			prv.nxt = &node{
				key: key,
				val: val,
				nxt: nil,
			}
		} else {
			np.val = val
		}
	}
}

func (m *Map) Remove(key int) {
	idx := key % len(m.list)
	np := m.list[idx]
	dummy := node{key: 0, val: 0, nxt: np}
	dummyP := &dummy
	for dummyP != nil {
		if dummyP.nxt != nil && dummyP.nxt.key == key {
			dummyP.nxt = dummyP.nxt.nxt
		} else {
			dummyP = dummyP.nxt
		}
	}
	m.list[idx] = dummy.nxt
}

func (m *Map) Get(key int) int {
	idx := key % len(m.list)
	np := m.list[idx]
	for ; np != nil && np.key != key; np = np.nxt {
	}
	if np == nil {
		return -1
	}
	return np.val
}
