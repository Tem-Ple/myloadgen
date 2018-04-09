package main

import (
	"time"
	"log"
	"fmt"
)

// 测试断续器

func main() {
	var throttle <-chan time.Time
	interval := time.Duration(time.Second*2)
	log.Printf("Setting throttle (%v)...", interval)
	throttle = time.Tick(interval)
	for{
		time := <- throttle
		fmt.Println(time)
	}
}
