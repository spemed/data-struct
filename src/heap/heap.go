package heap

type compare func(int, int) bool

type Heap struct {
	Data   []int
	length int
	compare
}

func NewMaxHeap(length int) *Heap {
	if length <= 0 {
		return nil
	}
	data := make([]int, length)
	return &Heap{
		Data:   data,
		length: 0,
		compare: func(a, b int) bool {
			return a > b
		},
	}
}

func NewMinHeap(length int) *Heap {
	data := make([]int, length)
	return &Heap{
		Data: data,
		compare: func(a, b int) bool {
			return a < b
		},
	}
}

func (heap *Heap) shiftUp(length int) {
	for length > 0 {
		j := (length - 1) / 2
		if heap.compare(heap.Data[length], heap.Data[j]) {
			heap.Data[length], heap.Data[j] = heap.Data[j], heap.Data[length]
		} else {
			break
		}
		length = (length - 1) / 2
	}
}

func (heap *Heap) PileUp(data int) {
	length := heap.length
	heap.length++ //堆的属性length自增
	heap.Data[length] = data
	heap.shiftUp(length)
}

func (heap *Heap) shiftDown(length int) {
	heap.Data[0], heap.Data[length] = heap.Data[length], heap.Data[0]
	i, j := 0, 0
	for {
		j = i*2 + 1
		if j >= length {
			break
		}
		if j+1 < length {
			if heap.compare(heap.Data[j+1], heap.Data[j]) {
				j = j + 1
			}
		}
		if heap.compare(heap.Data[j], heap.Data[i]) {
			heap.Data[j], heap.Data[i] = heap.Data[i], heap.Data[j]
			i = j
		} else {
			break
		}
	}
}

func (heap *Heap) Top() int {
	return heap.Data[0]
}

func (heap *Heap) PileDown() int {
	data := heap.Data[0]
	heap.length--
	heap.shiftDown(heap.length)
	return data
}

func (heap *Heap) Travel() []int {
	return heap.Data[:heap.length]
}

//todo 实现 heapfiy
func NewMaxHeapWithData(data []int) *Heap {
	heap := &Heap{
		Data:   data,
		length: len(data),
		compare: func(a, b int) bool {
			return a > b
		},
	}
	heap.heapfiy()
	return heap
}

func (heap *Heap) heapfiy() {
	for i := (heap.length - 2) / 2; i >= 0; i-- {
		index := i
		for {
			j := index*2 + 1
			if j >= (heap.length - 1) {
				break
			}
			if j+1 <= (heap.length - 1) {
				if heap.compare(heap.Data[j+1], heap.Data[j]) {
					j = j + 1
				}
			}
			if heap.compare(heap.Data[j], heap.Data[index]) {
				heap.Data[j], heap.Data[index] = heap.Data[index], heap.Data[j]
				index = j
			} else {
				break
			}
		}
	}
}
