package main

import "fmt"

var AttackChoice string
var PlayerAttack Skills
var CompteurTour int

func trainingFight(p *Character, Goblin *Monster) {
	initGoblin()
	CompteurTour = 0
	for p.Pv > 0 && Goblin.Pv > 0 {
		fmt.Println("★━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━★")
		fmt.Println("|        Tour de jeu n°", CompteurTour, "      |")
		fmt.Println("★━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━★")
		if p.Pv > 0 {
			IsDead(p)
		} else if Goblin.Pv > 0 {
			fmt.Println("Vous avez gagné ! ")
			p.Gold += 50
			fmt.Println("Vous avez reçu 50 pièces d'or")
		}
		CharacterTurn(p)
		switch PlayerActionTour {
		case "attack":
			for AttackChoice != "1" || AttackChoice != "2" {
				for i := 1; i < len(FirstCharacter.Skills); i++ {
					if FirstCharacter.Skills[i-1].Learned {
						fmt.Println("", i, "☆", FirstCharacter.Skills[i-1].Name)
					}
				}
				fmt.Println("Quel attaque voulez vous faire ?")
				fmt.Println("0 : pour revenir en arrière")
				fmt.Scanln(&AttackChoice)
				switch AttackChoice {
				case "0":
					PlayerActionTour = ""
					CharacterTurn(p)
				case "1":
					PlayerAttack = CoupDePoing
					Goblin.Pv -= PlayerAttack.damage
				case "2":
					PlayerAttack = Fireball
					Goblin.Pv -= PlayerAttack.damage
				default:
					fmt.Println("Choix invalide, réessayez")
				}
				fmt.Sprintf("%s inflige %d point de dégat à %s", p.Name, PlayerAttack.damage, Goblin.Name)
				fmt.Sprintf("%s a désormais %d/%d HP", Goblin.Name, Goblin.Pv, Goblin.Pvmax)
			}
		case "Inventory":
			fmt.Println("")
			fmt.Println("")
			fmt.Println("★¸„.-•~¹°”ˆ˜¨ ACTIONS ¨˜ˆ”°¹~•-.„¸★")
			fmt.Println("|                                 |")
			fmt.Println("★━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━★")
			fmt.Println("|        0 revenir au combat      |")
			fmt.Println("★━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━★")
			if lasterror != "" {
				fmt.Println("Erreur :", lasterror)
				lasterror = ""
			} else if ShowlastAction {
				fmt.Println("Dernière action :", lastAction)
				lastAction = ""
				ShowlastAction = false
			} else if ShowlastAction2 {
				fmt.Println("Dernière action :", lastAction2)
				lastAction2 = ""
				ShowlastAction2 = false
			}
			fmt.Println("\n★ tapez le numéro d'un objet pour l'utiliser")
			fmt.Scanln(&answer1)

			switch answer1 {
			case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40:
				if p.Inventory[answer1-1].Type == "Consumable" {
					fmt.Println("★ êtes-vous certain de vouloir consommer un(e)", p.Inventory[answer1-1].Name, "? [O/N]")
					fmt.Scanln(&answer2)
					switch answer2 {
					case "o", "O", "ou", "oU", "Ou", "OU", "oui", "ouI", "oUi", "oUI", "Oui", "OuI", "OUi", "OUI":
						if p.Inventory[answer1-1].Name == "potions de soin" {
							if p.Inventory[answer1-1].Quantity > 0 {
								p.Inventory[answer1-1].Quantity -= 1
								lastAction = fmt.Sprintf("vous avez consommé 1 %s et récupéré %d points de vie", p.Inventory[answer1-1].Name, p.Inventory[answer1-1].StatBoost)
								ShowlastAction = true
								TakePot(p) // lien vers la fonction TakePot
								accessInventory(p)
							} else {
								lasterror = fmt.Sprintf("Pas assez de", p.Inventory[answer1-1].Name)
								accessInventory(p)
							}
						} else if p.Inventory[answer1-1].Name == "Livre de sorts" {
							if p.Inventory[answer1-1].Quantity > 0 {
								p.Inventory[answer1-1].Quantity -= 1
								lastAction = fmt.Sprintf("vous avez consommé 1 %s et appris le sort Boule de feu", p.Inventory[answer1-1].Name)
								ShowlastAction = true
								SpeellBook(p, Fireball) // lien vers la fonction SpellBook
								accessInventory(p)
							} else {
								lasterror = fmt.Sprintf("Pas assez de", p.Inventory[answer1-1].Name)
								accessInventory(p)
							}

						} else if p.Inventory[answer1-1].Name == "Potion de poison" {
							if p.Inventory[answer1-1].Quantity > 0 {
								p.Inventory[answer1-1].Quantity -= 1
								PoisonPot(p)
								lastAction = fmt.Sprintf("vous avez consommé 1 %s et vous vous êtes infligé %d points de dégats", p.Inventory[answer1-1].Name, p.Inventory[answer1-1].StatBoost)
								ShowlastAction = true
								fmt.Println("\n★ Appuyez sur entrée pour continuer...")
								fmt.Scanln()
								accessInventory(p)
							} else {
								lasterror = fmt.Sprintf("Pas assez de", p.Inventory[answer1-1].Name)
								accessInventory(p)
							}

						} else {
							lasterror = "Erreur, l'objet sélectionné ne peut pas être utilisé"
						}
					case "N", "n", "no", "nO", "No", "NO", "non", "noN", "nOn", "nON", "Non", "NoN", "NOn", "NON":
						lasterror = "annulé"
						accessInventory(p)
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
							if p.Equipment[0].Name != "" {
								fmt.Println("Un équipement est déjà sur la tête :", p.Equipment[0].Name)
								fmt.Println("★ Voulez-vous le remplacer ? [O/N]")
								fmt.Scanln(&answer3)
								switch answer3 {
								case "o", "O", "ou", "oU", "Ou", "OU", "oui", "ouI", "oUi", "oUI", "Oui", "OuI", "OUi", "OUI":
									AddInventory(p, Items{p.Equipment[0].Name, 1, "Equipment", 0, p.Inventory[answer1-1].StatName, "Equipment"}) // remet l'ancien équipement dans l'inventaire
									lastAction = fmt.Sprintf("vous avez déséquipé 1 %s, il se trouve maintenant dans votre inventaire", p.Equipment[0].Name)
									ShowlastAction = true
									p.Equipment[0].Name = p.Inventory[answer1-1].Name
									RemoveInventory(p, p.Inventory[answer1-1], 1)
									lastAction2 = fmt.Sprintf("vous avez équipé 1 %s", p.Inventory[answer1-1].Name)
									ShowlastAction2 = true
									accessInventory(p)
								case "N", "n", "no", "nO", "No", "NO", "non", "noN", "nOn", "nON", "Non", "NoN", "NOn", "NON":
									lasterror = "annulé"
									accessInventory(p)
								}
							} else if p.Equipment[0].Name == "" {
								p.Equipment[0].Name = p.Inventory[answer1-1].Name
								RemoveInventory(p, p.Inventory[answer1-1], 1)
								lastAction = fmt.Sprintf("vous avez équipé 1 %s", p.Inventory[answer1-1].Name)
								ShowlastAction = true
								p.Equipment[0].Worn = true
								accessInventory(p)
							}

							/*##################################
							  # Si l'equipement est une armure #
							  ##################################*/
						} else if p.Inventory[answer1-1].StatName == "body" {
							if p.Equipment[1].Name != "" {
								fmt.Println("Une armure est déjà équipée :", p.Equipment[1].Name)
								fmt.Println("★ Voulez-vous la remplacer ? [O/N]")
								fmt.Scanln(&answer3)
								switch answer3 {
								case "o", "O", "ou", "oU", "Ou", "OU", "oui", "ouI", "oUi", "oUI", "Oui", "OuI", "OUi", "OUI":
									AddInventory(p, Items{p.Equipment[1].Name, 1, "Equipment", 0, p.Inventory[answer1-1].StatName, "Equipment"}) // remet l'ancien équipement dans l'inventaire
									lastAction = fmt.Sprintf("vous avez déséquipé 1 %s, il se trouve maintenant dans votre inventaire", p.Equipment[1].Name)
									ShowlastAction = true
									p.Equipment[1].Name = p.Inventory[answer1-1].Name
									RemoveInventory(p, p.Inventory[answer1-1], 1)
									lastAction2 = fmt.Sprintf("vous avez équipé 1 %s", p.Inventory[answer1-1].Name)
									ShowlastAction2 = true
									accessInventory(p)
								case "N", "n", "no", "nO", "No", "NO", "non", "noN", "nOn", "nON", "Non", "NoN", "NOn", "NON":
									lasterror = "annulé"
									accessInventory(p)
								}
							} else if p.Equipment[1].Name == "" {
								p.Equipment[1].Name = p.Inventory[answer1-1].Name
								RemoveInventory(p, p.Inventory[answer1-1], 1)
								lastAction = fmt.Sprintf("vous avez équipé 1 %s", p.Inventory[answer1-1].Name)
								ShowlastAction = true
								p.Equipment[1].Worn = true
								accessInventory(p)
							}

							/*###################################
							  # Si l'equipement est un pantalon #
							  ###################################*/
						} else if p.Inventory[answer1-1].StatName == "legs" {
							if p.Equipment[2].Name != "" {
								fmt.Println("Un pantalon est déjà équipé :", p.Equipment[2].Name)
								fmt.Println("★ Voulez-vous le remplacer ? [O/N]")
								fmt.Scanln(&answer3)
								switch answer3 {
								case "o", "O", "ou", "oU", "Ou", "OU", "oui", "ouI", "oUi", "oUI", "Oui", "OuI", "OUi", "OUI":
									AddInventory(p, Items{p.Equipment[2].Name, 1, "Equipment", 0, p.Inventory[answer1-1].StatName, "Equipment"}) // remet l'ancien équipement dans l'inventaire
									lastAction = fmt.Sprintf("vous avez déséquipé 1 %s, il se trouve maintenant dans votre inventaire", p.Equipment[2].Name)
									ShowlastAction = true
									p.Equipment[2].Name = p.Inventory[answer1-1].Name
									RemoveInventory(p, p.Inventory[answer1-1], 1)
									lastAction2 = fmt.Sprintf("vous avez équipé 1 %s", p.Inventory[answer1-1].Name)
									ShowlastAction2 = true
									accessInventory(p)
								case "N", "n", "no", "nO", "No", "NO", "non", "noN", "nOn", "nON", "Non", "NoN", "NOn", "NON":
									lasterror = "annulé"
									accessInventory(p)
								}
							} else if p.Equipment[2].Name == "" {
								p.Equipment[2].Name = p.Inventory[answer1-1].Name
								RemoveInventory(p, p.Inventory[answer1-1], 1)
								lastAction = fmt.Sprintf("vous avez équipé 1 %s", p.Inventory[answer1-1].Name)
								ShowlastAction = true
								p.Equipment[2].Worn = true
								accessInventory(p)
							}
						}
					case "N", "n", "no", "nO", "No", "NO", "non", "noN", "nOn", "nON", "Non", "NoN", "NOn", "NON":
						lasterror = "annulé"
						accessInventory(p)
					}

				} else {
					lasterror = "Erreur, l'objet sélectionné ne peut pas être utilisé"
				}
			case 0:
				break
			default:
				break
			}
		case "flee":
			fmt.Println("Vous avez fuis, mais, le goblin réussi à vous attaquer.")
			fmt.Println("Vous perdez donc 5 points de vie...")
			p.Pv -= 5
			main()
		}
		ct++
		CompteurTour++
		GoblinPattern(Goblin)
		fmt.Println(ActionTour)
		fmt.Println(ResultatTour)
	}
}
