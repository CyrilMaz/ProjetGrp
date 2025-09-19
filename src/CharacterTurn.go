package main

import "fmt"

var choice int
var PlayerActionTour string

func CharacterTurn(p *Character) string {

	fmt.Println("★¸„.-•~¹°”ˆ˜¨ ACTIONS ¨˜ˆ”°¹~•-.„¸★")
	fmt.Println("|         C'est ton tour !        |")
	fmt.Println("★━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━★")
	fmt.Println("|1:··········Attaquer             |")
	fmt.Println("|2:······Voir l'inventaire        |")
	fmt.Println("|                                 |")
	fmt.Println("★━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━★")
	fmt.Println("|       0 : Fuir du combat        |")
	fmt.Println("★━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━★")
	if lasterror != "" {
		fmt.Println("Erreur :", lasterror)
		lasterror = ""
	}

	fmt.Scanln(&choice) // Lire le choix de l'utilisateur

	switch choice {
	case 1:
		fmt.Println(p.Name, "a décidé d'attaquer")
		PlayerActionTour = "1"
		// faire la fonction d'attaque ici
	case 2:
		fmt.Println(p.Name, "a ouvert son inventaire")
		PlayerActionTour = "2"
	case 0:
		fmt.Println(p.Name, "a décidé de prendre la fuite")
		PlayerActionTour = "0"
		// Logique de passage de tour ici
	default:
		lasterror = "Choix invalide, veuillez réessayer."
		CharacterTurn(p) // Re-demande le choix si invalide
	}
	return ""
}
