package main

func InventoryLimitCheck(p *Character) bool {
	if len(p.Inventory) >= 9 {
		return false
	} else {
		return true
	}
}
