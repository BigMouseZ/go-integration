package main

import (
	"fmt"
	"index/suffixarray"
	"sort"
	"testing"
)

func Test(t *testing.T) {

	// suffixarray模块提供了基于前缀数组的子串检索功能，能够在byte数组中检索指定子串，并获得其索引下标
	source := []byte("hello world, hello china hello hello ")
	index := suffixarray.New(source)

	offsets := index.Lookup([]byte("hello"), -1)
	fmt.Printf("%v", offsets)
	sort.Ints(offsets)

	fmt.Printf("%v", offsets)
}
