package main

import "fmt"

func UpgradeInventorySlot(p *Character) {
	p.MaxInventory += 10
	fmt.Println("Votre inventaire peut maintenant contenir", p.MaxInventory, "objets !")
}
