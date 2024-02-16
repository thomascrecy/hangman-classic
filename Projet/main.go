package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(`
	 _   _    _    _   _  ____ __  __    _    _   _ 
	| | | |  / \  | \ | |/ ___|  \/  |  / \  | \ | |
	| |_| | / _ \ |  \| | |  _| |\/| | / _ \ |  \| |
	|  _  |/ ___ \| |\  | |_| | |  | |/ ___ \| |\  |
	|_| |_/_/   \_\_| \_|\____|_|  |_/_/   \_\_| \_|
	`)

	fmt.Println("Bienvenue dans le jeu du pendu !")
	fmt.Println("Le but du jeu est de deviner le mot caché.")
	fmt.Println("Vous pouvez entrer une lettre ou un mot.")
	fmt.Println("Vous avez 10 tentatives pour trouver le mot.")
	fmt.Println("Si vous entrez une lettre qui n'est pas dans le mot, vous perdez 1 tentative.")
	fmt.Println("Si vous entrez un mot qui n'est pas le mot à trouver, vous perdez 2 tentatives.")
	fmt.Println("Bonne chance ! \n")

	for { //Boucle pour rejouer
		playGame() 

		var playAgain string //Variable pour rejouer
		fmt.Print("Voulez-vous jouer à nouveau ? (y/n): ")
		_, err := fmt.Scanln(&playAgain) //Scan l'input
		if err != nil { //Si l'input n'est pas valide
			fmt.Println("Entrée invalide. Veuillez entrer 'y' pour rejouer ou 'n' pour quitter.")
			break
		}

		if strings.ToLower(playAgain) != "y" { //Si l'input est différent de 'y'
			fmt.Println("Merci d'avoir joué ! Au revoir.")
			break
		}
	}
}

func playGame() { //Fonction pour jouer
	file, err := os.Open("words.txt") //Ouvre le fichier
	if err != nil {
		fmt.Println("Erreur ouverture fichier:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) //Scanner pour lire le fichier ligne par ligne
	var words []string //Initialisation d'un tableau de string pour stocker les mots du fichier

	for scanner.Scan() { //Ajoute chaque ligne du fichier dans le tableau
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil { //Message d'erreur si le fichier ne peut pas être lu
		fmt.Println("Erreur lecture fichier:", err)
		return
	}

	if len(words) == 0 { //Message d'erreur si le fichier est vide
		fmt.Println("Aucun mots trouvés dans le fichier.")
		return
	}

	randomIndex := rand.Intn(len(words)) //Choisi un mot au hasard dans le tableau
	selectedWord := words[randomIndex] //Mot à deviner

	n := len(selectedWord)/2 - 1 //Nombre de lettres à révéler

	hiddenWordSlice := make([]rune, len(selectedWord)) //Initialisation d'un tableau de runes pour stocker le mot caché
	for i := range hiddenWordSlice {
		hiddenWordSlice[i] = '_' //Rempli le tableau de '_' pour cacher le mot
	}

	revealIndices := rand.Perm(len(selectedWord))[:n] //Choisi n indices au hasard dans le tableau

	for _, index := range revealIndices { //Remplace les '_' par les lettres du mot à révéler (indices)
		hiddenWordSlice[index] = []rune(selectedWord)[index]
	}

	attempts := 0
	maxAttempts := 10
	lettersTried := ""

	for {
		if attempts == maxAttempts { //Si le nombre de tentatives est égal au nombre de tentatives maximum, le joueur a perdu
			fmt.Println("José est mort †. Le mot était:", selectedWord)
			break
		}

		var input string  //Initialisation de la variable pour l'input
		fmt.Println("Mot à deviner :", string(hiddenWordSlice))
		fmt.Println("Lettres déjà essayées :", lettersTried)
		fmt.Print("Lettre ou mot : ")

		_, err := fmt.Scanln(&input) //Scan l'input
		if err != nil {
			fmt.Println("Entrée invalide. Veuillez entrer une lettre ou un mot. \n")
			continue
		}

		input = strings.ToLower(input) //Met l'input en minuscule

		if len(input) == 1 && isLetter(input) { //Si l'input est une lettre
			if strings.Contains(lettersTried, input) { //Si la lettre a déjà été essayée
				fmt.Println("Cette lettre a déjà été essayée. Essayez-en une autre. \n")
				continue
			}

			letterFound := false 
			for i, char := range selectedWord { //Boucle pour vérifier si la lettre est dans le mot
				if input == string(char) && hiddenWordSlice[i] == '_' { //Si la lettre est dans le mot et n'a pas déjà été trouvée
					hiddenWordSlice[i] = char //Remplace le '_' par la lettre
					letterFound = true //La lettre a été trouvée
				}
			}

			lettersTried += input + "-" //Ajoute la lettre essayée à la liste des lettres essayées

			if letterFound { //Si la lettre a été trouvée
				pendu(attempts) //Affiche le pendu
				fmt.Println("Vous avez trouvé une lettre ! \n")
			} else { //Si la lettre n'a pas été trouvée
				attempts++  //Ajoute 1 à la variable des tentatives
				pendu(attempts) //Affiche le pendu
				fmt.Println("Raté, cette lettre n'est pas dans le mot. Nombre de tentatives restantes:", maxAttempts-attempts, "\n")
			}
		} else if len(input) > 1 { //Si l'input est un mot
			if input == selectedWord { //Si le mot est le mot à trouver
				fmt.Println("Bravo, vous avez sauvé José, le mot était:", selectedWord)
				break
			} else { //Si le mot n'est pas le mot à trouver
				attempts += 2 //Ajoute 2 à la variable des tentatives
				pendu(attempts) //Affiche le pendu
				fmt.Println("Raté, ce n'est pas le mot. Nombre de tentatives restantes:", maxAttempts-attempts, "\n")
			}
		} else { //Si l'input n'est pas valide
			fmt.Println("Entrée invalide. Veuillez entrer une lettre ou un mot. \n")
		}
		
		if !strings.Contains(string(hiddenWordSlice), "_") { //Si le mot est trouvé
			fmt.Println("Bravo, vous avez sauvé José, le mot était:", selectedWord)
			break
		}
	}

	fmt.Println("Partie terminée.")
}

func isLetter(s string) bool { //Fonction pour vérifier si l'input est une lettre
	return len(s) == 1 && unicode.IsLetter([]rune(s)[0]) 
}

func pendu(stage_of_death int) { //Fonction pour afficher le pendu
	stages := []string{ //Tableau des étapes du pendu
		`
========`,
		`
========`,
		`
    	|
    	|
    	|
    	|
    	|
========`,
		`
    +---+
        |
        |
        |
        |
        |
========`,
		`
    +---+
    |   |
        |
        |
        |
        |
========`,
		`
    +---+
    |   |
    O   |
        |
        |
        |
========`,
		`
    +---+
    |   |
    O   |
    |   |
        |
        |
========`,
		`
    +---+
    |   |
    O   |
   /|   |
        |
        |
========`,
		`
    +---+
    |   |
    O   |
   /|\  |
        |
        |
========`,
		`
    +---+
    |   |
    O   |
   /|\  |
   /    |
        |
========`,
		`
    +---+
    |   |
    O   |
   /|\  |
   / \  |
        |
========`,
	}
	fmt.Println(stages[stage_of_death]) //Affiche l'étape du pendu correspondant au nombre de tentatives
}
