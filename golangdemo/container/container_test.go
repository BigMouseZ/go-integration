package main

import (
	"container/heap"
	"container/list"
	"container/ring"
	"fmt"
	"testing"
)

type Student struct {
	name  string
	score int
}

type StudentHeap []Student

func (h StudentHeap) Len() int { return len(h) }

func (h StudentHeap) Less(i, j int) bool {
	return h[i].score < h[j].score // 最小堆
	// return stu[i].score > stu[j].score //最大堆
}

func (h StudentHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *StudentHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Student))
}

func (h *StudentHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TestHeap(t *testing.T) {
	/*
		heap包提供了对任意类型（实现了heap.Interface接口）的堆操作。（最小）堆是具有“每个节点都是以其为根的子树中最小值”属性的树。

		树的最小元素为其根元素，索引0的位置。

		heap是常用的实现优先队列的方法。要创建一个优先队列，实现一个具有使用（负的）优先级作为比较的依据的Less方法的Heap接口，
		此一来可用Push添加项目而用Pop取出队列最高优先级的项目。
	*/
	h := &StudentHeap{
		{name: "xiaoming", score: 82},
		{name: "xiaozhang", score: 88},
		{name: "laowang", score: 85}}

	// 初始化一个堆。一个堆在使用任何堆操作之前应先初始化。
	// Init函数对于堆的约束性是幂等的（多次执行无意义），并可能在任何时候堆的约束性被破坏时被调用。
	// 本函数复杂度为O(n)，其中n等于h.Len()。
	heap.Init(h)

	// 向堆h中插入元素x，并保持堆的约束性。复杂度O(log(n))，其中n等于h.Len()。
	heap.Push(h, Student{name: "xiaoli", score: 66})

	for _, ele := range *h {
		fmt.Printf("student name %s,score %d\n", ele.name, ele.score)
	}

	for i, ele := range *h {
		if ele.name == "xiaozhang" {
			(*h)[i].score = 60

			// 在修改第i个元素后，调用本函数修复堆，比删除第i个元素后插入新元素更有效率。
			// 复杂度O(log(n))，其中n等于h.Len()。
			heap.Fix(h, i)
		}
	}

	fmt.Println("==========")

	for _, ele := range *h {
		fmt.Printf("student name %s,score %d\n", ele.name, ele.score)
	}

	fmt.Println("==========")

	for h.Len() > 0 {
		// 删除并返回堆h中的最小元素（取决于Less函数，最大堆或最小堆）（不影响堆de约束性）
		// 复杂度O(log(n))，其中n等于h.Len()。该函数等价于Remove(h, 0)
		item := heap.Pop(h).(Student)
		fmt.Printf("student name %s,score %d\n", item.name, item.score)
	}

}
func TestList(t *testing.T) {
	/*
		list包实现了双向链表。要遍历一个链表：
		双向链表是链表的一种，它的每个数据结点中都有两个指针，分别指向直接后继和直接前驱。
		所以，从双向链表中的任意一个结点开始，都可以很方便地访问它的前驱结点和后继结点。一般我们都构造双向循环链表。
	*/

	l := list.New()
	l.PushBack(1) // 尾插
	l.PushBack(2)
	print(l)

	fmt.Println("=========")

	l.PushFront(0) // 头插
	print(l)

	fmt.Println("=========")

	// 往后读取
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == 1 {
			l.InsertAfter(1.1, e)
		}

		if e.Value == 2 {
			l.InsertBefore(1.2, e)
		}
	}

	print(l)

	fmt.Println("=========")

	fmt.Println(l.Front().Value) // 返回链表的第一个元素
	fmt.Println("=========")

	fmt.Println(l.Back().Value) // 返回链表的最后一个元素
	fmt.Println("=========")

	// 将数据转移尾部
	l.MoveToBack(l.Front())
	print(l)

	fmt.Println("=========")

	// 从后往前读取
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

	fmt.Println("=========")
	// 前往往后读取
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
func TestRing(t *testing.T) {
	/*
		Ring类型代表环形链表的一个元素，同时也代表链表本身。环形链表没有头尾；
		指向环形链表任一元素的指针都可以作为整个环形链表看待。Ring零值是具有一个（Value字段为nil的）元素的链表。
	*/
	ring1 := ring.New(3)

	for i := 1; i <= 3; i++ {
		ring1.Value = i
		ring1 = ring1.Next()
	}

	ring2 := ring.New(3)

	for i := 4; i <= 6; i++ {
		ring2.Value = i
		ring2 = ring2.Next()
	}

	r := ring1.Link(ring2)

	fmt.Printf("ring length = %d\n", r.Len())

	r.Do(func(p interface{}) {
		fmt.Print(p.(int))
		fmt.Print(",")
	})

	fmt.Println()

	fmt.Printf("current ring is %v\n", r.Value)

	fmt.Printf("next ring is %v\n", r.Next().Value)

	fmt.Printf("prev ring is %v\n", r.Prev().Value)

	// ring 的遍历
	for p := r.Next(); p != r; p = p.Next() {
		fmt.Print(p.Value.(int))
		fmt.Print(",")
	}

}
