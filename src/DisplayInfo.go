package main

import "fmt"

func DisplayInfo(FirstCharacter *Character) {
	fmt.Println("")
	fmt.Println("◇─◇──◇────◇ INFORMATION ◇─────◇──◇─◇")
	fmt.Println("Nom :", FirstCharacter.Name)
	fmt.Println("Classe :", FirstCharacter.Class)
	fmt.Println("Niveau :", FirstCharacter.Level)
	fmt.Println("Points de vie actuels :", FirstCharacter.Pv)
	fmt.Println("Points de vie max :", FirstCharacter.Pvmax)
	fmt.Println("Pieces d'or :", FirstCharacter.Gold)
	fmt.Println("Compétences apprises :")
	for i := 0; i < len(FirstCharacter.Skills); i++ {
		if FirstCharacter.Skills[i].Learned {
			fmt.Println(FirstCharacter.Skills[i].Name)
		}
	}
	fmt.Println("")
	fmt.Println("◇─◇──◇────◇ INVENTAIRE ◇─────◇──◇─◇")
	for i := 1; i <= len(FirstCharacter.Inventory); i++ {
		if FirstCharacter.Inventory[i-1].Quantity > 0 {
			fmt.Println(i, ":", FirstCharacter.Inventory[i-1].Name, ", quantité :", FirstCharacter.Inventory[i-1].Quantity)
		}
	}
}
