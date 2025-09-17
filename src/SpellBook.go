package main

import "fmt"

func SpeellBook(p *Character, s Skills) {
	for i := 0; i < len((*p).Skills); i++ {
		if (*p).Skills[i].Name == s.Name {
			s.learned = true
			fmt.Println("Vous avez appris le sort :", s.Name)
			(*p).Skills = append((*p).Skills, s)
		} else {
			fmt.Println("Vous avez déjà appris le sort :", s.Name)
		}
	}
}
