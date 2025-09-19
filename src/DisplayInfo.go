package main

import "fmt"

func DisplayInfo(FirstCharacter *Character) {
	for spaces := 0; spaces < 40; spaces++ {
		fmt.Print("\n")
	}
	fmt.Println("┏◇─◇──◇────◇ INFORMATION ◇─────◇──◇─◇┓")
	fmt.Println("======================================")
	fmt.Println("Nom :", FirstCharacter.Name)
	fmt.Println("Classe :", FirstCharacter.Class)
	fmt.Println("Niveau :", FirstCharacter.Level)
	if FirstCharacter.Equipment[0].Worn {
		FirstCharacter.Pv += 10
		FirstCharacter.Pvmax += 10
	}
	if FirstCharacter.Equipment[1].Worn {
		FirstCharacter.Pv += 10
		FirstCharacter.Pvmax += 10
	}
	if FirstCharacter.Equipment[2].Worn {
		FirstCharacter.Pv += 10
		FirstCharacter.Pvmax += 10
	}
	fmt.Println("Points de vie actuels :", FirstCharacter.Pv)
	fmt.Println("Points de vie max :", FirstCharacter.Pvmax)
	fmt.Println("Pieces d'or :", FirstCharacter.Gold)
	fmt.Println("Emplacements d'inventaire max :", FirstCharacter.MaxInventory)
	fmt.Println("Compétences apprises :")
	for i := 0; i < len(FirstCharacter.Skills); i++ {
		if FirstCharacter.Skills[i].Learned {
			fmt.Print("☆", FirstCharacter.Skills[i].Name)
		}
		if i == len(FirstCharacter.Skills)-1 {
			fmt.Print(".")
		} else {
			fmt.Print(", ")
		}
	}
	fmt.Println("\nÉquipements portés :")
	for i := 0; i < len(FirstCharacter.Equipment); i++ {
		if FirstCharacter.Equipment[i].Worn {
			fmt.Print(FirstCharacter.Equipment[i].Name)
			if i+1 > len(FirstCharacter.Equipment) {
				fmt.Print(".")
			} else if i < len(FirstCharacter.Equipment) {
				fmt.Print(", ")
			}
		}
	}
	fmt.Print("\n")
	fmt.Println("\n★━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━★")
	fmt.Println("|★ Appuyez sur une Entrée pour sortir du menu |")
	fmt.Println("★━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━★")
	fmt.Scanln(&answer)
	switch answer {
	default:
		main()
	}
}
