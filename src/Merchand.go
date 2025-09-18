package main

import "fmt"

var PotionGratuite = true // Variable globale pour suivre si la potion gratuite a été prise
var DisplayAddItem = true
var ShowlastPurchase = false
var lastPurchase = ""
var c int = 0

func Merchand(p *Character) {
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

	fmt.Println(" ┈ ┈ ┈ ┈ ┈ ⋞ <  ⏣  > ⋟ ┈ ┈ ┈ ┈ ┈")
	fmt.Println("\n   Bienvenue chez le marchand !")

	fmt.Println("\n¸„.-•~¹°”ˆ˜¨ BOUTIQUE ¨˜ˆ”°¹~•-.„¸")

	if PotionGratuite {
		fmt.Println("| 1 : Potion de soin ···(Gratuite)")
	} else {
		fmt.Println("| 1 : Potion de soin ·······(3 PO)")
	}
	fmt.Println("| 2 : Potion de poison ·····(6 PO)")
	fmt.Println("| 3 : Livre de sorts ······(25 PO)")
	fmt.Println("| 4 : Fourrure de loup ·····(4 PO)")
	fmt.Println("| 5 : Peau de troll ········(7 PO)")
	fmt.Println("| 6 : Cuir de sanglier ·····(3 PO)")
	fmt.Println("| 7 : Plume de corbeau ·····(1 PO)")
	fmt.Println("| 8 : +10 slots inventaire·(30 PO)")
	fmt.Println("|                                |")
	fmt.Println("★━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━★")
	fmt.Println("|    0 : Sortir de la boutique   |")
	fmt.Println("★━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━★")
	if lasterror != "" {
		fmt.Println("Erreur :", lasterror)
		lasterror = ""
	} else if ShowlastPurchase {
		fmt.Println("Dernier achat :", lastPurchase)
		ShowlastPurchase = false
	}
	var choice int
	fmt.Scanln(&choice) // Lire le choix de l'utilisateur

	switch choice {
	case 1:
		ShowlastPurchase = true
		if PotionGratuite {
			DisplayAddItem = false
			AddInventory(p, PotionSoin) // Ajoute la potion de soin à l'inventaire via la fonction AddInventory
			lastPurchase = fmt.Sprintf("Vous avez ajouté %d %s à votre inventaire.", PotionSoin.Quantity, PotionSoin.Name)
			PotionGratuite = false // Met à jour la variable pour indiquer que la potion gratuite a été prise
			Merchand(p)
		} else {
			if FirstCharacter.Gold >= 3 {
				DisplayAddItem = false
				AddInventory(p, PotionSoin) // Ajoute la potion de soin à l'inventaire via la fonction AddInventory
				FirstCharacter.Gold -= 3    // Déduit le coût de la potion de l'argent du personnage
				lastPurchase = fmt.Sprintf("Vous avez ajouté %d %s à votre inventaire.", PotionSoin.Quantity, PotionSoin.Name)
				Merchand(p)
				fmt.Println("Vous avez ajouté", PotionSoin.Quantity, PotionSoin.Name, "à votre inventaire.")
			} else {
				lasterror = "Vous n'avez pas assez d'argent pour acheter une potion de soin."
				Merchand(p)
			}
		}
	case 2:
		ShowlastPurchase = true
		if FirstCharacter.Gold >= 6 {
			DisplayAddItem = false
			AddInventory(p, PotionPoison)
			FirstCharacter.Gold -= 6
			lastPurchase = fmt.Sprintf("Vous avez ajouté %d %s à votre inventaire.", PotionPoison.Quantity, PotionPoison.Name)
			Merchand(p)
			fmt.Println("Vous avez ajouté", PotionPoison.Quantity, PotionPoison.Name, "à votre inventaire.")
		} else {
			lasterror = "Vous n'avez pas assez d'argent pour acheter une potion de poison."
			Merchand(p)
		}
	case 3:
		ShowlastPurchase = true
		if FirstCharacter.Gold >= 25 {
			DisplayAddItem = false
			AddInventory((p), LivreSorts)
			FirstCharacter.Gold -= 25
			lastPurchase = fmt.Sprintf("Vous avez ajouté %d %s à votre inventaire.", LivreSorts.Quantity, LivreSorts.Name)
			Merchand(p)
			fmt.Println("Vous avez ajouté", LivreSorts.Quantity, LivreSorts.Name, "à votre inventaire.")
		} else {
			lasterror = "Vous n'avez pas assez d'argent pour acheter un livre de sorts."
			Merchand(p)
		}
	case 4:
		ShowlastPurchase = true
		if FirstCharacter.Gold >= 4 {
			DisplayAddItem = false
			AddInventory(p, FourrureLoup)
			FirstCharacter.Gold -= 4
			lastPurchase = fmt.Sprintf("Vous avez ajouté %d %s à votre inventaire.", FourrureLoup.Quantity, FourrureLoup.Name)
			Merchand(p)
			fmt.Println("Vous avez ajouté", FourrureLoup.Quantity, FourrureLoup.Name, "à votre inventaire.")
		} else {
			lasterror = "Vous n'avez pas assez d'argent pour acheter une fourrure de loup."
			Merchand(p)
		}
	case 5:
		ShowlastPurchase = true
		if FirstCharacter.Gold >= 7 {
			DisplayAddItem = false
			AddInventory(p, PeauTroll)
			FirstCharacter.Gold -= 7
			lastPurchase = fmt.Sprintf("Vous avez ajouté %d %s à votre inventaire.", PeauTroll.Quantity, PeauTroll.Name)
			Merchand(p)
			fmt.Println("Vous avez ajouté", PeauTroll.Quantity, PeauTroll.Name, "à votre inventaire.")
		} else {
			lasterror = "Vous n'avez pas assez d'argent pour acheter une peau de troll."
			Merchand(p)
		}
	case 6:
		ShowlastPurchase = true
		if FirstCharacter.Gold >= 3 {
			DisplayAddItem = false
			AddInventory(p, CuirSanglier)
			FirstCharacter.Gold -= 3
			lastPurchase = fmt.Sprintf("Vous avez ajouté %d %s à votre inventaire.", CuirSanglier.Quantity, CuirSanglier.Name)
			Merchand(p)
			fmt.Println("Vous avez ajouté", CuirSanglier.Quantity, CuirSanglier.Name, "à votre inventaire.")
		} else {
			lasterror = "Vous n'avez pas assez d'argent pour acheter un cuir de sanglier."
			Merchand(p)
		}
	case 7:
		ShowlastPurchase = true
		if FirstCharacter.Gold >= 1 {
			DisplayAddItem = false
			AddInventory(p, PlumeCorbeau)
			FirstCharacter.Gold -= 1
			lastPurchase = fmt.Sprintf("Vous avez ajouté %d %s à votre inventaire.", PlumeCorbeau.Quantity, PlumeCorbeau.Name)
			Merchand(p)
			fmt.Println("Vous avez ajouté", PlumeCorbeau.Quantity, PlumeCorbeau.Name, "à votre inventaire.")
		} else {
			lasterror = "Vous n'avez pas assez d'argent pour acheter une plume de corbeau."
			Merchand(p)
		}
	case 8:
		ShowlastPurchase = true
		if FirstCharacter.Gold >= 30 {
			if c < 3 {
				DisplayAddItem = false
				UpgradeInventorySlot(p)
				c++
				FirstCharacter.Gold -= 30
				lastPurchase = "Vous avez augmenté votre inventaire de 10 slots."
				Merchand(p)
			} else if c >= 3 {
				lasterror = "Vous avez déjà acheté le maximum de slots d'inventaire (3)."
				Merchand(p)
			}
		} else {
			lasterror = "Vous n'avez pas assez d'argent pour acheter des slots d'inventaire."
			Merchand(p)

		}
	case 0:
		fmt.Println("Merci de votre visite !")
		main() // Retour au menu principal
	default: // fonction default sert à gérer les cas non prévus
		lasterror = "Choix invalide, réessayez." // Message d'erreur pour un choix invalide
		Merchand(p)
	}
}
