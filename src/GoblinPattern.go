package main

import "fmt"

var ActionTour string
var ResultatTour string
var ct int = 0

func GoblinPattern(Goblin *Monster, p *Character) {
	ActionTour = ""
	ResultatTour = ""
	if ct == 3 {
		ActionTour = fmt.Sprintf("%s inflige %d point de dégats à %s", Goblin.Name, Goblin.Attack, p.Name)
		p.Pv -= Goblin.Attack * 2
		ResultatTour = fmt.Sprintf("%s a désormais %d/%d HP", p.Name, p.Pv, p.Pvmax)
		ct = 0
	} else {
		ActionTour = fmt.Sprintf("%s inflige %d point de dégats à %s", Goblin.Name, Goblin.Attack, p.Name)
		p.Pv -= Goblin.Attack
		ResultatTour = fmt.Sprintf("%s a désormais %d/%d HP", p.Name, p.Pv, p.Pvmax)
	}
}
