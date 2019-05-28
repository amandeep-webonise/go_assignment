package main

import "fmt"
import "time"

func main() {
	/*switch statement*/
	i := 3
	switch i {
	case 1:
		fmt.Println("Value is ONE")
	case 2:
		fmt.Println("Value is Two")
	case 3:
		fmt.Println("Value is Three")
	}

	//switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Today is weekend")
	default:
		fmt.Println("Today is weekday")

	}

	whatIAm := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a integer")
		default:
			fmt.Println("Do not know type %T\n", t)
		}
	}
	whatIAm(true)
	whatIAm(1)
}
