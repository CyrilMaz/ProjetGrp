package main

import "fmt"

func DisplayInfo(p Character) {
	fmt.Println("Nom :", p.Name)
	fmt.Println("Classe :", p.Class)
	fmt.Println("Niveau :", p.Level)
	fmt.Println("Points de vie actuels :", p.Pv)
	fmt.Println("Points de vie max :", p.Pvmax)
	for i := 0; i < len(p.Inventory); i++ {
		fmt.Println("tu as", p.Inventory[i].Quantity, p.Inventory[i].Name)
	}
}
