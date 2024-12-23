package datareaderwriter

import (
	"fmt"
	game "game/Game"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

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
			temp.Score = Atoi(temp1)
			temp1 = ""
			users.Pseudos = append(users.Pseudos, temp)
			temp = game.User{}
			temp2 = 0
		case ' ': //cas d'un espace : on change de donnée (pseudo, nbparite joué, score, level, langue)
			temp2++
			if temp2 != 3 {
				switch temp2 {
				case 1:
					temp.Pseudo = temp1
				case 2:
					temp.NbPartieJoué = Atoi(temp1)
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
	filename = filepath.Join("../Data/", langue, fmt.Sprintf("File%d.txt", difficulte))
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
	return words
}

func VictoireReader(victoire bool, langue string) []string {
	var filename string
	if victoire {
		filename = "../Data/" + langue + "/victoire.txt"
	} else {
		filename = "../Data/" + langue + "/defaite.txt"
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

func PackageLangage(langue string) game.LangueText {
	L := game.LangueText{}
	var filepath string
	switch langue {
	case "Français":
		filepath = "../Data/Français/Langue.txt"
	case "English":
		filepath = "../Data/English/Langue.txt"
	case "Español":
		filepath = "../Data/Español/Langue.txt"
	default:
		fmt.Println("Langue non supportée")
		return L
	}
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return L
	}
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	//assignez les lignes aux champs de la structure LangueText
	val := reflect.ValueOf(&L).Elem()
	for i := 0; i < val.NumField(); i++ {
		if i < len(lines) {
			val.Field(i).SetString(lines[i])
		}
	}
	fmt.Println(L.Logo)
	return L
}
