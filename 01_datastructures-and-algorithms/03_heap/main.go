package main

import (
	"container/heap"
	"fmt"
)

type IntegerHeap []int

func (iHeap IntegerHeap) Len() int{
	return len(iHeap)
}

func (iHeap IntegerHeap) Less(i, j int) bool{
	return iHeap[i] < iHeap[j]
}

func (iHeap IntegerHeap) Swap(i, j int) {
	iHeap[i], iHeap[j] = iHeap[j], iHeap[i]
}

func (iHeap *IntegerHeap) Push(heapintf interface{}) {
	*iHeap = append(*iHeap, heapintf.(int))
}

func (iHeap *IntegerHeap) Pop() interface{} {
	var n int
	var x1 int
	previousHeap := *iHeap
	n = len(*iHeap)
	x1 = previousHeap[n-1]
	*iHeap = previousHeap[0:n-1]
	return x1
}
func main() {
	var intHeap *IntegerHeap = &IntegerHeap{1,4,5}
	heap.Init(intHeap)
	fmt.Printf("max %v\n", (*intHeap)[0])
	for intHeap.Len() > 0 {
		fmt.Println(heap.Pop(intHeap))
	}


}
