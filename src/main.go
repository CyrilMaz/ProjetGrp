package main

import "fmt"

var initalized bool = false
var answer string

func main() {
	if initalized == false {
		initCharacter()
		initalized = true
	}
	for answer != "0" {
		fmt.Println("\nQue voulez-vous faire ?")
		for i := 0; i < len(MENU01); i++ {
			fmt.Println(MENU01[i])
		}

		fmt.Print("Votre choix : ")
		fmt.Scanln(&answer)

		switch answer {
		case "1":
			DisplayInfo(FirstCharacter)
		case "2":
			accessInventory(&FirstCharacter)
		case "3":
			Merchand(&FirstCharacter)
		case "0":
			fmt.Println("Au revoir", FirstCharacter.Name, "!")
		default:
			fmt.Println("Choix invalide, réessayez.")
		}
	}
}
