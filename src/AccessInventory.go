package main

import "fmt"

var answer1 int
var answer2 string
var UseItem bool = false 

/*#############################
# AFFICHAGE DE L'INVENTAIRE #
#############################*/
func accessInventory(FirstCharacter *Character) {
	fmt.Println("\n")
	fmt.Println("◇─◇──◇────◇ INVENTAIRE ◇─────◇──◇─◇")
	for i := 1; i <= len(FirstCharacter.Inventory); i++ {
		fmt.Println(i, ":", FirstCharacter.Inventory[i - 1].Name,", quantité :",FirstCharacter.Inventory[i - 1].Quantity)
	}


	/*######################
	# GESTION DES OBJETS #
	######################*/
	fmt.Println("\n")
	fmt.Println("◇─◇──◇────◇ ACTIONS ◇─────◇──◇─◇")
	fmt.Println("0 revenir au menu précédent")
	fmt.Println("tapez un numéro pour utiliser un objet ou pour faire une action")
	fmt.Scanln(&answer1)

	switch answer1 {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
		if FirstCharacter.Inventory[answer1 - 1].Type == "Consumable" {
			fmt.Println("êtes-vous certain de vouloir consommer un(e)", FirstCharacter.Inventory[answer1 - 1].Name, "? [O/N]")
			fmt.Scanln(&answer2)
			switch answer2 {
			case "O", "o", "Oui", "oui":
				if FirstCharacter.Inventory[answer1 - 1].Quantity > 0 {
					FirstCharacter.Inventory[answer1 - 1].Quantity -= 1
					fmt.Println("consommé")
					TakePot(FirstCharacter) // lien vers la fonction TakePot
					accessInventory(FirstCharacter)
				} else {
					fmt.Println("Pas assez de", FirstCharacter.Inventory[answer1 - 1].Name)
					accessInventory(FirstCharacter)
				}
			case "N", "n", "non", "no", "No", "Non":
				fmt.Println("annulé")
				accessInventory(FirstCharacter)
			default:
				fmt.Println("Retour au menu précédent")
			}
		} else if FirstCharacter.Inventory[answer1 - 1].Type == "Equippable" {
			// il faut faire la même chose pour les objets équipalbe comme outil, arme, arumure
		} else {
			fmt.Println("Erreur, l'objet sélectionné ne peut pas être utilisé")
		}
	case 0:
		return 
	default:
		return
	}
}
