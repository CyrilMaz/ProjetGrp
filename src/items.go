package Projetgrp

import "fmt"

type Items struct {
	name string
	quantity int
	statName string
	statBoost int
	level int
	EXP int
	NextLvlExp int
	RespectiveClass string
}

func ShowItems() {
	fmt.Println()
}