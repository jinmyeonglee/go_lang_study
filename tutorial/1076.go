package main

import "fmt"
import "math"

func main() {
	register_val := map[string] int {
		"black": 0,
		"brown": 1,
		"red": 2,
		"orange": 3,
		"yellow": 4,
		"green": 5,
		"blue": 6,
		"violet": 7,
		"grey": 8,
		"white": 9,
	}

	var first, second, third string
	fmt.Scan(&first, &second, &third)

	fmt.Println(register_val[first] * int(math.Pow10(register_val[third]))+ register_val[second] * int(math.Pow10(register_val[third])))

}