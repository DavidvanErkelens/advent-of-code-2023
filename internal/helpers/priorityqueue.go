package helpers

import "slices"

type PriorityQueue[T any] struct {
	keys     []int
	values   map[int][]T
	elements int
}

func BuildPriorityQueue[T any](from T, priority int) PriorityQueue[T] {
	return PriorityQueue[T]{
		keys: []int{priority},
		values: map[int][]T{
			priority: {from},
		},
		elements: 1,
	}
}

func (pq *PriorityQueue[T]) Push(value T, priority int) {
	pq.elements += 1
	if _, ok := pq.values[priority]; ok {
		pq.values[priority] = append(pq.values[priority], value)
		return
	}

	pq.values[priority] = []T{value}
	pq.keys = append(pq.keys, priority)
}

func (pq *PriorityQueue[T]) Pop() T {
	pq.elements -= 1
	lowestPrio := slices.Min(pq.keys)
	value := pq.values[lowestPrio][0]

	if len(pq.values[lowestPrio]) > 1 {
		pq.values[lowestPrio] = pq.values[lowestPrio][1:]
	} else {
		delete(pq.values, lowestPrio)
		pq.keys = RemoveIndex(pq.keys, FindIndex(pq.keys, lowestPrio))
	}

	return value
}

func (pq *PriorityQueue[T]) HasElements() bool {
	return pq.elements > 0
}

func (pq *PriorityQueue[T]) Elements() int {
	return pq.elements
}
