package main

import (
	"fmt"
	"sort"
)

/*

type Interface
type Interface interface {

    Len() int    // Len 为集合内元素的总数

    Less(i, j int) bool　//如果index为i的元素小于index为j的元素，则返回true，否则返回false

    Swap(i, j int)  // Swap 交换索引为 i 和 j 的元素
}

任何实现了 sort.Interface 的类型（一般为集合），均可使用该包中的方法进行排序。这些方法要求集合内列出元素的索引为整数。
*/

type People struct {
	Name string  `json:"name"`
	Age  float64 `json:"age"`
}

func isNaN(f float64) bool {
	return f != f
}

type testSlice []People

func (l testSlice) Len() int      { return len(l) }
func (l testSlice) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l testSlice) Less(i, j int) bool {
	return l[i].Age < l[j].Age || isNaN(l[i].Age) && !isNaN(l[j].Age)
}

func main() {
	ls := testSlice{
		{Name: "n1", Age: 12.12},
		{Name: "n2", Age: 11.11},
		{Name: "n3", Age: 10.10},
	}

	fmt.Println(ls) // [{n1 12.12} {n2 11.11} {n3 10.1}]
	sort.Sort(ls)
	fmt.Println(ls) // [{n3 10.1} {n2 11.11} {n1 12.12}]
}
