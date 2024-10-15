package main

import (
	"fmt"
	datareaderwriter "game/DataReaderWriter"
	game "game/Game"
	"html/template"
	"log"
	"net/http"
	"os"
)

var temp, err = template.ParseGlob("../html/*.html")
var GameData = game.GameAffichage{
	Start:                 true,
	DerniereEssaieReussie: false,
	Mot:                   "test",
	ListeLettre:           []string{"a", "g"},
	ListeMots:             []string{""},
	PvRestant:             10,
	Finie:                 true,
	Victoire:              false,
}

func AcceilRoute(w http.ResponseWriter, r *http.Request) {
	err1 := temp.ExecuteTemplate(w, "acceuil", nil)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func GameRoute(w http.ResponseWriter, r *http.Request) {
	err1 := temp.ExecuteTemplate(w, "game", GameData)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func ScoreRoute(w http.ResponseWriter, r *http.Request) {
	Users := []game.User{
		game.User{
			Pseudo:       "Adrien",
			Level:        10,
			Langue:       "Français",
			NbPartieJoué: 9,
			Score:        150,
		},
		game.User{
			Pseudo:       "Jonathan",
			Level:        9,
			Langue:       "English",
			NbPartieJoué: 12,
			Score:        300,
		},
	}
	data := game.Tableau{
		Pseudos: Users,
	}
	err1 := temp.ExecuteTemplate(w, "score", data)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func TraitementUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.FormValue("pseudo"))
	http.Redirect(w, r, "/game", http.StatusSeeOther)
	if r.FormValue("pseudo") != "tintin" {
		//erreur, mauvais pseudo
	}
}

func TraitementGame(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.FormValue("answer"))
	GameData.ListeMots = append(GameData.ListeMots, r.FormValue("answer"))
	//traitement de la reponse de l'utilisateur
	if GameData.Finie {
		http.Redirect(w, r, "/ending", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/game", http.StatusSeeOther)
	}
}

func Ending(w http.ResponseWriter, r *http.Request) {
	err1 := temp.ExecuteTemplate(w, "ending", GameData)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func main() {
	game.Game()
	datareaderwriter.Reader("test")
	if err != nil {
		fmt.Printf("Erreur => %s\n", err.Error())
		os.Exit(02)
	}
	http.HandleFunc("/acceil", AcceilRoute)
	http.HandleFunc("/game", GameRoute)
	http.HandleFunc("/score", ScoreRoute)
	http.HandleFunc("/traitement", TraitementUser)
	http.HandleFunc("/game/traitement", TraitementGame)
	http.HandleFunc("/ending", Ending)
	//Initialisation des assets
	fileserver := http.FileServer(http.Dir("../assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	//Initialisation du serveur
	http.ListenAndServe("localhost:8000", nil)
}
