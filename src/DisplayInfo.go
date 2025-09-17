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
	fmt.Println("Points de vie actuels :", FirstCharacter.Pv)
	fmt.Println("Points de vie max :", FirstCharacter.Pvmax)
	fmt.Println("Pieces d'or :", FirstCharacter.Gold)
	fmt.Println("Compétences apprises :")
	for i := 0; i < len(FirstCharacter.Skills); i++ {
		if FirstCharacter.Skills[i].Learned {
			fmt.Print(FirstCharacter.Skills[i].Name)
		}
		if i == len(FirstCharacter.Skills)-1 {
			fmt.Print(".")
		} else {
			fmt.Print(", ")
		}
	}
	fmt.Print("\n")
	fmt.Println("\n★━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━★")
	fmt.Println("|★ Appuyez sur une touche pour sortir du menu |")
	fmt.Println("★━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━★")
	fmt.Scanln(&answer)
	switch answer {
	default:
		main()
	}
}
