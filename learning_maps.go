package main

import "fmt"

func main() {
	birthdays := map[string]string{
		"Keith": "02/29/1988",
		"Robs": "28/02/1996",
		"Emily": "05/12/1998",
	}

//	delete(birthdays, "Robs")

	ages := map[string]int{}
	ages["Keith"] = 28
	ages["Robs"] = 27
	ages["Emily"] = 24

//	fmt.Println(ages["Robs"])
	fmt.Println(birthdays)
}
