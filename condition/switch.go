package condition

import (
	"fmt"
)

func GeneralSwitch()  {
	var point = 6

	switch point {
		case 8:
			fmt.Println("perfect")

		case 7:
			fmt.Println("awesome")

		default:
			fmt.Println("not bad")
	}

}

func MultipleCaseSwitch()  {
	var point = 6

	switch point {
		case 8:
			fmt.Println("perfect")

		case 7,6,5,4:{
			fmt.Println("case & default with curly braces")
			fmt.Println("awesome")
			fmt.Println("lumayanlah")
		}
			

		default:{
			fmt.Println("not bad")
			fmt.Println("this is default with curlybraces")
		}		
	}
}

func SwitchCaseLikeIfElse()  {
	var point = 6

	switch {
		case point == 8:
			fmt.Println("Perfect")
		case (point < 8) && (point > 3):
			fmt.Println("awesome")
		default:
			{
				fmt.Println("not bad")
				fmt.Println("you need to learn more")
			}
	}
}

func SwithWithFallthrough()  {
	// fallthrough will force switch case to run next case, and consider it as true	
	var point = 6

	switch {
		case point == 8:
			fmt.Println("Perfect")
		case (point < 8) && (point > 3):
			fmt.Println("awesome")
			fallthrough
		default:
			{
				fmt.Println("not bad")
				fmt.Println("you need to learn more")
			}
	}
}