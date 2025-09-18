package main

import (
	"fmt"
	"time"
)

func PoisonPot(p *Character) {
	fmt.Println("Vous avez bu une potion empoisonnée !")
	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Vous perdez 10 points de vie.")
		p.Pv -= 10
		if p.Pv <= 0 {
			fmt.Println("Vous êtes mort !")
			IsDead(p)
		} else {
			fmt.Println("Il vous reste", p.Pv, "/", p.Pvmax, "points de vie.")
		}
	}
	RemoveInventory(p, PotionPoison, 1)
}
