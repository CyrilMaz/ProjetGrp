package main

var PotionSoin = Items{"potions de soin", 1, "soin", 20, "all", "Consumable"}
var LivreSorts = Items{"Livre de sorts", 1, "FireBall", 0, "all", "Consumable"}
var PotionPoison = Items{"Potion de poison", 1, "poison", 0, "all", "Consumable"}
var FourrureLoup = Items{"Fourrure de loup", 1, "craft", 0, "all", "Material"}
var PeauTroll = Items{"Peau de troll", 1, "craft", 0, "all", "Material"}
var CuirSanglier = Items{"Cuir de sanglier", 1, "craft", 0, "all", "Material"}
var PlumeCorbeau = Items{"Plume de corbeau", 1, "craft", 0, "all", "Material"}
var ChapeauAventurier = Items{
	Name:            "Chapeau d'aventurier",
	Quantity:        1,
	StatName:        "Head",
	StatBoost:       10, // ou la valeur que tu veux
	RespectiveClass: "all",
	Type:            "Equipment",
}
var TuniqueAventurier = Items{
	Name:            "Tunique d'aventurier",
	Quantity:        1,
	StatName:        "body",
	StatBoost:       25,
	RespectiveClass: "all",
	Type:            "Equipment",
}
var BottesAventurier = Items{
	Name:            "Bottes d'aventurier",
	Quantity:        1,
	StatName:        "legs",
	StatBoost:       15,
	RespectiveClass: "all",
	Type:            "Equipment",
}
