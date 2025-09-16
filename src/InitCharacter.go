package main

import "fmt"

var pv int = 40
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
	FirstCharacter = Character{nom, "Elfe", 1, pv, 100, []Items{PotionSoin}}
}
