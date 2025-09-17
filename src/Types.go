package main

type Character struct {
	Name      string
	Class     string
	Level     int
	Pv        int
	Pvmax     int
	Inventory []Items
	Argent    int
}

type Items struct {
	Name            string
	Quantity        int
	StatName        string
	StatBoost       int
	RespectiveClass string
	Type            string
}
