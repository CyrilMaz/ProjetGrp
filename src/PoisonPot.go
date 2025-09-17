package main

import (
	"fmt"
	"time"
)

func PoisonPot(p *Character) {
	fmt.Println("Vous avez bu une potion empoisonnée !")
	time.Sleep(3 * time.Second)
	fmt.Println("Vous perdez 10 points de vie.")
	p.Pv -= 10
	if p.Pv <= 0 {
		fmt.Println("Vous êtes mort !")
		//IsDead(p) // Appelle la fonction IsDead pour gérer la mort du personnage
	} else {
		fmt.Println("Il vous reste", p.Pv, "/", p.Pvmax, "points de vie.")
	}
}
