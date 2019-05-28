package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(2 * time.Second)
	go func() {
		for t := range ticker.C {
			fmt.Println("First call of Ticker", t)
		}
	}()

	time.Sleep(10 * time.Second)
	ticker.Stop()
	fmt.Println("Ticker STOP CALL")
}
