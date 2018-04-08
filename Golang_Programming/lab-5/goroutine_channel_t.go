package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 创建一个类型为chan bool的切片，每一项是能发送bool值的通道,缓存区容量为6
	channels := make([]chan bool, 6)
	for i := range channels {
		channels[i] = make(chan bool)
	}

	go func() {
		for {
			channels[rand.Intn(6)] <- true
		}
	}()
	fmt.Println(channels)
	for i := 0; i < 36; i++ {
		var x int
		select {
		case <-channels[0]:
			x = 1
		case <-channels[1]:
			x = 2
		case <-channels[2]:
			x = 3
		case <-channels[3]:
			x = 4
		case <-channels[4]:
			x = 5
		case <-channels[5]:
			x = 6
		}
		fmt.Printf("%d ", x)
	}
	fmt.Println()
}
