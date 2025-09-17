package main

type Character struct {
	Name      string
	Class     string
	Level     int
	Pv        int
	Pvmax     int
	Inventory []Items
	Gold      int
	Skills    []Skills
	Equip     Equipment
}

type Items struct {
	Name            string
	Quantity        int
	StatName        string
	StatBoost       int
	RespectiveClass string
	Type            string
}

type Skills struct {
	Name    string
	Learned bool
}

type Equipment struct {
	Head string
	Body string
	Legs string
}
