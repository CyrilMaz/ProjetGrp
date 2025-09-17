package main

import "fmt"

var PotionGratuite = true // Variable globale pour suivre si la potion gratuite a été prise

func Merchand(FirstCharacter *Character) {
	for spaces := 0; spaces < 40; spaces++ {
		fmt.Print("\n")
	}
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡠⢔⢒⡿⠯⠥⢦⣦⣾⣄⠀⠀⠀⠀⠀⠀⠀")
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣾⢮⠊⠁⠀⠀⠀⠀⠈⠉⠛⠳⡀⠀⠀⠀⠀⠀")
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣧⣿⣝⡴⡔⠀⠀⠀⠀⠀⠀⠀⠀⠘⡀⠀⠀⠀⠀")
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⣀⣀⣦⣶⣿⣿⣯⣿⢽⠁⢰⣢⣶⣦⣌⠠⠴⠆⠘⣀⠀⠀⠀")
	fmt.Println("⠀⠀⠀⠀⢀⠔⠁⠀⢂⠘⢻⢛⣛⠿⣝⠁⠀⠼⣁⡴⣖⣫⠙⠙⠿⡳⡅⠀⠀⠀")
	fmt.Println("⠀⠀⠀⢀⠂⠰⠀⠀⠈⡄⢠⢓⣺⢇⡇⣊⠐⠀⠉⠁⠲⠒⠀⠀⠀⠑⠅⠀⠀⠀")
	fmt.Println("⠀⠀⠀⣨⠀⡇⠀⠀⠀⠰⢸⠄⠄⣸⣷⡦⣄⢤⠄⢄⡀⣀⠤⠠⡀⠀⠈⡄⠀⠀")
	fmt.Println("⠀⠀⢠⠁⠙⣇⠀⠀⠀⠀⢾⠘⢠⣿⣟⣿⣿⣪⣮⣶⣸⣮⣖⣢⣌⠁⠀⠁⠀⠀")
	fmt.Println("⠀⠀⢸⠀⠀⢹⠀⠀⠀⠀⠸⠀⢝⣻⣯⣿⢿⡫⢺⡩⠍⣉⣉⣨⡗⠉⠂⠊⠀⠀")
	fmt.Println("⠀⠀⡈⠂⠀⣾⠀⠀⠀⠀⠀⠆⡸⣹⡿⣿⣯⣷⣱⣙⠫⠧⠷⣦⠀⠀⠀⠀⠀⠀")
	fmt.Println("⠀⢠⠇⠀⠀⢉⠀⠀⠀⠀⠀⠈⠕⣻⣿⣿⣿⣟⣿⣿⡷⣶⣾⠿⠒⡁⠀⠀⠀⠀")
	fmt.Println("⠀⠀⠀⠀⠀⠈⡇⠀⠀⠀⠀⠀⠈⢇⠻⠽⣿⡿⢟⣿⣻⡟⠁⠰⣯⡕⡰⠀⠀⠀")
	fmt.Println("⠀⠀⠀⠀⠀⠀⢩⠀⠀⠀⠀⠀⢀⠬⣍⠱⠨⠯⠛⠙⢏⠀⢀⡀⣨⡀⠤⢚⠀⠀")
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⢂⠀⠀⠀⠀⠀⢠⠁⠀⠀⠀⠀⠀⢸⠄⠰⡶⠲⢦⠓⠍⠀⠀")
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠡⡀⠀⠀⠀⠀⠀⠀⡄⠀⢀⠄⠊⠉⠙⠑⠒⠊⠉⠀⠀")

	fmt.Println(" ┈ ┈ ┈ ┈ ┈ ⋞ 〈  ⏣ 	〉 ⋟ ┈ ┈ ┈ ┈ ┈")
	fmt.Println("   Bienvenue chez le marchand !")

	fmt.Println("¸„.-•~¹°”ˆ˜¨ BOUTIQUE ¨˜ˆ”°¹~•-.„¸")

	if PotionGratuite {
		fmt.Println("1 : Potion de soin (Gratuite)")
	} else {
		fmt.Println("1 : Potion de soin (3 pièces d'or)")
	}
	fmt.Println("2 : potion de poison (6 pièces d'or)")
	fmt.Println("3 : Livre de sorts (25 pièces d'or)")
	fmt.Println("4 : Fourrure de loup (4 pièces d'or)")
	fmt.Println("5 : Peau de troll (7 pièces d'or)")
	fmt.Println("6 : Cuir de sanglier (3 pièces d'or)")
	fmt.Println("7 : plume de corbeau (1 pièces d'or)")
	fmt.Println("\n0 : sortir de la boutique")

	var choice int
	fmt.Scanln(&choice) // Lire le choix de l'utilisateur

	switch choice {
	case 1:
		if PotionGratuite {
			fmt.Println("Vous avez reçu une potion de soin gratuite !")
			AddInventory(FirstCharacter, PotionSoin) // Ajoute la potion de soin à l'inventaire via la fonction AddInventory
			PotionGratuite = false                   // Met à jour la variable pour indiquer que la potion gratuite a été prise
		} else {
			if FirstCharacter.Gold >= 3 {
				AddInventory(FirstCharacter, PotionSoin) // Ajoute la potion de soin à l'inventaire via la fonction AddInventory
				FirstCharacter.Gold -= 3                 // Déduit le coût de la potion de l'argent du personnage
				fmt.Println("Vous avez acheté une potion de soin.")
			} else {
				fmt.Println("Vous n'avez pas assez d'argent pour acheter une potion de soin.")
			}
		}
	case 2:
		if FirstCharacter.Gold >= 6 {
			AddInventory(FirstCharacter, PotionPoison)
			FirstCharacter.Gold -= 6
			fmt.Println("Vous avez acheté une potion de poison.")
		} else {
			fmt.Println("Vous n'avez pas assez d'argent pour acheter une potion de poison.")
		}
	case 3:
		if FirstCharacter.Gold >= 25 {
			AddInventory((FirstCharacter), LivreSorts)
			FirstCharacter.Gold -= 25
			fmt.Println("Vous avez acheté un livre de sorts.")
		}
	case 4:
		if FirstCharacter.Gold >= 4 {
			AddInventory(FirstCharacter, FourrureLoup)
			FirstCharacter.Gold -= 4
			fmt.Println("Vous avez acheté une fourrure de loup.")
		} else {
			fmt.Println("Vous n'avez pas assez d'argent pour acheter une fourrure de loup.")
		}
	case 5:
		if FirstCharacter.Gold >= 7 {
			AddInventory(FirstCharacter, PeauTroll)
			FirstCharacter.Gold -= 7
			fmt.Println("Vous avez acheté une peau de troll.")
		} else {
			fmt.Println("Vous n'avez pas assez d'argent pour acheter une peau de troll.")
		}
	case 6:
		if FirstCharacter.Gold >= 3 {
			AddInventory(FirstCharacter, CuirSanglier)
			FirstCharacter.Gold -= 3
			fmt.Println("Vous avez acheté un cuir de sanglier.")
		} else {
			fmt.Println("Vous n'avez pas assez d'argent pour acheter un cuir de sanglier.")
		}
	case 7:
		if FirstCharacter.Gold >= 1 {
			AddInventory(FirstCharacter, PlumeCorbeau)
			FirstCharacter.Gold -= 1
			fmt.Println("Vous avez acheté une plume de corbeau.")
		} else {
			fmt.Println("Vous n'avez pas assez d'argent pour acheter une plume de corbeau.")
		}
	case 0:
		fmt.Println("Merci de votre visite !")
		main() // Retour au menu principal
	default: // fonction default sert à gérer les cas non prévus
		fmt.Println("Choix invalide, réessayez.") // Message d'erreur pour un choix invalide
		fmt.Scanln(&choice)                       // Lire le choix de l'utilisateur
	}
}
