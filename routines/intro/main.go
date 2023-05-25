package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(time.Second)
	}
}

// 1
func main() {
	// 2
	go task("A")
	// 3
	go task("B")
	// 4
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "C")
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 20)
}
