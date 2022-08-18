package test

import (
	"fmt"
	"testing"
)

func RpcTest(t *testing.T) {
	tlist := []int{1, 2, 3}
	tlist = append(tlist, 1)
	fmt.Print(tlist)
}
