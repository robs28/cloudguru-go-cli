package main
import "fmt"
func main(){
	ages := map[string]int{}
	ages["Robs"] = 27

	switch ages ["Robs"] {
		case 1, 2, 3, 4, 5, 7, 11, 13, 17, 19:
			fmt.Println("Robs age is a prime number") 
		case 16:
			fmt.Println("Robs can drive now")
		case 18: 
			fmt.Println("Robs can vote now")
		case 67:
			fmt.Println("Robs can retire now")
		default:
			fmt.Println("There's nothing special about Robs current age")
		}


	/*
	switch {
	case ages["Robs"] < 18:
		fmt.Println("Robs can't vote yet")
	case ages["Robs"] < 67:
		fmt.Println("Robs is not of retirement age just yet")
	default:
		fmt.Println("Enjoy retirement")
	}

	if ages["Robs"] < 67 {
		fmt.Println("Robs is not of retirement age")
	} else if ages ["Robs"] < 18 {
		fmt.Println("Robs can't vote")
	} else {
		fmt.Println("Enjoy retirement!")
	}
*/
}
