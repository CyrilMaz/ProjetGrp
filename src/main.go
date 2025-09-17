package main

import "fmt"

var initalized bool = false
var answer string

func main() {
	for spaces := 0; spaces < 40; spaces++ {
		fmt.Print("\n")
	}
	if !initalized {
		initCharacter()
		initalized = true
		fmt.Println("╔══════════════════════════════════════════════╗")
		fmt.Println("║                                              ║")
		fmt.Println("║      BIENVENUE DANS LE MONDE D'AVENTURE      ║")
		fmt.Println("║                                              ║")
		fmt.Println("║        ──◇◆◇──  Projet Groupe  ──◇◆◇──       ║")
		fmt.Println("║                                              ║")
		fmt.Println("║    Prépare-toi à forger ton destin, héros !  ║")
		fmt.Println("║                                              ║")
		fmt.Println("╚══════════════════════════════════════════════╝")
	}
	for answer != "0" {
		fmt.Println("★━━━━━━━━━━━━━━━━━━━━━━━━━★")
		fmt.Println("| Que voulez-vous faire ? |")
		fmt.Println("★━━━━━━━━━━━━━━━━━━━━━━━━━★")
		for i := 0; i < len(MENU01); i++ {
			fmt.Println(MENU01[i])
		}
		fmt.Print("\n")
		fmt.Println(MENU00)
		fmt.Print("\n")
		fmt.Println("★━━━━━━━━━━━━━━━━━━━━━━★")
		fmt.Println("| Quel est ton choix ? |")
		fmt.Println("★━━━━━━━━━━━━━━━━━━━━━━★")
		fmt.Scanln(&answer)

		switch answer {
		case "1":
			DisplayInfo(&FirstCharacter)
		case "2":
			accessInventory(&FirstCharacter)
		case "3":
			Merchand(&FirstCharacter)
		case "0":
			fmt.Println("Au revoir", FirstCharacter.Name, "!")
		default:
			fmt.Println("Choix invalide, réessayez.")
		}
	}
}
