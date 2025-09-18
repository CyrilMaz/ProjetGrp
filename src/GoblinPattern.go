package main

import "fmt"

var ActionTour string
var ResultatTour string
var ct int = 0

func GoblinPattern(Goblin *Monster) {
	ActionTour = ""
	ResultatTour = ""
	if ct == 3 {
		Goblin.Attack *= 2
	}
	ActionTour = fmt.Sprintf("%s inflige %d point de dégats à %s", Goblin.Name, Goblin.Attack, FirstCharacter.Name)
	ResultatTour = fmt.Sprintf("%s a désormais %d/%d HP", FirstCharacter.Name, FirstCharacter.Pv, FirstCharacter.Pvmax)
}
