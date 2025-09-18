package main

import "fmt"

func RemoveInventory(p *Character, i Items, n int) {
	for j := 0; j < len(p.Inventory); j++ {
		if p.Inventory[j].Name == i.Name {
			if p.Inventory[j].Quantity > 0 {
				for k := 0; k < n; k++ {
					if p.Inventory[j].Quantity > 0 {
						p.Inventory[j].Quantity--
					}
				}
				if p.Inventory[j].Type == "consommable" {
					fmt.Println("Vous avez utilisé une", i.Name)
				} else if p.Inventory[j].Type == "équipement" {
					fmt.Println("Vous avez équipé", i.Name)
				} else {
					fmt.Println("Vous avez enlevé", i.Name, "de votre inventaire")
				}
			}
		}
	}
}
