package main

import "fmt"

var answerRevive string

func IsDead(p *Character) {
	if p.Pv <= 0 {
		fmt.Println("vous êtes mort")
		fmt.Scanln(&answerRevive)
		switch answerRevive {
		case "O", "o", "Oui", "oui":
			fmt.Println("vous avez été ressuscité avec la moitié de vos points de vie")
			p.Pv = p.Pvmax / 2
			main() // lien vers la fonction main
		case "N", "n", "non", "no", "No", "Non":
			fmt.Println("vous avez choisi de ne pas être ressuscité, fin du jeu")
			return
		default:
			fmt.Println("réponse invalide, vous avez été ressuscité avec la moitié de vos points de vie")
			main() // lien vers la fonction main
		}

	}
}
