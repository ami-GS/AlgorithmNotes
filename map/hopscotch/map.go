package hopscotch

import (
	"math"

	"github.com/ami-GS/AlgorithmNotes/map/utils"
)

type hopInfo uint64

func (h hopInfo) swap(i, j int) {
	if j == i {
		panic("swap should be on different bit")
	}
	iBit := (uint64(h) & (1 << uint64(i))) >> uint64(i)
	jBit := (uint64(h) & (1 << uint64(j))) >> uint64(j)
	if iBit+jBit != 1 {
		panic("swap should be between 0 and 1")
	}

	h = hopInfo(uint64(h) ^ (1 << uint64(i)))
	h = hopInfo(uint64(h) ^ (1 << uint64(j)))
}

func (h hopInfo) checkInsertability() bool {
	if h == math.MaxUint64 {
		return false
	}
	return true
}

func (h hopInfo) getAvailableLeftBitNextTo(loc int) byte {
	// bad implementation
	for i := uint64(1 << uint(63-loc)); i > 0; i >>= 1 {
		if uint64(h)&i == 0 {
			return byte(i)
		}
	}
	return 255 // not availabe, rehash would be required
}

func (h hopInfo) getSelfAvailability() bool {
	return (uint64(h) & (1 << 63)) == 0
}

func (h *hopInfo) setBit(i int) {
	*h |= (1 << uint64(i))
}

type node struct {
	elememt int
	info    hopInfo
}

type Map struct {
	list []node
}

func NewMap(sizeHint int) *Map {
	return &Map{
		list: make([]node, int(utils.GetClosestPrime(uint32(sizeHint)))),
	}
}

func (m *Map) Insert(key, val int) {
	idx := key % len(m.list)
	if m.list[idx].hopInfo.getSelfAvailability() {
		m.list[idx].element = val
		m.list[idx].setBit(63)
	}

}

func (m *Map) swap(i, j, k int) {
	m.list[i].info.swap(i, j)

	// update all hopInfos which are affected by above swap
	for {

	}
}
