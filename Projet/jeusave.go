package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode"
	"encoding/json"
	//"flag"
)

func main() {
	fmt.Println(`
	 _   _    _    _   _  ____ __  __    _    _   _ 
	| | | |  / \  | \ | |/ ___|  \/  |  / \  | \ | |
	| |_| | / _ \ |  \| | |  _| |\/| | / _ \ |  \| |
	|  _  |/ ___ \| |\  | |_| | |  | |/ ___ \| |\  |
	|_| |_/_/   \_\_| \_|\____|_|  |_/_/   \_\_| \_|
	`)
	
	//Ouvre le fichier words.txt qui contient les mots à deviner
	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println("Erreur ouverture fichier:", err) //Message d'erreur si le fichier n'est pas trouvé
		return
	}
	defer file.Close() //Ferme le fichier à la fin de la fonction

	scanner := bufio.NewScanner(file) //Scanner pour lire le fichier ligne par ligne

	var words []string //Initialisation d'un tableau de string pour stocker les mots du fichier

	for scanner.Scan() {
		words = append(words, scanner.Text()) //Ajoute chaque ligne du fichier dans le tableau
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lecture fichier:", err) //Message d'erreur si le fichier ne peut pas être lu
		return
	}

	if len(words) == 0 {
		fmt.Println("Aucun mots trouvés dans le fichier.") //Message d'erreur si le fichier est vide
		return
	}

	randomIndex := rand.Intn(len(words)) //Choisi un mot au hasard dans le tableau
	selectedWord := words[randomIndex]

	n := len(selectedWord)/2 - 1 //Nombre de lettres à révéler

	hiddenWordSlice := make([]rune, len(selectedWord)) //Initialisation d'un tableau de runes pour stocker le mot caché
	for i := range hiddenWordSlice {
		hiddenWordSlice[i] = '_' //Rempli le tableau de '_' pour cacher le mot
	}

	revealIndices := rand.Perm(len(selectedWord))[:n] //Choisi n indices au hasard dans le tableau

	for _, index := range revealIndices {
		hiddenWordSlice[index] = []rune(selectedWord)[index] //Remplace les '_' par les lettres du mot à révéler (indices)
	}

	attempts := 0 //Nombre de tentatives
	maxAttempts := 10 //Nombre de tentatives maximum
	lettersTried := "" //Lettres déjà essayées
	lettersFound := "" //Lettres trouvées

	for {
		if attempts == maxAttempts { //Si le nombre de tentatives est égal au nombre de tentatives maximum, le joueur a perdu
			fmt.Println("José est mort †. Le mot était:", selectedWord)
			break
		}
	
		var input string
		fmt.Println("Mot à deviner :", string(hiddenWordSlice)) //Affiche le mot caché
		fmt.Println("Lettres déjà essayées :", lettersTried) //Affiche les lettres déjà essayées
		fmt.Print("Lettre ou mot : ") //Demande à l'utilisateur d'entrer une lettre ou un mot
	
		_, err := fmt.Scanln(&input) //Lis l'input de l'utilisateur
		if err != nil {
			fmt.Println("Entrée invalide. Veuillez entrer une lettre ou un mot. \n") //Message d'erreur si l'input n'est pas valide
			continue
		}
	
		input = strings.ToLower(input) //Met l'input en minuscule pour la compatibilité
	
		if input == "stop" { //Si l'input est "stop"
			SaveGame(attempts, lettersTried, lettersFound, string(hiddenWordSlice)) //Sauvegarde la partie
			break //Arrête le jeu
		}

		if len(input) == 1 && isLetter(input) { //Boucle si l'input est une lettre
			if strings.Contains(lettersTried, input) { //Si la lettre a déjà été essayée
				fmt.Println("Cette lettre a déjà été essayée. Essayez-en une autre. \n")
				continue
			}
	
			letterFound := false //Initialisation d'une variable pour savoir si la lettre est dans le mot
			for i, char := range selectedWord { //Boucle pour vérifier si la lettre est dans le mot
				if input == string(char) && hiddenWordSlice[i] == '_' { //Si la lettre est dans le mot et n'a pas déjà été trouvée
					hiddenWordSlice[i] = char //Remplace le '_' par la lettre
					letterFound = true
					lettersFound += input + "-" //Ajoute la lettre à la liste des lettres trouvées

				}
			}
	
			lettersTried += input + "-" //Ajoute la lettre essayée à la liste des lettres essayées
	
			if letterFound { //Si la lettre est dans le mot
				pendu(attempts) //Affiche le pendu
				fmt.Println("Vous avez trouvé une lettre ! \n")
			} else { //Si la lettre n'est pas dans le mot
				attempts++ //Ajoute 1 au nombre de tentatives
				pendu(attempts) //Affiche le pendu
				fmt.Println("Raté, cette lettre n'est pas dans le mot. Nombre de tentatives restantes:", maxAttempts-attempts, "\n")
			}
		} else if len(input) > 1 { //Si l'input est un mot
			if input == selectedWord { //Si l'inpu est le mot à trouver
				fmt.Println("Bravo, vous avez sauvé José, le mot était:", selectedWord)
				break
			} else { //Si l'input n'est pas le mot à trouver
				attempts += 2 //Ajoute 2 au nombre de tentatives
				pendu(attempts) //Affiche le pendu
				fmt.Println("Raté, ce n'est pas le mot. Nombre de tentatives restantes:", maxAttempts-attempts, "\n")
			}
		} else { //Si l'input n'est pas valide
			fmt.Println("Entrée invalide. Veuillez entrer une lettre ou un mot. \n")
		}
	}
}

func isLetter(s string) bool { //Fonction pour vérifier si l'input est une lettre
	return len(s) == 1 && unicode.IsLetter([]rune(s)[0])
}

func pendu(stage_of_death int) { //Fonction pour afficher le pendu
	stages := []string{
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
	fmt.Println(stages[stage_of_death])
}

func SaveGame(essais int, lettersTried string, letterFound string, pendustate string) {
    saveData := map[string]interface{}{
        "essais": essais,
        "lettersTried": lettersTried,
        "letterFound": letterFound,
        "pendustate": pendustate,
    }

    jsonData, err := json.Marshal(saveData)
    if err != nil {
        fmt.Println("Erreur encodage JSON:", err)
        return
    }

    file, err := os.Create("save.txt")
    if err != nil {
        fmt.Println("Erreur création fichier:", err)
        return
    }
    defer file.Close()

    _, err = file.Write(jsonData)
    if err != nil {
        fmt.Println("Erreur écriture dans le fichier:", err)
        return
    }

    fmt.Println("Partie sauvegardée dans le fichier save.txt")
}