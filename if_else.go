package main

import "fmt"

func main() {
	/*IF ELSE statement*/

	j := 8
	if j%2 == 0 {
		fmt.Println("J is Even")
	} else {
		fmt.Println("J is Odd")
	}

	/*If Else-If Else along with initial*/
	if k := 10; k%3 == 0 {
		fmt.Println("K is divisible by 3")
	} else if k < 11 {
		fmt.Println("K is less than 11")
	} else {
		fmt.Println("K value is 10")
	}
}
