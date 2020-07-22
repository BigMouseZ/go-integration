package main

import (
	"fmt"
)

func main() {

	var (
		name    string
		age     int
		married bool
	)
	// 1:name 2:10 3:t
	_, err := fmt.Scanln(&name, &age, &married)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("扫描结果 name:%s age:%d married:%t", name, age, married)
}
