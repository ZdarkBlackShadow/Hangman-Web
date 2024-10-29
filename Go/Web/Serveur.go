package web

import (
	"fmt"
	game "game/Game"
	"html/template"
	"net/http"
	"os"
)

var temp, err = template.ParseGlob("../html/*.html")
var GameData game.GameAffichage
var Erreur game.Erreur = game.Erreur{
	Message: "",
	BackTo:  "Acceuil",
}
var Identified bool = false
var InGame bool = false
var Replay bool = false
var UserIn game.User
var Langue string
var LanguePackage game.LangueText

func InitialisationServeur() {
	if err != nil {
		//erreur lors de l'ouverture des templates
		fmt.Printf("Erreur => %s\n", err.Error())
		os.Exit(02)
	}
	//Initilalisation des routes
	http.HandleFunc("/acceil", AcceilRoute)
	http.HandleFunc("/game", GameRoute)
	http.HandleFunc("/score", ScoreRoute)
	http.HandleFunc("/traitement", TraitementUser)
	http.HandleFunc("/game/traitement", TraitementGame)
	http.HandleFunc("/ending", Ending)
	http.HandleFunc("/temporisation", Temporisation)
	http.HandleFunc("/score/traitement", ScoreTraitement)
	http.HandleFunc("/score/traitement/replay", ScoreTraitementGame)
	http.HandleFunc("/score/traitement/acceil", ScoreTraitementAcceil)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/acceil":
			AcceilRoute(w, r)
		case "/game":
			GameRoute(w, r)
		case "/score":
			ScoreRoute(w, r)
		case "/traitement":
			TraitementUser(w, r)
		case "/game/traitement":
			TraitementGame(w, r)
		case "/ending":
			Ending(w, r)
		case "/temporisation":
			Temporisation(w, r)
		case "/score/traitement/acceil":
			ScoreTraitementAcceil(w, r)
		case "/score/traitement/replay":
			ScoreTraitementGame(w, r)
		default:
			http.Redirect(w, r, "/temporisation", http.StatusTemporaryRedirect)
		}
	})
	Langue = DefaultLangage()
	//LanguePackage = datareaderwriter.PackageLangage(Langue)
	LanguePackage = game.LangueText{
		Inscription: Langue,
		Logo:        "test logo",
	}
	//Initialisation des assets
	fileserver := http.FileServer(http.Dir("../assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	//Initialisation du serveur
	http.ListenAndServe("localhost:8000", nil)
}
