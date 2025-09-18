package main

import "fmt"

var answer1 int
var answer2 string
var answer3 string
var UseItem bool = false

/*
#############################
# AFFICHAGE DE L'INVENTAIRE #
#############################
*/
func accessInventory(p *Character) {
	for spaces := 0; spaces < 40; spaces++ {
		fmt.Print("\n")
	}
	fmt.Println("┏◇─◇──◇────◇ INVENTAIRE ◇─────◇──◇─◇┓")
	fmt.Println("╠===================================╣")
	for i := 1; i <= len(p.Inventory); i++ {
		if p.Inventory[i-1].Quantity > 0 {
			fmt.Println(" ", i, ":", p.Inventory[i-1].Name, "| quantité:", p.Inventory[i-1].Quantity)
		}
	}
	fmt.Println("┗===================================┛")

	/*######################
	# GESTION DES OBJETS #
	######################*/
	fmt.Println("")
	fmt.Println("")
	fmt.Println("★¸„.-•~¹°”ˆ˜¨ ACTIONS ¨˜ˆ”°¹~•-.„¸★")
	fmt.Println("|                                 |")
	fmt.Println("★━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━★")
	fmt.Println("|   0 revenir au menu principal   |")
	fmt.Println("★━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━★")
	fmt.Println("\n★ tapez le numéro d'un objet pour l'utiliser")
	fmt.Scanln(&answer1)

	switch answer1 {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
		if p.Inventory[answer1-1].Type == "Consumable" {
			fmt.Println("★ êtes-vous certain de vouloir consommer un(e)", p.Inventory[answer1-1].Name, "? [O/N]")
			fmt.Scanln(&answer2)
			switch answer2 {
			case "o", "O", "ou", "oU", "Ou", "OU", "oui", "ouI", "oUi", "oUI", "Oui", "OuI", "OUi", "OUI":
				if p.Inventory[answer1-1].Name == "potions de soin" {
					if p.Inventory[answer1-1].Quantity > 0 {
						p.Inventory[answer1-1].Quantity -= 1
						fmt.Println("consommé")
						TakePot(p) // lien vers la fonction TakePot
						accessInventory(p)
					} else {
						fmt.Println("Pas assez de", p.Inventory[answer1-1].Name)
						accessInventory(p)
					}
				} else if p.Inventory[answer1-1].Name == "Livre de sorts" {
					if p.Inventory[answer1-1].Quantity > 0 {
						p.Inventory[answer1-1].Quantity -= 1
						fmt.Println("consommé")
						SpeellBook(p, Fireball) // lien vers la fonction SpellBook
					} else {
						fmt.Println("Pas assez de", p.Inventory[answer1-1].Name)
						accessInventory(p)
					}

				} else if p.Inventory[answer1-1].Name == "Potion de poison" {
					if p.Inventory[answer1-1].Quantity > 0 {
						p.Inventory[answer1-1].Quantity -= 1
						fmt.Println("consommé")
						PoisonPot(p)
						fmt.Println("\n★ Appuyez sur entrée pour continuer...")
						fmt.Scanln()
						accessInventory(p)
					} else {
						fmt.Println("Pas assez de", p.Inventory[answer1-1].Name)
						accessInventory(p)
					}

				} else {
					fmt.Println("Erreur, l'objet sélectionné ne peut pas être utilisé")
				}
			case "N", "n", "no", "nO", "No", "NO", "non", "noN", "nOn", "nON", "Non", "NoN", "NOn", "NON":
				fmt.Println("annulé")
				accessInventory(p)
			default:
				fmt.Println("Retour au menu précédent")
			}
		} else if p.Inventory[answer1-1].Type == "Equipment" {
			fmt.Println("★ êtes-vous certain de vouloir équiper :", p.Inventory[answer1-1].Name, "? [O/N]")
			fmt.Scanln(&answer2)
			switch answer2 {
			case "o", "O", "ou", "oU", "Ou", "OU", "oui", "ouI", "oUi", "oUI", "Oui", "OuI", "OUi", "OUI":

				/*#################################
				  # Si l'equipement est un casque #
				  #################################*/
				if p.Inventory[answer1-1].StatName == "Head" {
					if p.Equipment[0].Worn {
						fmt.Println("Un équipement est déjà sur la tête :", p.Equipment[0].Name)
						fmt.Println("Voulez-vous le remplacer ? [O/N]")
						fmt.Scanln(&answer3)
						switch answer3 {
						case "o", "O", "ou", "oU", "Ou", "OU", "oui", "ouI", "oUi", "oUI", "Oui", "OuI", "OUi", "OUI":
							AddInventory(p, Items{p.Equipment[0].Name, 1, "Equipment", 0, p.Inventory[answer1-1].StatName, "Equipment"}) // remet l'ancien équipement dans l'inventaire
							fmt.Println("Ancien équipement ajouté à l'inventaire :", p.Equipment[0].Name)
							p.Equipment[0].Name = p.Inventory[answer1-1].Name
							RemoveInventory(p, p.Inventory[answer1-1], 1)
							fmt.Println("Vous avez équipé :", p.Inventory[answer1-1].Name)
							accessInventory(p)
						case "N", "n", "no", "nO", "No", "NO", "non", "noN", "nOn", "nON", "Non", "NoN", "NOn", "NON":
							fmt.Println("annulé")
							accessInventory(p)
						}
					} else {
						p.Equipment[0].Worn = true
						p.Equipment[0].Name = p.Inventory[answer1-1].Name
						fmt.Println("Vous avez équipé :", p.Inventory[answer1-1].Name)
						accessInventory(p)
					}

					/*##################################
					  # Si l'equipement est une armure #
					  ##################################*/
				} else if p.Inventory[answer1-1].StatName == "body" {
					if p.Equipment[1].Worn {
						fmt.Println("Une armure est déjà équipée :", p.Equipment[1].Name)
						fmt.Println("Voulez-vous la remplacer ? [O/N]")
						fmt.Scanln(&answer3)
						switch answer3 {
						case "o", "O", "ou", "oU", "Ou", "OU", "oui", "ouI", "oUi", "oUI", "Oui", "OuI", "OUi", "OUI":
							AddInventory(p, Items{p.Equipment[1].Name, 1, "Equipment", 0, p.Inventory[answer1-1].StatName, "Equipment"})
							fmt.Println("Ancienne armure ajoutée à l'inventaire :", p.Equipment[1].Name)
							p.Equipment[1].Name = p.Inventory[answer1-1].Name
							RemoveInventory(p, p.Inventory[answer1-1], 1)
							fmt.Println("Vous avez équipé :", p.Inventory[answer1-1].Name)
							accessInventory(p)
						case "N", "n", "no", "nO", "No", "NO", "non", "noN", "nOn", "nON", "Non", "NoN", "NOn", "NON":
							fmt.Println("annulé")
							accessInventory(p)
						}
					} else {
						p.Equipment[1].Worn = true
						p.Equipment[1].Name = p.Inventory[answer1-1].Name
						fmt.Println("Vous avez équipé :", p.Inventory[answer1-1].Name)
						accessInventory(p)
					}

					/*###################################
					  # Si l'equipement est un pantalon #
					  ###################################*/
				} else if p.Inventory[answer1-1].StatName == "legs" {
					if p.Equipment[2].Worn {
						fmt.Println("Un pantalon est déjà équipé :", p.Equipment[2].Name)
						fmt.Println("Voulez-vous le remplacer ? [O/N]")
						fmt.Scanln(&answer3)
						switch answer3 {
						case "o", "O", "ou", "oU", "Ou", "OU", "oui", "ouI", "oUi", "oUI", "Oui", "OuI", "OUi", "OUI":
							AddInventory(p, Items{p.Equipment[2].Name, 1, "Equipment", 0, p.Inventory[answer1-1].StatName, "Equipment"})
							fmt.Println("Ancien pantalon ajouté à l'inventaire :", p.Equipment[2].Name)
							p.Equipment[2].Name = p.Inventory[answer1-1].Name
							RemoveInventory(p, p.Inventory[answer1-1], 1)
							fmt.Println("Vous avez équipé :", p.Inventory[answer1-1].Name)
							accessInventory(p)
						case "N", "n", "no", "nO", "No", "NO", "non", "noN", "nOn", "nON", "Non", "NoN", "NOn", "NON":
							fmt.Println("annulé")
							accessInventory(p)
						}
					} else {
						p.Equipment[2].Worn = true
						p.Equipment[2].Name = p.Inventory[answer1-1].Name
						fmt.Println("Vous avez équipé :", p.Inventory[answer1-1].Name)
						accessInventory(p)
					}
				}
			case "N", "n", "no", "nO", "No", "NO", "non", "noN", "nOn", "nON", "Non", "NoN", "NOn", "NON":
				fmt.Println("annulé")
				accessInventory(p)
			}

		} else {
			fmt.Println("Erreur, l'objet sélectionné ne peut pas être utilisé")
		}
	case 0:
		main()
	default:
		main()
	}
}
