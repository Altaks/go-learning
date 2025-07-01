package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Print("Write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Local().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Its the week")
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Its before noon")
	default:
		fmt.Printf("Its after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("It's a bool")
		case int:
			fmt.Println("Its an int")
		default:
			fmt.Printf("It's something else, idk type : %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("heyyy")
}
