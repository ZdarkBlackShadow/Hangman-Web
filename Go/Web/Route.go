package web

import (
	datareaderwriter "game/DataReaderWriter"
	game "game/Game"
	"log"
	"net/http"
	"os"
	"strings"
)

func AcceilRoute(w http.ResponseWriter, r *http.Request) {
	if InGame {
		Erreur.Message = "code 403 : accés refusé, you are in game"
		http.Redirect(w, r, "/temporisation", http.StatusSeeOther)
	} else {
		LanguePackage.Logo = "test logo description"
		data := game.TestLangue{
			Pack:        LanguePackage,
			Description: "description",
		}
		err1 := temp.ExecuteTemplate(w, "acceuil", data)
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
		LanguePackage.ImageGame = "Description d'image in game"
		GameData.Pack = LanguePackage
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
	GameData.Phrasefinal = game.RandomString(datareaderwriter.VictoireReader(GameData.Victoire, UserIn.Langue))
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
