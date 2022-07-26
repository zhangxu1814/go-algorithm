package structures

type MaxPriorityQueue struct {
	pq   []int
	size int
}

func NewMaxPriorityQueue(cap int) *MaxPriorityQueue {
	return &MaxPriorityQueue{
		pq:   make([]int, cap+1, cap+1),
		size: 0,
	}
}

func (m *MaxPriorityQueue) max() int {
	return m.pq[1]
}

func (m *MaxPriorityQueue) insert(n int) {
	m.size++
	m.pq[m.size] = n
	m.swim(m.size)
}

func (m *MaxPriorityQueue) delete() {
	m.swap(1, m.size)
	m.pq[m.size] = 0
	m.size--
	m.sink(1)
}

//-------------二叉堆操作--------------//

func (m *MaxPriorityQueue) swim(i int) {
	for i > 1 && m.cmp(i, parent(i)) {
		m.swap(i, parent(i))
		i = parent(i)
	}
}

func (m *MaxPriorityQueue) sink(i int) {
	for left(i) <= m.size {
		max := left(i)
		if right(i) <= m.size && m.cmp(right(i), left(i)) {
			max = right(i)
		}
		if m.cmp(i, max) {
			break
		}
		m.swap(i, max)
		i = max
	}
}

func (m *MaxPriorityQueue) swap(i, j int) {
	m.pq[i], m.pq[j] = m.pq[j], m.pq[i]
}

func (m *MaxPriorityQueue) cmp(i, j int) bool {
	return m.pq[i] > m.pq[j]
}

func left(root int) int {
	return root * 2
}

func right(root int) int {
	return root*2 + 1
}

func parent(root int) int {
	return root / 2
}
