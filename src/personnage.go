package ProjetGrp

import "fmt"

type Character struct {
	Name       string
	Class      string
	level      int
	pv         int
	pvmax      int
	inventaire []string
}

func initCharacter() {
	nom := Character.Name
	fmt.Println("Quel est le nom de ton personnage")
	fmt.Scanln(&nom)
}