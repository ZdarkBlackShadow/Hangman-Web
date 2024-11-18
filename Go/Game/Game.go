package game

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode"
)

// Variable globale
var Word string
var Difficulte int
var PV int
var ListeMot []string
var ListeLettre []string
var LettreAlreadyFind []rune
var WordToDisplay []string

//Fonction

func GameInit(word string, difficulte int) {
	/*
		Initialisation du jeu
	*/
	fmt.Println(word)
	Word = word
	Difficulte = difficulte
	switch difficulte {
	case 1:
		PV = 12
	case 2:
		PV = 10
	case 3:
		PV = 8
	case 4:
		PV = 6
	case 5:
		PV = 4
	}
	ListeLettre = []string{}
	ListeMot = []string{}
	LettreAlreadyFind = []rune{}
	WordToDisplay = []string{}
	rand.Seed(time.Now().UnixNano())
	RandomLetter := rune(Word[rand.Intn(len(word))])
	LettreAlreadyFind = append(LettreAlreadyFind, RandomLetter)
	for i := 0; i < len(word); i++ {
		if rune(Word[i]) == RandomLetter {
			WordToDisplay = append(WordToDisplay, string(RandomLetter))
		} else {
			WordToDisplay = append(WordToDisplay, "_")
		}
	}
}

func GameLap(submit_answer string) []bool {
	if submit_answer == Word {
		return []bool{true, len(LettreAlreadyFind) <= len(Word)/2}
	}
	if len(submit_answer) > 1 {
		PV -= 2
		ListeMot = append(ListeMot, submit_answer)
		return []bool{false, false}
	}
	letter := rune(submit_answer[0])
	if strings.ContainsRune(Word, letter) {
		if !strings.ContainsRune(string(LettreAlreadyFind), letter) {
			LettreAlreadyFind = append(LettreAlreadyFind, letter)
			for i, char := range Word {
				if char == letter {
					WordToDisplay[i] = string(letter)
				}
			}
		}
		temp := true
		for i, element := range WordToDisplay {
			if element != string(Word[i]) {
				temp = false
			}
		}
		if temp {
			return []bool{true}
		}
	} else {
		PV -= 1
		ListeLettre = append(ListeLettre, submit_answer)
	}
	return []bool{false, false}
	/*
		Verifie si submit_answer est un mot ou une lettre
		et le comparer avec le mot :
		Retirer les point si faux (-2 si le mot est faux, -1 si la lettre est fausse)
		Si la lettre est juste, rajouter à LettreAlreadyFind
		Si le mot est faux, rajouter à ListeMot
		Si la lettre est fausse, rajouter à ListeLettre
		Si le mot est juste ou que toutes les lettres on été trouvé renvoyer true
		Sinon renvoyer False
		Renvoie true ou false si le joueurs à trouvé le mot alors qu'il n'a pas la moitiè des
		lettres de dévoilé.(Fonctionalité bonus)
	*/
}

func RandomString(liste []string) string {
	rand.Seed(time.Now().UnixNano())
	return liste[rand.Intn(len(liste))]
}

func ChoiceLangue(langueChoisi string, answer string) bool {
	langueChoisi = strings.ToLower(langueChoisi)
	if langueChoisi == "français" || langueChoisi == "english" || langueChoisi == "español" {
		fmt.Printf("Langue choisi : %s\n", langueChoisi)
		fmt.Println("Réponse", answer)
		return true
	}
	fmt.Println("Langue non supportée. Veuillez choisir parmis les langue proposée.")
	return false
}

func removeAccents(input string) string {
	var result strings.Builder
	for _, r := range input {
		if unicode.IsLetter(r) {
			r = unicode.ToLower(r)
			switch r {
			case 'à', 'â', 'ä', 'á', 'ã', 'å':
				r = 'a'
			case 'ç':
				r = 'c'
			case 'é', 'è', 'ê', 'ë':
				r = 'e'
			case 'ì', 'í', 'î', 'ï':
				r = 'i'
			case 'ñ':
				r = 'n'
			case 'ò', 'ó', 'ô', 'õ', 'ö':
				r = 'o'
			case 'ù', 'ú', 'û', 'ü':
				r = 'u'
			}
		}
		result.WriteRune(r)
	}
	return result.String()
}

func isValidInput(input string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	return re.MatchString(input)
}
