package main

import (
	"fmt"
)

var FirstCharacter Character

func initCharacter() {
	var nom string
	for {
		fmt.Println("Quel est le nom de ton personnage")
		fmt.Scanln(&nom) // scan le nom du personnage
		match := true
		for i := 0; i < len(nom); i++ {
			if (nom[i] < 97 || nom[i] > 122) && (nom[i] < 65 || nom[i] > 90) && len(nom) > 0 {
				match = false
			}
		}
		if !match {
			fmt.Println("Ton nom ne peut contenir que des lettres ! : A_Z, a_z")
			fmt.Println("réessaie !")
		} else {
			break
		}
	}
	if nom[0] >= 'a' && nom[0] <= 'z' {
		nom = string(nom[0]-32) + nom[1:]
	}
	for i := 1; i < len(nom); i++ {
		if nom[i] >= 'A' && nom[i] <= 'Z' {
			nom = nom[:i] + string(nom[i]+32) + nom[i+1:]
		}
	}
	FirstCharacter = Character{nom, "Elfe", 1, 40, 100, []Items{Items{"potions de soin", 3, "soin", 20, "all", "Consumable"}}, 100, []Skills{{"coup de poing", true}}}
}
