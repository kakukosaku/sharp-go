package time_test

import (
	"fmt"
	"time"
)

func Example() {
	fmt.Println(1 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	done := make(chan bool)
	go func() {
		time.Sleep(3 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("3s elapsed, end!")
			return
		case <-ticker.C:
			fmt.Println("1s elapsed")
		}
	}
	// Output:
	// 1s
	// 1s elapsed
	// 1s elapsed
	// 1s elapsed
	// 3s elapsed, end!
}
