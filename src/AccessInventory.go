package main

import "fmt"

var answer1 int
var answer2 string
var UseItem bool = false

/*
#############################
# AFFICHAGE DE L'INVENTAIRE #
#############################
*/
func accessInventory(FirstCharacter *Character) {
	for spaces := 0; spaces < 40; spaces++ {
		fmt.Print("\n")
	}
	fmt.Println("┏◇─◇──◇────◇ INVENTAIRE ◇─────◇──◇─◇┓")
	fmt.Println("╠===================================╣")
	for i := 1; i <= len(FirstCharacter.Inventory); i++ {
		if FirstCharacter.Inventory[i-1].Quantity > 0 {
			fmt.Println(" ", i, ":", FirstCharacter.Inventory[i-1].Name, "| quantité:", FirstCharacter.Inventory[i-1].Quantity)
		}
	}
	fmt.Println("┗===================================┛")

	/*######################
	# GESTION DES OBJETS #
	######################*/
	fmt.Println("")
	fmt.Println("\n")
	fmt.Println("★¸„.-•~¹°”ˆ˜¨ ACTIONS ¨˜ˆ”°¹~•-.„¸★")
	fmt.Println("|                                 |")
	fmt.Println("★━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━★")
	fmt.Println("|   0 revenir au menu principal   |")
	fmt.Println("★━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━★")
	fmt.Println("\n★ tapez le numéro d'un objet pour l'utiliser")
	fmt.Scanln(&answer1)

	switch answer1 {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
		if FirstCharacter.Inventory[answer1-1].Type == "Consumable" {
			fmt.Println("★ êtes-vous certain de vouloir consommer un(e)", FirstCharacter.Inventory[answer1-1].Name, "? [O/N]")
			fmt.Scanln(&answer2)
			switch answer2 {
			case "o", "O", "ou", "oU", "Ou", "OU", "oui", "ouI", "oUi", "oUI", "Oui", "OuI", "OUi", "OUI":
				if FirstCharacter.Inventory[answer1-1].Name == "potions de soin" {
					if FirstCharacter.Inventory[answer1-1].Quantity > 0 {
						FirstCharacter.Inventory[answer1-1].Quantity -= 1
						fmt.Println("consommé")
						TakePot(FirstCharacter) // lien vers la fonction TakePot
						accessInventory(FirstCharacter)
					} else {
						fmt.Println("Pas assez de", FirstCharacter.Inventory[answer1-1].Name)
						accessInventory(FirstCharacter)
					}
				} else if FirstCharacter.Inventory[answer1-1].Name == "Livre de sorts" {
					if FirstCharacter.Inventory[answer1-1].Quantity > 0 {
						FirstCharacter.Inventory[answer1-1].Quantity -= 1
						fmt.Println("consommé")
						SpeellBook(FirstCharacter, Fireball) // lien vers la fonction SpellBook
					} else {
						fmt.Println("Pas assez de", FirstCharacter.Inventory[answer1-1].Name)
						accessInventory(FirstCharacter)
					}

				} else if FirstCharacter.Inventory[answer1-1].Name == "Potion de poison" {
					if FirstCharacter.Inventory[answer1-1].Quantity > 0 {
						FirstCharacter.Inventory[answer1-1].Quantity -= 1
						fmt.Println("consommé")
						PoisonPot(FirstCharacter)
						fmt.Println("\n★ Appuyez sur entrée pour continuer...")
						fmt.Scanln()
						accessInventory(FirstCharacter)
					} else {
						fmt.Println("Pas assez de", FirstCharacter.Inventory[answer1-1].Name)
						accessInventory(FirstCharacter)
					}

				} else {
					fmt.Println("Erreur, l'objet sélectionné ne peut pas être utilisé")
				}
			case "N", "n", "no", "nO", "No", "NO", "non", "noN", "nOn", "nON", "Non", "NoN", "NOn", "NON":
				fmt.Println("annulé")
				accessInventory(FirstCharacter)
			default:
				fmt.Println("Retour au menu précédent")
			}
		} else if FirstCharacter.Inventory[answer1-1].Type == "Equippable" {
			// il faut faire la même chose pour les objets équipalbe comme outil, arme, arumure
		} else {
			fmt.Println("Erreur, l'objet sélectionné ne peut pas être utilisé")
		}
	case 0:
		main()
	default:
		main()
	}
}
