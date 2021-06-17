package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greetings := "hello there friends!"
	fmt.Println(strings.Contains(greetings, "hello"))
	fmt.Println(strings.ReplaceAll(greetings, "hello", "hi"))
	fmt.Println(strings.Index(greetings, "ll"))
	fmt.Println(strings.Split(greetings, " "))

	ages := []int{45, 25, 35, 56, 21, 12, 3}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 60)
	fmt.Println(index)

	names := []string{"John", "Peter", "Jared", "Gilfoyle"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "Jared"))
}
