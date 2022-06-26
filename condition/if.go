package condition

import "fmt"

func GaneralIf()  {
	var point = 8

	if point == 10 {
		fmt.Println("Anda lulus dengan sempurna")
	} else if point >= 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("Tidak lulus, nilai anda %d\n", point)
	}
}

func TemporaryVariable()  {
	var point = 8840.0
	
	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70{
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}
}