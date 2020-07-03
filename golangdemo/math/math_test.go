package main

import (
	"fmt"
	"math"
	"testing"
)

func TestOne(t *testing.T) {
	num := math.Abs(float64(2302.2))
	fmt.Printf("type:%T value:%#v\n", num, num)

	// 四舍五入
}
