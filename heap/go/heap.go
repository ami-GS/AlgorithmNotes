package heap

type HeapType bool

const MaxSize = 1024

const (
	Max = HeapType(false)
	Min = HeapType(true)
)

type Heap struct {
	array     []int
	size      int
	direction HeapType
}

func NewHeap(heapType HeapType) *Heap {
	return &Heap{
		array:     make([]int, MaxSize),
		size:      0,
		direction: heapType,
	}
}

func NewFromSlice(data []int, heapType HeapType) *Heap {
	h := &Heap{
		array:     make([]int, MaxSize),
		size:      0,
		direction: heapType,
	}
	for _, item := range data {
		h.Push(item)
	}
	return h
}

func (h *Heap) Top() int {
	return h.array[0]
}

func (h *Heap) Pop() int {
	out := h.array[0]
	h.array[0] = h.array[h.size-1]
	if h.direction == Min {
		for i := 0; 2*i+1 < h.size-1 && (h.array[i] > h.array[2*i+1] || h.array[i] > h.array[2*i+2]); {
			if h.array[2*i+1] < h.array[2*i+2] {
				h.array[i], h.array[2*i+1] = h.array[2*i+1], h.array[i]
				i = 2*i + 1
			} else {
				h.array[i], h.array[2*i+2] = h.array[2*i+2], h.array[i]
				i = 2*i + 2
			}
		}
	} else if h.direction == Max {
		for i := 0; 2*i+1 < h.size-1 && (h.array[i] < h.array[2*i+1] || h.array[i] < h.array[2*i+2]); {
			if h.array[2*i+1] < h.array[2*i+2] {
				h.array[i], h.array[2*i+2] = h.array[2*i+2], h.array[i]
				i = 2*i + 2
			} else {
				h.array[i], h.array[2*i+1] = h.array[2*i+1], h.array[i]
				i = 2*i + 1
			}
		}
	}
	h.size--
	return out
}

func (h *Heap) Push(item int) {
	if h.size == MaxSize {
		panic("out of memory")
	}
	h.size++
	h.array[h.size-1] = item

	if h.direction == Min {
		for i := h.size - 1; h.array[i] < h.array[(i-1)/2]; i = (i - 1) / 2 {
			h.array[i], h.array[(i-1)/2] = h.array[(i-1)/2], h.array[i]
		}
	} else if h.direction == Max {
		for i := h.size - 1; h.array[i] > h.array[(i-1)/2]; i = (i - 1) / 2 {
			h.array[i], h.array[(i-1)/2] = h.array[(i-1)/2], h.array[i]
		}
	}
}
