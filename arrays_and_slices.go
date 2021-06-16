package main

import (
	"fmt"
)

func main() {
	// var ages []int = []int{10, 20, 30}
	var ages = []int{10, 20, 30}
	fmt.Println(ages, len(ages))

	names := []string{"John", "Peter", "Jared", "Gilfoyle"}
	fmt.Println(names, len(names))
	fmt.Println(names[0])

	// Slices

	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 40)
	fmt.Println(scores)

	// Slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

}
