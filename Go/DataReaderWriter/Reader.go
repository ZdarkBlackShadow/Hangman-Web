package datareaderwriter

import (
	"bufio"
	"fmt"
	game "game/Game"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Reader(filename string) {
	// Créer le chemin complet vers le fichier
	path := filepath.Join("data", "français", filename)
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Erreur lors de l'ouverture du fichier '%s': %v\n", path, err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier '%s': %v\n", path, err)
	}
}

// Reader lit le contenu d'un fichier dans la langue spécifiée (français ou anglais) et l'affiche dans la console.
func DifferentLanguages(filename, langue string) {
	var path string
	// Différenciation entre les chemins pour le français et l'anglais
	if strings.ToLower(langue) == "français" {
		path = filepath.Join("data", "français", filename)
	} else if strings.ToLower(langue) == "anglais" || strings.ToLower(langue) == "english" {
		path = filepath.Join("data", "english", filename)
	} else {
		fmt.Printf("Erreur : Langue inconnue '%s'. Choisissez entre 'français' et 'anglais'.\n", langue)
		return
	}
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Erreur lors de l'ouverture du fichier '%s': %v\n", path, err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier '%s': %v\n", path, err)
	}
}

func AnalyseLangue(motCache, langue string) {
	var path string
	if strings.ToLower(langue) == "français" {
		path = filepath.Join("data", "français", "Files.txt")
	} else if strings.ToLower(langue) == "anglais" || strings.ToLower(langue) == "english" {
		path = filepath.Join("data", "english", "File1")
	} else {
		fmt.Printf("Erreur : Langue inconnue '%s. Choisissez entre 'français' et 'anglais'.\n", langue)
		return
	}
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Erreur lors de l'ouverture du fichier '%v': %v\n", path, err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	found := false
	for scanner.Scan() {
		if strings.ToLower(scanner.Text()) == strings.ToLower(motCache) {
			found = true
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Erreurr llrs de la lecture du fichier '%s' : %v\n", path, err)
	}
	if found {
		fmt.Printf("Le mot caché '%s' est reconnu comme étant en %s.\n", motCache, langue)
	} else {
		fmt.Printf("Le mt caché '%s' n'est pas trouvé dans la liste des mots en %s.\n", motCache, langue)
	}
}

func ReaderUser() game.Tableau {
	/*
		Fonction qui lit le ficher users.txt et qui met les données dans la structure tableau
	*/
	content, error := ioutil.ReadFile("../Data/users.txt") //ouverture du ficher, content contient tout le ficher dans une liste de byte
	if error != nil {
		fmt.Println("Error when opening file")
	}
	users := game.Tableau{
		Pseudos: []game.User{},
	}
	temp := game.User{}
	temp1 := ""
	temp2 := 0
	for _, element := range string(content) { //on regarde tout les element de content en tant que string, ducoup on va analyser un type rune (un seul caratere à la fois)
		switch element {
		case '\n': //cas où l'on saute de ligne : nouvel utilisateur
			temp.Langue = temp1
			temp1 = ""
			users.Pseudos = append(users.Pseudos, temp)
			temp = game.User{}
			temp2 = 0
		case ' ': //cas d'un espace : on change de donnée (pseudo, nbparite joué, score, level, langue)
			temp2++
			if temp2 != 6 {
				switch temp2 {
				case 1:
					temp.Pseudo = temp1
				case 2:
					temp.NbPartieJoué = Atoi(temp1)
				case 3:
					temp.Score = Atoi(temp1)
				case 4:
					temp.Level = Atoi(temp1)
				}
			}
			temp1 = "" //reset de temp quand on change de donnée
		default:
			if element != '\r' { //on ajoute à temp chaque caractere lorsque que l'on change de donnée
				temp1 += string(element)
			}
		}
	}
	return users

}

func Atoi(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("Erreur de conversion de '%s' en entier : %v\n", s, err)
		return 0
	}
	return value
}

//Fonction pour Jonathan, n'hesite pas à regarder le code de ReaderUser(), glhf

func WordReader(langue string, difficulte int) []string {
	var filename string
	if langue == "français" {
		filename = filepath.Join("../Data/Français", fmt.Sprintf("File%d.txt", difficulte))
	} else if langue == "english" {
		filename = filepath.Join("../Data/English", fmt.Sprintf("File%d.txt", difficulte))
	} else {
		fmt.Println("Langue non supportée")
		return nil
	}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier : %v\n", err)
		return nil
	}
	var words []string
	temp := ""
	for _, element := range string(content) {
		if element == '\n' {
			words = append(words, temp)
			temp = ""
		} else {
			if element != '\r' {
				temp += string(element)
			}
		}
	}
	/*
		Fonction qui va renvoyer une liste de string qui contiendra
		tout les mots d'un ficher.txt exemple : WordReader("Français", 3)
		devra lire le ficher file3.txt dans le dossier français.
		Une ligne = un mot
		Nombre de ligne maximum  : 30
	*/
	return words
}

func WordsReadBothFiles(frLangue, enLangue string, difficulte int) []string {
	frWords := WordReader(frLangue, difficulte)
	enWords := WordReader(enLangue, difficulte)
	combinewords := append(frWords, enWords...)
	return combinewords
}

func VictoireReader(victoire bool) []string {
	var filename string
	if victoire {
		filename = "../Data/victoire.txt"
	} else {
		filename = "../Data/defaite.txt"
	}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier :%v\n", err)
		return nil
	}
	var phrases []string
	temp := ""
	for _, element := range string(content) {
		if element == '\n' {
			phrases = append(phrases, temp)
			temp = ""
		} else {
			if element != '\r' {
				temp += string(element)
			}
		}
	}
	/*
		Fonction qui va renvoyer une liste de string qui contiendra
		toutes les phrases du ficher victoire.txt ou du ficher defaite.txt
		en fonction de victoire = true ou false
		(une ligne = une phrase)
		Nombre de ligne maximum  : 40
	*/
	return phrases
}
