package main

import "fmt"

func main(){

	m := make(map[string]int)
	m["l1"]  = 2
	m["l2"] = 3
	fmt.Println(m)
	delete(m,"l2")
	fmt.Println("after deleting", m)

}
