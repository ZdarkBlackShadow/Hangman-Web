package main

import (
	"fmt"
	datareader "game/DataReader"
	game "game/Game"
	"html/template"
	"log"
	"net/http"
	"os"
)

var temp, err = template.ParseGlob("../html/*.html")

func AcceilRoute(w http.ResponseWriter, r *http.Request) {
	err1 := temp.ExecuteTemplate(w, "acceuil", nil)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func GameRoute(w http.ResponseWriter, r *http.Request) {
	data := game.GameAffichage{
		Start:                 true,
		DerniereEssaieReussie: false,
		Mot:                   "test",
		ListeLettre:           []string{"a", "g"},
		ListeMots:             []string{"salut"},
		PvRestant:             10,
		Finie:                 false,
		Victoire:              false,
	}
	err1 := temp.ExecuteTemplate(w, "game", data)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func ScoreRoute(w http.ResponseWriter, r *http.Request) {
	Users := []game.User{
		game.User{
			Pseudo: "Adrien",
			Level:  10,
			Langue: "Français",
		},
		game.User{
			Pseudo: "Jonathan",
			Level:  9,
			Langue: "English",
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

func main() {
	game.Game()
	datareader.Reader("test")
	if err != nil {
		fmt.Printf("Erreur => %s\n", err.Error())
		os.Exit(02)
	}
	http.HandleFunc("/acceil", AcceilRoute)
	http.HandleFunc("/game", GameRoute)
	http.HandleFunc("/score", ScoreRoute)
	//Initialisation des assets
	fileserver := http.FileServer(http.Dir("../assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	//Initialisation du serveur
	http.ListenAndServe("localhost:8000", nil)
}
