package main
import "fmt"

func main(){
	ages := map[string]int{}
	ages["Robs"] = 11
	ages["Emily"] = 28
	ages["Daniel"] = 67
	ages["James"] = 18
	ages["Leigha"] = 16

	for name, age := range ages {
		switch age{
		case 1,2,3,5,7,11,13,17,19:
			fmt.Println(fmt.Sprintf("%s's age is a prime number", name))
		case 16:
			fmt.Println(name, "can drive now")
		case 18:
			fmt.Println(name, "can retire now")
		case 67:
			fmt.Println(fmt.Sprintf("There's nothing special about %s's current age", name))
		}
	}


		for i := 1; i <= 10; i++ {
			fmt.Println("We're counting:", i)
		
	}

	a := 0
	for a < 10 {
		fmt.Println("We're counting (again):", a)
		a++
	}

	b := 0
	for b < 10{
		if b%2 == 0{
			b++
			continue
		} else if b == 5 {
			break
		}
		fmt.Println("We're counting again...:", b)
		b++
	}

}
