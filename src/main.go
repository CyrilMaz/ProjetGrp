package main

import "fmt"

var answer string 

func main() {
	initCharacter()
	fmt.Println("Que voulez vous faire ?")
	for i := 0; i < len(MENU01); i++ {
		fmt.Println(MENU01[i])
	}
	fmt.Scanln(&answer)
	switch answer {
	case "0":
		return
	case "1":
		DisplayInfo(FirstCharacter)
	case "2":
		accessInventory(FirstCharacter)
	}
}
