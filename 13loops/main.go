package main

import "fmt"

func main() {
	fmt.Println("welcome to loops in golang")

	// days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i:= range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and day is %v\n", index, days)
	// }

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 4 {
			goto label_lco
		}

		if rogueValue == 5 {
			rogueValue++
			continue
		}

		fmt.Println("Value is: ", rogueValue)
		rogueValue++
	}

label_lco:
	fmt.Println("Jumping at LCO")
}
