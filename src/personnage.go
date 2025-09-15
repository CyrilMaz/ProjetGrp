package main

import "fmt"

type Character struct {
	Name       string
	Class      string
	level      int
	pv         int
	pvmax      int
	inventaire []string
}

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
}