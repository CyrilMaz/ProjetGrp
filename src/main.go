package main

import "fmt"

var answer string 

func main() {
	initCharacter()
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
			accessInventory(FirstCharacter)
		case "0":
			fmt.Println("Au revoir Cyril !")
		default:
			fmt.Println("Choix invalide, réessayez.")
		}
	}
}
