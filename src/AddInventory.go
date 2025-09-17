package main

import "fmt"

func AddInventory(p *Character, i Items) { // Ajoute un item à l'inventaire du personnage , si jamais l'item est déjà dans l'inventaire, incrémente la quantité
	if len(p.Inventory) == 0 { // si l'inventaire est vide, on ajoute l'item directement
		p.Inventory = append(p.Inventory, i) // ajoute l'item à l'inventaire via la fonction append
		if DisplayAddItem {
			fmt.Println("Vous avez ajouté", i.Quantity, i.Name, "à votre inventaire.")
		}
	} else if !InventoryLimitCheck(p) { // si l'inventaire est plein, on ne peut pas ajouter d'item
		fmt.Println("Votre inventaire est plein, vous ne pouvez pas ajouter d'item.")
	} else {
		for j := 0; j < len(p.Inventory); j++ { // cette boucle fort incrémente la quantité d'un item si il est déjà dans l'inventaire
			if p.Inventory[j].Name == i.Name { // compare le nom de l'item à ajouter avec les items déjà dans l'inventaire
				p.Inventory[j].Quantity += i.Quantity // incrémente la quantité de l'item
				if DisplayAddItem {
					fmt.Println("Vous avez ajouté", i.Quantity, i.Name, "à votre inventaire.")
				}
			} else if j == len(p.Inventory)-1 { // si l'item n'est pas dans l'inventaire, on l'ajoute
				p.Inventory = append(p.Inventory, i) // ajoute l'item à l'inventaire via la fonction append
				if DisplayAddItem {
					fmt.Println("Vous avez ajouté", i.Quantity, i.Name, "à votre inventaire.")
				}
			}
		}
	}
	DisplayAddItem = true
}
