package main

type Character struct {
	Name                 string
	Class                string
	Level                int
	Pv                   int
	Pvmax                int
	Inventory            []Items
	Gold                 int
	Skills               []Skills
	Equipment            []Equipment
	MaxInventory         int
	InventoryUpgradeUsed bool
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
	damage  int
}

type Equipment struct {
	Slot string
	Worn bool
	Name string
}

type Monster struct {
	Name   string
	Pv     int
	Pvmax  int
	Attack int
}
