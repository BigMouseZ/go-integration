package strconvtest

import (
	"fmt"
	"strconv"
	"testing"
)

func TestOne(t *testing.T) {
	fmt.Println(strconv.FormatInt(int64(10), 10))
	fmt.Println(strconv.Itoa(10))

}
