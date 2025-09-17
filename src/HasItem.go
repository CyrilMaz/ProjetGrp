package main

func HasItem(p *Character, i Items, n int) bool {
	for j := 0; j < len(p.Inventory); j++ {
		if p.Inventory[j].Name == i.Name && p.Inventory[j].Quantity >= n {
			return true
		}
	}
	return false
}
