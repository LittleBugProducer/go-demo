package main

import (
	"container/heap"
	"fmt"
)

func main() {
	h := hp{}
	heap.Init(&h) // 堆化

	heap.Push(&h, &ListNode{Val: 2})
	heap.Push(&h, &ListNode{Val: 22})
	heap.Push(&h, &ListNode{Val: 342})
	heap.Push(&h, &ListNode{Val: 162})
	heap.Push(&h, &ListNode{Val: 12})

	fmt.Println(heap.Pop(&h))
	fmt.Println(heap.Pop(&h))
	fmt.Println(heap.Pop(&h))
	fmt.Println(heap.Pop(&h))
	fmt.Println(heap.Pop(&h))

}

type ListNode struct {
	Val  int
	Next *ListNode
}

type hp []*ListNode

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].Val < h[j].Val } // 最小堆
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(*ListNode)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
