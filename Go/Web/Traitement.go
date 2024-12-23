package web

import (
	"fmt"
	datareaderwriter "game/DataReaderWriter"
	game "game/Game"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func PseudoVerification(pseudo string) bool {
	temp, err := regexp.MatchString("^[a-zA-Z]{6,32}$", pseudo)
	if err != nil {
		fmt.Println(err)
	}
	return temp
}

func DefaultLangage() string {
	lang := os.Getenv("LANG")
	if strings.HasPrefix(lang, "fr") {
		return "Français"
	} else if strings.HasPrefix(lang, "en") {
		return "English"
	} else if strings.HasPrefix(lang, "es") {
		return "Español"
	} else {
		return "English"
	}
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
	GameData.Pack = datareaderwriter.PackageLangage(UserIn.Langue)
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

	if temp[0] || game.PV <= 0 {
		GameData.Double = temp[1]
		GameData.Finie = true
		GameData.Victoire = game.PV > 0
		InGame = false
		http.Redirect(w, r, "/ending", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/game", http.StatusSeeOther)
	}
}

func ScoreTraitement(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Erreur.Message = "code 403 : Accés refusé"
		Erreur.BackTo = "acceuil"
		http.Redirect(w, r, "/temporisation", http.StatusSeeOther)
	} else {
		fmt.Println(GameData.Victoire)
		if GameData.Victoire {
			if GameData.Double {
				UserIn.Score += (UserIn.Level * 500) * 2
			} else {
				UserIn.Score += UserIn.Level * 500
			}
		}
		datareaderwriter.Writer(UserIn)
		http.Redirect(w, r, "/score", http.StatusSeeOther)
	}
}

func ScoreTraitementGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Erreur.Message = "code 403 : Accés refusé"
		Erreur.BackTo = "acceuil"
		http.Redirect(w, r, "/temporisation", http.StatusSeeOther)
	}
	Identified = true
	game.GameInit(game.RandomString(datareaderwriter.WordReader(UserIn.Langue, UserIn.Level)), UserIn.Level)
	UserIn.Score = 0
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
	UserIn = game.User{}
	http.Redirect(w, r, "/acceil", http.StatusSeeOther)
}
