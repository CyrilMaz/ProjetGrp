package main

import "fmt"

var PotionGratuite = true // Variable globale pour suivre si la potion gratuite a été prise

func Merchand(FirstCharacter *Character) {
	fmt.Println("Bienvenue chez le marchand !")
	fmt.Println("Que souhaitez-vous acheter ?")
	fmt.Println("1 : Potion de soin (20 Argent)")
	fmt.Println("0 : Quitter le marchand")

	var choice int
	fmt.Scanln(&choice) // Lire le choix de l'utilisateur

	switch choice {
	case 1:
		if PotionGratuite {
			fmt.Println("Vous avez reçu une potion de soin gratuite !")
			FirstCharacter.Inventory = append(FirstCharacter.Inventory, PotionSoin) //Ajoute la potion de soin à l'inventaire via la fonction append
			PotionGratuite = false                                                  // Met à jour la variable pour indiquer que la potion gratuite a été prise
		} else {
			if FirstCharacter.Gold >= 20 {
				FirstCharacter.Inventory = append(FirstCharacter.Inventory, PotionSoin) //Ajoute la potion de soin à l'inventaire via la fonction append
				FirstCharacter.Gold -= 20                                               // Déduit le coût de la potion de l'argent du personnage
				fmt.Println("Vous avez acheté une potion de soin.")
			} else {
				fmt.Println("Vous n'avez pas assez d'argent pour acheter une potion de soin.")
			}
		}
	case 0:
		fmt.Println("Merci de votre visite !")
		main() // Retour au menu principal
	default: // fonction default sert à gérer les cas non prévus
		fmt.Println("Choix invalide, réessayez.") // Message d'erreur pour un choix invalide
		fmt.Scanln(&choice)                       // Lire le choix de l'utilisateur
	}

}
