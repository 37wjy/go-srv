package main

import (
	"UnicornServer/core"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func TestCfg(t *testing.T) {

	config := &core.ServerCfg{}

	f, _ := ioutil.ReadFile("config/config.json")
	err := json.Unmarshal(f, config)
	if err != nil {
		fmt.Println("unmarshal faild ", err)
	}
	fmt.Println(config.Name)
	fmt.Println(config.Port)
	fmt.Println(config.IP)
}
