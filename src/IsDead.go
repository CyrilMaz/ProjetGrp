package main

import (
	"fmt"
	"os"
	"time"
)

var answerRevive string

func IsDead(p *Character) {
	for spaces := 0; spaces < 40; spaces++ {
		fmt.Print("\n")
	}
	fmt.Println("vous êtes mort")
	fmt.Println("★ voulez-vous être ressuscité ? (O/N)")
	fmt.Scanln(&answerRevive)
	switch answerRevive {
	case "o", "O", "ou", "oU", "Ou", "OU", "oui", "ouI", "oUi", "oUI", "Oui", "OuI", "OUi", "OUI":
		for spaces := 0; spaces < 40; spaces++ {
			fmt.Print("\n")
		}
		fmt.Print("Résurrection en cours : ")
		loading := []string{"▂", "▃", "▄", "▅", "▆", "▇", "▉"}
		for i := 0; i < 7; i++ { // 8*0.50s ≈ 4s
			fmt.Print(loading[i%len(loading)])
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println()
		fmt.Println("vous avez été ressuscité avec la moitié de vos points de vie")
		(*p).Pvmax /= 2
		main() // lien vers la fonction main

	case "N", "n", "no", "nO", "No", "NO", "non", "noN", "nOn", "nON", "Non", "NoN", "NOn", "NON":
		for spaces := 0; spaces < 40; spaces++ {
			fmt.Print("\n")
		}
		fmt.Print("Fin du jeu : ")
		loading := []string{"▉", "▇", "▆", "▅", "▄", "▃", "▂"}
		for i := 0; i < 7; i++ {
			fmt.Print(loading[i%len(loading)])
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println("\nvous avez choisi de ne pas être ressuscité, fin du jeu")
		os.Exit(0)

	default:
		fmt.Println("réponse invalide, vous avez été ressuscité avec la moitié de vos points de vie")
		main() // lien vers la fonction main
	}
}
