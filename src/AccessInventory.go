package main

import "fmt"

func accessInventory(FirstCharacter Character) {
<<<<<<< HEAD
	fmt.Println("0 : sortir")	
	for i := 1; i < len(FirstCharacter.Inventory) - 1; i++ {
		fmt.Println(i, ":", FirstCharacter.Inventory[i].Name, FirstCharacter.Inventory[i].Quantity)
	}
	fmt.Scanln(&answer)
	switch answer {
	case "0":
		return
	case "1":

	case "2":

	case "3":

	case "4":

	case "5":

	case "6":

	case "7":

	case "8":

	case "9":

	case "10":
=======
	for i := 0; i < len(FirstCharacter.Inventory); i++ {
		fmt.Println(FirstCharacter.Inventory[i].Name, FirstCharacter.Inventory[i].Quantity)
>>>>>>> 46971e2 (fix add boucle)
	}
}
