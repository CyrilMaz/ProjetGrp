package main

import "fmt"

var Showlastcraft bool = false
var lastcraft string = ""
var lasterror string = ""

func BlackSmith(p *Character) {
	for spaces := 0; spaces < 40; spaces++ {
		fmt.Print("\n")
	}
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣰⣿⣿⣷⡄⠀⠀⠀⠀⠀⠀⠀⢠⣄⣤⣦⣤⣀⠀⠀")
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢿⣿⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⠛⠿⠟⠀⠀")
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⣠⠀⠘⢿⣿⠟⠀⢠⡀⠀⠀⠀⠀⠀⠀⠀⣰⡗⠀⠀⠀⠀")
	fmt.Println("⠀⠀⠀⠀⠀⢠⣾⠀⣿⠀⣷⣦⣤⣴⡇⢸⡇⠀⣷⠀⠀⠀⠀⣰⡟⠀⠀⠀⠀⠀")
	fmt.Println("⠀⠀⠀⠀⠀⣿⣿⠀⣿⣤⣈⣉⣉⣉⣠⣼⡇⠀⣿⡆⠀⠀⣰⡟⠀⠀⠀⠀⠀⠀")
	fmt.Println("⠀⠀⠀⠀⠀⣿⣿⠀⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⣿⠇⠀⠀⠛⠀⠀⠀⠀⠀⠀⠀")
	fmt.Println("⠀⠀⠀⠀⠀⠛⠛⠀⠛⠛⠛⠛⠛⠛⠛⠛⠃⠀⠛⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
	fmt.Println("⠀⠀⠀⠀⣤⣤⣤⣤⣤⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣇⠀⠀⠀⠀⠀⠀⠀")
	fmt.Println("⠀⠀⠀⠀⠈⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣤⠀⠀⠀⠀")
	fmt.Println("⠀⠀⠀⠀⠀⠀⠈⠙⠛⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠟⠋⠀⠀⠀⠀⠀")
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠿⠿⠿⣿⣿⣿⣿⣿⣿⣿⠿⠿⠿⠇⠀⠀⠀⠀⠀⠀⠀")
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣰⣿⣿⣿⣿⣿⣿⣿⣧⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣦⡀⠀⠀⠀⠀⠀⠀⠀")
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠘⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠃⠀⠀⠀⠀⠀")

	fmt.Println(" ┈ ┈ ┈ ┈ ┈ ⋞ <  ⏣  > ⋟ ┈ ┈ ┈ ┈ ┈")
	fmt.Println("\n   Bienvenue chez le forgeron !")

	fmt.Println("\n,„.-•~¹°”ˆ˜¨  FABRICATION  ¨˜ˆ”°¹~•-.„,")

	fmt.Println("| 1 : Chapeau de l’aventurier ·(5 PO) |")
	fmt.Println("| 2 : Tunique de l’aventurier ·(5 PO) |")
	fmt.Println("| 3 : Bottes de l’aventurier ··(5 PO) |")
	fmt.Println("|                                     |")
	fmt.Println("★━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━★")
	fmt.Println("|        0 : Sortir de la forge       |")
	fmt.Println("|      /inv : Aller à l'inventaire    |")
	fmt.Println("★━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━★")
	if lasterror != "" {
		fmt.Println("Erreur :", lasterror)
		lasterror = ""
	} else if Showlastcraft {
		fmt.Println("Dernière fabriation :", lastcraft)
		Showlastcraft = false
	}
	var choice string
	DisplayAddItem = false
	fmt.Scanln(&choice) // Lire le choix de l'utilisateur

	switch choice {
	case "1":
		Showlastcraft = true
		if p.Gold >= 5 && HasItem(p, PlumeCorbeau, 1) && HasItem(p, CuirSanglier, 1) {
			p.Gold -= 5
			RemoveInventory(p, PlumeCorbeau, 1)
			RemoveInventory(p, CuirSanglier, 1)
			AddInventory(p, ChapeauAventurier)
			lastcraft = fmt.Sprintf("Vous avez ajouté %d %s à votre inventaire.", ChapeauAventurier.Quantity, ChapeauAventurier.Name)
		} else if p.Gold < 5 {
			lastcraft = "Vous n'avez pas assez d'or pour fabriquer cet objet."
		} else if !HasItem(p, PlumeCorbeau, 1) {
			lastcraft = "Il vous manque une Plume de Corbeau pour fabriquer cet objet."
		} else if !HasItem(p, CuirSanglier, 1) {
			lastcraft = "Il vous manque 1 de Cuir de Sanglier pour fabriquer cet objet."
		}
		BlackSmith(p)
	case "2":
		Showlastcraft = true
		if p.Gold >= 5 && HasItem(p, FourrureLoup, 2) && HasItem(p, PeauTroll, 1) {
			p.Gold -= 5
			RemoveInventory(p, FourrureLoup, 2)
			RemoveInventory(p, PeauTroll, 1)
			AddInventory(p, TuniqueAventurier)
			lastcraft = fmt.Sprintf("Vous avez ajouté %d %s à votre inventaire.", TuniqueAventurier.Quantity, TuniqueAventurier.Name)
		} else if p.Gold < 5 {
			lasterror = "Vous n'avez pas assez d'or pour fabriquer cet objet."
		} else if !HasItem(p, FourrureLoup, 2) {
			lasterror = "Il vous manque de la Fourrures de Loup pour fabriquer cet objet."
		} else if !HasItem(p, PeauTroll, 1) {
			lasterror = "Il vous manque une Peau de Troll pour fabriquer cet objet."
		}
		BlackSmith(p)
	case "3":
		Showlastcraft = true
		if p.Gold >= 5 && HasItem(p, FourrureLoup, 1) && HasItem(p, CuirSanglier, 1) {
			p.Gold -= 5
			RemoveInventory(p, FourrureLoup, 1)
			RemoveInventory(p, CuirSanglier, 1)
			AddInventory(p, BottesAventurier)
			lastcraft = fmt.Sprintf("Vous avez ajouté %d %s à votre inventaire.", BottesAventurier.Quantity, BottesAventurier.Name)
		} else if p.Gold < 5 {
			lasterror = "Vous n'avez pas assez d'or pour fabriquer cet objet."
		} else if !HasItem(p, FourrureLoup, 1) {
			lastcraft = "Il vous manque 1 de Fourrure de Loup pour fabriquer cet objet."
		} else if !HasItem(p, CuirSanglier, 1) {
			lastcraft = "Il vous manque 1 de Cuir de Sanglier pour fabriquer cet objet."
		}
		BlackSmith(p)
	case "0":
		main()
	case "/inv":
		accessInventory(p)
	case "give all":
		AddInventory(p, ChapeauAventurier)
		AddInventory(p, TuniqueAventurier)
		AddInventory(p, BottesAventurier)
		AddInventory(p, FourrureLoup)
		AddInventory(p, PeauTroll)
		AddInventory(p, CuirSanglier)
		AddInventory(p, PlumeCorbeau)
	default:
		BlackSmith(p)
	}
}
