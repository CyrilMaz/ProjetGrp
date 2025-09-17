package main

import "fmt"

func DisplayInfo(p *Character) {
	fmt.Println("")
	fmt.Println("◇─◇──◇────◇ INFORMATION ◇─────◇──◇─◇")
	fmt.Println("Nom :", p.Name)
	fmt.Println("Classe :", p.Class)
	fmt.Println("Niveau :", p.Level)
	fmt.Println("Points de vie actuels :", p.Pv)
	fmt.Println("Points de vie max :", p.Pvmax)
	fmt.Println("Pieces d'or :", p.Gold)
	fmt.Println("Compétences apprises :")
	for i := 0; i < len(p.Skills); i++ {
		fmt.Printf(p.Skills[i].Name)
	}
	fmt.Println("")
	fmt.Println("◇─◇──◇────◇ INVENTAIRE ◇─────◇──◇─◇")
	for i := 0; i < len(p.Inventory); i++ {
		fmt.Println("tu as", p.Inventory[i].Quantity, p.Inventory[i].Name)
	}
}
