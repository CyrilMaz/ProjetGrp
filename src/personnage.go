package main

import (
	"fmt"
)

type Character struct {
	Name       string
	Class      string
	Level      int
	Pv         int
	Pvmax      int
	Inventory []Items
}

type Items struct {
	Name            string
	Quantity        int
	StatName        string
	StatBoost       int
	RespectiveClass string
}

// Variable globale accessible partout
var FirstCharacter Character
var MENU01 = []string{"0 : sortir", "1 : infos du personnage", "2 : accéder à l'inventaire"}



<<<<<<< HEAD
func main() {
	var nom string
	fmt.Println("Quel est le nom de ton personnage")
	fmt.Scanln(&nom) // scan le nom du personnage
	match := true
	for i := 0; i < len(nom); i++ {
		if (nom[i] < 97 || nom[i] > 122) && (nom[i] < 65 || nom[i] > 90) {
			match = false
		}
	}
	if match == false { 
		fmt.Println("Ton nom ne peut contenir que des lettres ! : A_Z, a_z")
		fmt.Println("réessaie !")
		fmt.Scanln(&nom)
	}
	fmt.Println("Ton nom est : " + nom)
	PotionSoin := Items{"potion de soin", 3, "soin", 20, "all"}
	FirstCharacter := Character{nom, "Elfe", 1, 40, 100, []Items{PotionSoin}}
	fmt.Println(FirstCharacter.Inventaire[0].Name, FirstCharacter.Inventaire[0].Quantity)
=======
func initCharacter() {
	var nom string
	for {
		fmt.Println("Quel est le nom de ton personnage")
		fmt.Scanln(&nom) // scan le nom du personnage
		match := true
		for i := 0; i < len(nom); i++ {
			if (nom[i] < 97 || nom[i] > 122) && (nom[i] < 65 || nom[i] > 90) {
				match = false
			}
		}
		if match == false {
			fmt.Println("Ton nom ne peut contenir que des lettres ! : A_Z, a_z")
			fmt.Println("réessaie !")
		} else {
			break
		}
	}
	PotionSoin := Items{"potions de soin", 3, "soin", 20, "all"}
	FirstCharacter = Character{nom, "Elfe", 1, 40, 100, []Items{PotionSoin}}
}

func DisplayInfo(p Character) {
	fmt.Println("Nom :", p.Name)
	fmt.Println("Classe :", p.Class)
	fmt.Println("Niveau :", p.Level)
	fmt.Println("Points de vie actuels :", p.Pv)
	fmt.Println("Points de vie max :", p.Pvmax)
	for i := 0; i < len(p.Inventory); i++ {
		fmt.Println("tu as", p.Inventory[i].Quantity, p.Inventory[i].Name)
	}
}

func accessInventory(FirstCharacter Character) {
	for i := 0; i < len(FirstCharacter.Inventory); i++ {
		fmt.Println(FirstCharacter.Inventory[i].Name, FirstCharacter.Inventory[i].Quantity)
	}
}


func main() {
	var answer string 
	initCharacter()
	fmt.Println("Que voulez vous faire ?")
	for i := 0; i < len(MENU01); i++ {
		fmt.Println(MENU01[i])
	}
	fmt.Scanln(&answer)
	switch answer {
	case "0":
		return
	case "1":
		DisplayInfo(FirstCharacter)
	case "2":
		accessInventory(FirstCharacter)
	}
>>>>>>> 0329b560724255582950509edaa6ae1dfbbbf246
}