package collection

import "fmt"

func MapCollection()  {
	fmt.Println("Initialiation map")
	var chicken = map[string]int{
		"january":4,
		"february":5,
		"march":6,
		"april":7,
	}

	for key, item := range chicken{
		fmt.Println(key," \t", item)
	}

	fmt.Print("\n\n\n")
	fmt.Println("Delete map")
	delete(chicken, "january")
	for key, item := range chicken{
		fmt.Println(key," \t", item)
	}

	fmt.Print("\n\n\n")
	fmt.Println("Get value of map")
	item, isExists := chicken["march"]
	if isExists {
		fmt.Println("march of chicken is ",item)
	} else {
		fmt.Println("item is not exists")
	}
}