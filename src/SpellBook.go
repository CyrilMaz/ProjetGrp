package main

import "fmt"

func SpeellBook(p *Character, s Skills) {
<<<<<<< HEAD
	for i := 0; i < len(p.Skills); i++ {
		if p.Skills[i].Name == s.Name {
=======
	for i := 0; i > len((*p).Skills); i++ {
		if (*p).Skills[i].Name == s.Name {
			s.learned = true
			fmt.Println("Vous avez appris le sort :", s.Name)
			(*p).Skills = append((*p).Skills, s)
		} else {
>>>>>>> 49dad57a2d0bffb6659c2e8104131526b708fd94
			fmt.Println("Vous avez déjà appris le sort :", s.Name)
			return
		} else {
			s.Learned = true
			p.Skills = append(p.Skills, s)
			fmt.Println("Vous avez appris le sort :", s.Name)
			Merchand(p)
			return
		}
	}
}
