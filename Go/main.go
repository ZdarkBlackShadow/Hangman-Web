package main

import (
	"fmt"
	datareaderwriter "game/DataReaderWriter"
	game "game/Game"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
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

/*{
	Start:                 true,
	DerniereEssaieReussie: false,
	Mot:                   "test",
	ListeLettre:           []string{"a", "g"},
	ListeMots:             []string{""},
	PvRestant:             10,
	Finie:                 true,
	Victoire:              false,
}*/

func AcceilRoute(w http.ResponseWriter, r *http.Request) {
	if InGame {
		Erreur.Message = "code 403 : accés refusé, you are in game"
		http.Redirect(w, r, "/temporisation", http.StatusSeeOther)
	} else {
		err1 := temp.ExecuteTemplate(w, "acceuil", nil)
		if err1 != nil {
			log.Fatal(err1)
		}
	}
}

func GameRoute(w http.ResponseWriter, r *http.Request) {
	if !Identified {
		Erreur.Message = "code 403 : accès refusé"
		Erreur.BackTo = "acceuil"
		http.Redirect(w, r, "/temporisation", http.StatusSeeOther)
	} else {
		InGame = true
		err1 := temp.ExecuteTemplate(w, "game", GameData)
		if err1 != nil {
			log.Fatal(err1)
		}
	}
}

func ScoreRoute(w http.ResponseWriter, r *http.Request) {
	//A remplacer par une fonction dans Reader.go (fonction qui lit le fucher users.txt et qui renvoie sous forme de structure game.User)
	if InGame {
		Erreur.Message = "code 403 : accés refusé, you are in game"
		Erreur.BackTo = "game"
		http.Redirect(w, r, "/temporisation", http.StatusSeeOther)
	} else {
		data := datareaderwriter.ReaderUser()
		err1 := temp.ExecuteTemplate(w, "score", data)
		if err1 != nil {
			log.Fatal(err1)
		}
	}
}

func PseudoVerification(pseudo string) bool {
	temp, err := regexp.MatchString("^[a-zA-Z]{6,32}$", pseudo)
	if err != nil {
		fmt.Println(err)
	}
	return temp
}

func TraitementUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Erreur.Message = "code 403 : Accés refusé"
		Erreur.BackTo = "acceuil"
		http.Redirect(w, r, "/temporisation", http.StatusSeeOther)
	}
	if !PseudoVerification(r.FormValue("pseudo")) {
		// Gestion de l'erreur : mauvais pseudo
		Erreur.Message = "code 401 : Pseudo incorect"
		Erreur.BackTo = "acceuil"
		http.Redirect(w, r, "/temporisation", http.StatusSeeOther)
	}
	UserIn = game.User{
		Pseudo:       r.FormValue("pseudo"),
		Langue:       r.FormValue("langue"),
		Level:        datareaderwriter.Atoi(r.FormValue("level")),
		NbPartieJoué: 0,
		Score:        0,
	}
	Identified = true
	temp1, _ := strconv.Atoi(r.FormValue("level"))
	game.GameInit(game.RandomString(datareaderwriter.WordReader(r.FormValue("langue"), temp1)), temp1)
	GameData = game.GameAffichage{
		Start:                 true,
		DerniereEssaieReussie: false,
		Mot:                   game.WordToDisplay,
		ListeLettre:           game.ListeLettre,
		ListeMots:             game.ListeMot,
		PvRestant:             game.PV,
		Finie:                 false,
		Victoire:              false,
	}
	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

func TraitementGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Erreur.Message = "code 403 : Accés refusé"
		Erreur.BackTo = "acceuil"
		http.Redirect(w, r, "/temporisation", http.StatusSeeOther)
	}
	//traitement de la reponse de l'utilisateur
	temp1 := []string{}
	for _, ele := range game.WordToDisplay {
		temp1 = append(temp1, ele)
	}
	temp := game.GameLap(r.FormValue("answer"))
	temp2 := false
	for i, element := range game.WordToDisplay {
		if element != temp1[i] {
			temp2 = true
			break
		}
	}
	GameData = game.GameAffichage{
		Start:                 false,
		DerniereEssaieReussie: temp2,
		Mot:                   game.WordToDisplay,
		ListeLettre:           game.ListeLettre,
		ListeMots:             game.ListeMot,
		PvRestant:             game.PV,
		Finie:                 false,
		Victoire:              false,
	}

	if temp || game.PV <= 0 {
		GameData.Finie = true
		GameData.Victoire = game.PV > 0

		datareaderwriter.Writer(UserIn)
		InGame = false
		http.Redirect(w, r, "/ending", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/game", http.StatusSeeOther)
	}
}

func Ending(w http.ResponseWriter, r *http.Request) {
	if !GameData.Finie {
		if !Identified {
			Erreur.Message = "code 403 : Not identified"
			Erreur.BackTo = "acceuil"
			http.Redirect(w, r, "/temporisation", http.StatusSeeOther)
		} else {
			Erreur.Message = "code 403 : accés refusé, you're in game"
			Erreur.BackTo = "game"
			http.Redirect(w, r, "/temporisation", http.StatusSeeOther)
		}
	}
	GameData.Phrasefinal = game.RandomString(datareaderwriter.VictoireReader(GameData.Victoire))
	err1 := temp.ExecuteTemplate(w, "ending", GameData)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func Temporisation(w http.ResponseWriter, r *http.Request) {
	//rajouter une condition si l'utilisateur s'était déja authentifié et mettre sa langue
	if Erreur.Message == "" {
		lang := os.Getenv("LANG")
		if strings.HasPrefix(lang, "fr") {
			Erreur.Message = "code 404 : Mauvaise URL"
		} else if strings.HasPrefix(lang, "en") {
			Erreur.Message = "code 404 : Wrong URL"
		} else if strings.HasPrefix(lang, "es") {
			Erreur.Message = "código 404 : URL incorrecta"
		} else if strings.HasPrefix(lang, "it") {
			Erreur.Message = "codice  404 : URL errato"
		} else if strings.HasPrefix(lang, "nl") {
			Erreur.Message = "code 404 : slechte URL"
		} else {
			Erreur.Message = "code 404 : Wrong URL"
		}
		Erreur.BackTo = "acceuil"
	}
	err1 := temp.ExecuteTemplate(w, "temporisation", Erreur)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func ScoreTraitementGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Erreur.Message = "code 403 : Accés refusé"
		Erreur.BackTo = "acceuil"
		http.Redirect(w, r, "/temporisation", http.StatusSeeOther)
	}
	Identified = true
	temp1, _ := strconv.Atoi(r.FormValue("level"))
	game.GameInit("hello", temp1)
	GameData = game.GameAffichage{
		Start:                 true,
		DerniereEssaieReussie: false,
		Mot:                   game.WordToDisplay,
		ListeLettre:           game.ListeLettre,
		ListeMots:             game.ListeMot,
		PvRestant:             game.PV,
		Finie:                 false,
		Victoire:              false,
	}
	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

func ScoreTraitementAcceil(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Erreur.Message = "code 403 : Accés refusé"
		Erreur.BackTo = "acceuil"
		http.Redirect(w, r, "/temporisation", http.StatusSeeOther)
	}
	Identified = false
	http.Redirect(w, r, "/acceil", http.StatusSeeOther)
}

func OpenRoads() {
	http.HandleFunc("/acceil", AcceilRoute)
	http.HandleFunc("/game", GameRoute)
	http.HandleFunc("/score", ScoreRoute)
	http.HandleFunc("/traitement", TraitementUser)
	http.HandleFunc("/game/traitement", TraitementGame)
	http.HandleFunc("/ending", Ending)
	http.HandleFunc("/temporisation", Temporisation)
	http.HandleFunc("/score/traitement/replay", ScoreTraitementGame)
	http.HandleFunc("/score/traitement/acceil", ScoreTraitementAcceil)
}

func main() {
	if err != nil {
		//erreur lors de l'ouverture des templates
		fmt.Printf("Erreur => %s\n", err.Error())
		os.Exit(02)
	}
	OpenRoads()
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
	//Initialisation des assets
	fileserver := http.FileServer(http.Dir("../assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	//Initialisation du serveur
	http.ListenAndServe("localhost:8000", nil)
}
