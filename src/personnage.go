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
	Inventaire []Items
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
	FirstCharacter = Character{nom, "Elfe", 1, 40, 100, []Items{PotionSoin}} // = au lieu de :=
	fmt.Println("tu as désormais", FirstCharacter.Inventaire[0].Quantity, FirstCharacter.Inventaire[0].Name)
}

func DisplayInfo(p Character) {
	fmt.Println("Nom :", p.Name)
	fmt.Println("Classe :", p.Class)
	fmt.Println("Niveau :", p.Level)
	fmt.Println("Points de vie actuels :", p.Pv)
	fmt.Println("Points de vie max :", p.Pvmax)
	for i := 0; i < len(p.Inventaire); i++ {
		fmt.Println("tu as", p.Inventaire[i].Quantity, p.Inventaire[i].Name)
	}
}

func main() {
	initCharacter()
	DisplayInfo(FirstCharacter) // Maintenant FirstCharacter est accessible
}