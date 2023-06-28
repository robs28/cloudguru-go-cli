/* If there is a variable you won't use, you get an error
Ex: Var not used etc etc
So if you don't mean to use a variable you can swap it with underscore 
when declaring it... not a good example though */

package main

import "fmt"

func main() {
	var name string
	var myInt int = 16
	var val, ok = "yes", true
	Numb := 18
	name = "Robs"
//you can directly declare a variable as well

	fmt.Println("myInt is:", myInt)
	fmt.Println("myInt times two:", myInt*2)
	fmt.Println("val is:", val)
	fmt.Println("ok is:", ok)
	fmt.Println("Number:", Numb)
	fmt.Println("Name is:", name)
}
