package main

import "time"

//测试select

func main() {
	c1 := make(chan byte)
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- 1
	}()

	loop:
	for {
		select{
		case <- c1:
			println("end")
			break loop
		default:
		}
	}

}
