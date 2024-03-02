package concurrency

import (
	"fmt"
	"time"
)

func DefaultSelection() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(1000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-boom:
			fmt.Println("Boom")
			return
		default:
			fmt.Println("   .")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
