package main

import "fmt"

func Merchand() {
	fmt.Println("Bienvenue chez le marchand !")
	fmt.Println("Que souhaitez-vous acheter ?")
	fmt.Println("1 : Potion de soin (20 pièces d'or)")
	fmt.Println("0 : Quitter le marchand")

	var choice int 
	fmt.Scanln(&choice) // Lire le choix de l'utilisateur

	switch choice {
	case 1:
		fmt.Println("Vous avez acheté une potion de soin.")
		FirstCharacter.Inventory = append(FirstCharacter.Inventory, PotionSoin) //Ajoute la potion de soin à l'inventaire via la fonction append
		// Ici, vous pouvez ajouter la logique pour déduire l'argent du personnage
	case 0:
		fmt.Println("Merci de votre visite !")//
		main() // Retour au menu principal
	default: // fonction default sert à gérer les cas non prévus
		fmt.Println("Choix invalide, réessayez.")// 
	
	}
