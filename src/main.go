package main

import (
	"fmt"
	"os"
)

var initalized bool = false
var answer string

func main() {
	for spaces := 0; spaces < 40; spaces++ {
		fmt.Print("\n")
	}
	if !initalized {
		initCharacter()
		initalized = true
		for spaces := 0; spaces < 40; spaces++ {
			fmt.Print("\n")
		}
		fmt.Println("╔══════════════════════════════════════════════╗")
		fmt.Println("║                                              ║")
		fmt.Println("║      BIENVENUE DANS LE MONDE D'AVENTURE      ║")
		fmt.Println("║                                              ║")
		fmt.Println("║        ──◇◆◇──  FORGE  DESTIN  ──◇◆◇──       ║")
		fmt.Println("║                                              ║")
		fmt.Println("║    Prépare-toi à forger ton destin, héros !  ║")
		fmt.Println("║                                              ║")
		fmt.Println("╚══════════════════════════════════════════════╝")
	}
	fmt.Println("")
	fmt.Println("★¸„.-•~¹°”ˆ˜¨ ACTIONS ¨˜ˆ”°¹~•-.„¸★")
	fmt.Println("|                                 |")
	fmt.Println("★━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━★")
	for i := 0; i < len(MENU01); i++ {
		fmt.Println(MENU01[i])
	}
	fmt.Println("|                                 |")
	fmt.Println("★━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━★")
	fmt.Println("|        0 : Sortir du jeu        |")
	fmt.Println("★━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━★")
	fmt.Print("\n")
	if lasterror != "" {
		fmt.Println("Erreur :", lasterror)
		lasterror = ""
	}

	fmt.Println("★ Quel est ton choix ? ")

	fmt.Scanln(&answer)

	switch answer {
	case "1":
		DisplayInfo(&FirstCharacter)
	case "2":
		accessInventory(&FirstCharacter)
	case "3":
		Merchand(&FirstCharacter)
	case "4":
		BlackSmith(&FirstCharacter)
	case "5":
		trainingFight(&FirstCharacter, &Goblin)
	case "/kill":
		FirstCharacter.Pv = 0
		IsDead(&FirstCharacter)
	case "/gold":
		FirstCharacter.Gold += 1000
		main()
	case "0":
		fmt.Println("Au revoir", FirstCharacter.Name, "!")
		os.Exit(0)
	default:
		lasterror = "Choix invalide, réessaie."
		main()
	}
}
