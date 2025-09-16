package main

import "fmt"

var answer1 int 
var answer2 string
var UseItem bool = false 

/*#############################
  # AFFICHAGE DE L'INVENTAIRE #
  #############################*/
func accessInventory(FirstCharacter Character) {
	fmt.Println("===============INVENTAIRE================")
	for i := 0; i < len(FirstCharacter.Inventory); i++ {
		fmt.Println(i, ":", FirstCharacter.Inventory[i].Name,", quantité :",FirstCharacter.Inventory[i].Quantity)
	}


	/*######################
	  # GESTION DES OBJETS #
	  ######################*/
	fmt.Println("================ACTIONS==================")
	fmt.Println("tapez un numéro pour utiliser un objet, sinon tapez \"exit\"")
	fmt.Scanln(&answer1)
	switch answer1 {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:	
		if FirstCharacter.Inventory[answer1].Type == "Consumable" {
			fmt.Println("êtes-vous certain de vouloir consommer un(e)", FirstCharacter.Inventory[answer1].Name, "? [O/N]")
			fmt.Scanln(&answer2)
			switch answer2 {
			case "O", "o", "Oui", "oui":
				FirstCharacter.Inventory[answer1].Quantity -= 1
				fmt.Println("consommé")
				accessInventory(FirstCharacter)
				TakePot(FirstCharacter) // lien vers la fonction TakePot
			case "N", "n", "non", "no", "No", "Non":
				fmt.Println("annulé")
				accessInventory(FirstCharacter)
			default:
				fmt.Println("Erreur, Réponse non appropriée -> Retour au menu précédent")
			}
		} else if FirstCharacter.Inventory[answer1].Type == "Equippable" {
			// il faut faire la même chose pour les objets équipalbe comme outil, arme, arumure
		} else {
			fmt.Println("Erreur, l'objet sélectionné ne peut pas être utilisé")
		}
	case "exit":
		return
	default:
		fmt.Println("Erreur, Réponse non prise en charge")
	}
}
