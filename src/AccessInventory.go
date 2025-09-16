package main

import "fmt"

func accessInventory(FirstCharacter Character) {
	for i := 0; i < len(FirstCharacter.Inventory); i++ {
		fmt.Println(FirstCharacter.Inventory[i].Name, FirstCharacter.Inventory[i].Quantity)
	}
}
