package main

import (
	"fmt"
	"time"
)

func main() {
	tick1 := time.Tick(3 * time.Second)
	tick2 := make(chan struct{})

	for {
		select {
		case <-tick1:
			fmt.Printf("tick 1\n")
			go func() {
				time.Sleep(5 * time.Second)
				tick2 <- struct{}{}
			}()
		case <-tick2:
			fmt.Printf("tick 2\n")
		}
	}
}
