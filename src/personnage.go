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
	fmt.Println("tu as désormais", FirstCharacter.Inventory[0].Quantity, FirstCharacter.Inventory[0].Name)
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
	initCharacter()
	DisplayInfo(FirstCharacter)
	accessInventory(FirstCharacter)
}
