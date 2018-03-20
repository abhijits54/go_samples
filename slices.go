package main

import "fmt"


func main(){

	s := make([]string, 3)
	s[0] = "hey"
	s[1] = "there"
	s[2] = "wadup"
	fmt.Println("S: ", s)
	s = append(s, "yo", "lo")


	fmt.Println("appended s:", s)


	var k [3] string
	k[0] = "hi"
	k[1] = "there"
	k[2] = "ll"
	fmt.Println("k is an array not a slice: ", k)

	c := make([]string, len(s))
	copy(c,s)
	fmt.Println(" c is a copy of s:", c)

}

