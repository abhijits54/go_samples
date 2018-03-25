package main

import "fmt"

import "sort"

func main(){
	strs := []string{"c", "b", "d", "a", "l"}
	sort.Strings(strs)
	fmt.Println("Strings: ", strs)

	ints := []int{8,2,4,6}
	sort.Ints(ints)
	fmt.Println("Ints: ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

}
