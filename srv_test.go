package main

import (
	"UnicornServer/core"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"runtime"
	"testing"
	"time"
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

func TestRef(t *testing.T) {
	arr := []int{1, 2, 3}
	newArr := []*int{}
	for _, v := range arr {
		fmt.Print(&v)
		newArr = append(newArr, &v)
	}
	fmt.Print("\n")
	for _, v := range newArr {
		fmt.Print(*v, " ")
	}
	fmt.Print("\n")
	fmt.Println(newArr)
	for _, v := range newArr {
		fmt.Print(*v, " ")
	}
	fmt.Print("\n")
}

func TestAA(t *testing.T) {
	var id int32 = 11001
	if id < core.MsgID.GM_ID_START || id >= core.MsgID.RK_ID_START {
		print(1)
	}
	print(2)
}

func TestPR(t *testing.T) {
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*2)
	go func() {
		defer func() {
			// 发生宕机时，获取panic传递的上下文并打印
			err := recover()
			switch err.(type) {
			case nil:
				fmt.Println("no err")
			case runtime.Error: // 运行时错误
				fmt.Println("runtime error:", err)
			default: // 非运行时错误
				fmt.Println("error:", err)
			}
		}()

		select {
		case <-ctx.Done():
			fmt.Println("ctx done")
			return
		}
	}()
	select {
	case <-ctx.Done():
		time.Sleep(time.Second * 1)
		return
	}
	print(111)

}
