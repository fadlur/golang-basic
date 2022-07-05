package collection

import "fmt"

/*
Array is a group of data with the same type. array should be declare the length when it initiated.
*/

func ArrayCollection()  {
	var names[4]string

	names[0] = "fadlur"
	names[1] = "rohman"
	names[2] = "baba"
	names[3] = "bubu"

	fmt.Println(names)

	// if we add one more data, it will be produce an error
	// names[4] = "halo"

	// looping array using for to get data
	fmt.Print("\n\n\n")
	fmt.Println("iterate array using for")
	for i := 0; i < len(names); i++ {
		fmt.Printf("Element in index %d : %s\n", i, names[i])
	}
	fmt.Print("\n\n\n")
	fmt.Println("iterate array using range")
	for index, item := range names {
		fmt.Printf("Element index %d: is %s\n", index, item)
	}
	fmt.Print("\n\n\n")
	fmt.Println("Array multidimensi")
	var number1 = [2][3]int{{1,2,3}, {3,4,5}}
	var number2 = [2][3]int{{3,2,1}, {5,4,3}}

	fmt.Println("number 1", number1)
	fmt.Println("number 2", number2)

	fmt.Print("\n\n\n")
	fmt.Println("Array using make")
	var fruits = make([]string, 2)
	fruits[0] = "orange"
	fruits[1] = "watermelon"
	fmt.Println("fruit array", fruits)

	fmt.Print("\n\n\n")
	fmt.Println("Array without the number of content")
	var number3 = [...]int{1,2,3}
	fmt.Println("Data number3\t", number3)
	fmt.Println("Len of array\t", len(number3))

}
