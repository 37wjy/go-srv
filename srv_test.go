package main

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	scene := make(map[string]int)
	// 准备map数据
	scene["route"] = 66
	scene["china"] = 960
	delete(scene, "brazil")
	for k, v := range scene {
		fmt.Println(k, v)
	}
}
