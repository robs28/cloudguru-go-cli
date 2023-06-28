package main

import "fmt"

func main(){
	names := [3]string{"Robs", "Emily", "Isa"}
	var names_two [3]string

	names_two[0] = "Robs"
	names_two[1] = "Emily"
	names_two[2] = "Isa"

	fmt.Println(names_two)
	fmt.Println(names)
	//fmt.Println("names[2] is nil:", names[2] == nil)
}
