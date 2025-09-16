package main

import "fmt"

func accessInventory(FirstCharacter Character) {
	fmt.Println("===============INVENTAIRE================")
	for i := 0; i < len(FirstCharacter.Inventory); i++ {
		fmt.Println(i, ":", FirstCharacter.Inventory[i].Name,", quantité :",FirstCharacter.Inventory[i].Quantity)
	}
	fmt.Println("tapez un numéro pour utiliser un objet, sinon tapez \"exit\"")
	fmt.Scanln(&answer)
	if FirstCharacter.Inventory[answer].Consumables == true {
		FirstCharacter.Inventory[answer].Quantity -= 2
	}

	switch answer {
	case "0":

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

	case "exit":
		return
	}
}
