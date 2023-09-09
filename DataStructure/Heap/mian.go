package main

import (
	// use container heap
	// we can use list & ring
	"container/heap"
	"fmt"

	"github.com/kr/pretty"
)

type HeapFloat []float64

func (h *HeapFloat) Pop() any {
	old := *h

	x := old[len(old)-1]
	new := old[0 : len(old)-1]
	*h = new
	return x
}

func (h *HeapFloat) Push(x any) {
	*h = append(*h, x.(float64))
}

func (h HeapFloat) Len() int {
	return len(h)
}

func (h HeapFloat) Less(a, b int) bool {
	return h[a] < h[b]
}

func (h HeapFloat) Swap(a, b int) {
	h[a], h[b] = h[b], h[a]
}

func main() {
	h := &HeapFloat{1.2, 2.1, 3.1, -100.1}
	heap.Init(h)

	fmt.Printf("Heap size: %d\n", len(*h))
	fmt.Printf("%# v\n", pretty.Formatter(h))

	h.Push(-100.2)
	h.Push(0.2)
	fmt.Printf("Heap size: %d\n", len(*h))

	fmt.Printf("%# v\n", pretty.Formatter(h))
	heap.Init(h)
	fmt.Printf("%# v\n", pretty.Formatter(h))
}
