package main

import "fmt"

func SpeellBook(p *Character, s Skills) {
	for i := 0; i < len(p.Skills); i++ {
		if p.Skills[i].Name == s.Name {
			fmt.Println("Vous avez déjà appris le sort :", s.Name)
		} else {
			s.Learned = true
			p.Skills = append(p.Skills, s)
			accessInventory(p)
			fmt.Println("Vous avez appris le sort :", s.Name)
		}
	}
}
