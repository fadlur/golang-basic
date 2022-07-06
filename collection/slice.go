package collection

import (
	"fmt"
	"reflect"
)

func SliceCollection()  {
	// slice is array reference
	var fruits = []string{"apple", "grape", "melon", "orange"}
	fmt.Println(fruits[0])// apple
	fmt.Print("\n\n\n")
	fmt.Println("Slice or array")
	var fruitA = []string{"melon", "orange"}// slice
	var fruitB = [2]string{"banana", "papaya"}// array
	var fruitC = [...]string{"pineapple", "grape"}// array
	fmt.Printf("fruitA is %v\n", reflect.TypeOf(fruitA).Kind())
	fmt.Printf("fruitB is %v\n", reflect.TypeOf(fruitB).Kind())
	fmt.Printf("fruitC is %v\n", reflect.TypeOf(fruitC).Kind())
	fmt.Print("\n\n\n")
	fmt.Println("Get the element using slice")
	fmt.Printf("fruits[0:2] is %v\n", fruits[0:2])
	fmt.Printf("fruits[0:4] is %v\n", fruits[0:4])
	fmt.Printf("fruits[0:0] is %v\n", fruits[0:0])
	fmt.Printf("fruits[4:4] is %v\n", fruits[4:4])
	// fmt.Printf("fruits[4:0] is %v\n", fruits[4:0])// error invalid slice index: 4 > 0
	fmt.Printf("fruits[:] is %v\n", fruits[:])
	fmt.Printf("fruits[2:] is %v\n", fruits[2:])
	fmt.Printf("fruits[:2] is %v\n", fruits[:2])

	fmt.Print("\n\n\n")
	fmt.Println("Slice is array reference")
	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"

	fmt.Println(fruits)   // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]
	fmt.Println("Capacity is ", cap(fruits))
	fmt.Println("Len is ", len(fruits))
}