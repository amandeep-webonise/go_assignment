package main

import (
	"fmt"
	"time"
)

func main() {
	//x := 2 * time.Second
	//fmt.Println(x)
	timer2 := time.NewTimer(time.Second)
	<-timer2.C
	fmt.Println("This is timer format")

	timer3 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer3.C
		fmt.Println("Timer 2 experied")
	}()

	stop3 := timer3.Stop()
	if stop3 {
		fmt.Println("Timer3 Stopper")
	}

}
