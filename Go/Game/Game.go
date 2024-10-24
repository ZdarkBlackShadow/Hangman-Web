package game

import (
	"math/rand"
	"time"
)

// Variable globale
var Word string
var Difficulte int
var PV int
var ListeMot []string
var ListeLettre []string
var LettreAlreadyFind []rune
var WordToDisplay []rune

//Fonction

func GameInit(word string, difficulte int) {
	/*
		Initialisation du jeu
	*/
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
	WordToDisplay = []rune{}
	rand.Seed(time.Now().UnixNano())
	RandomLetter := rune(Word[rand.Intn(len(word))])
	LettreAlreadyFind = append(LettreAlreadyFind, RandomLetter)
	for i := 0; i < len(word); i++ {
		if rune(Word[i]) == RandomLetter {
			WordToDisplay = append(WordToDisplay, RandomLetter)
		} else {
			WordToDisplay = append(WordToDisplay, '_')
		}
	}
}

func GameLap(submit_answer string) bool {
	/*
		Verifie si submit_answer est un mot ou une lettre
		et le comparer avec le mot :
		Retirer les point si faux (-2 si le mot est faux, -1 si la lettre est fausse)
		Si la lettre est juste, rajouter à LettreAlreadyFind
		Si le mot est faux, rajouter à ListeMot
		Si la lettre est fausse, rajouter à ListeLettre
		Si le mot est juste ou que toutes les lettres on été trouvé renvoyer true
		Sinon renvoyer False
	*/
	return false //a remplacer avec le code
}
