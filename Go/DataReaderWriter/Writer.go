package datareaderwriter

import (
	"fmt"
	game "game/Game"
	"io/ioutil"
	"os"
	"strconv"
)

func Writer(u game.User) {
	content, error := ioutil.ReadFile("../Data/users.txt")
	if error != nil {
		fmt.Println("Error when opening file")
	}
	res := []game.User{}
	temp := game.User{}
	temp1 := ""
	temp2 := 0
	for _, element := range string(content) {
		switch element {
		case '\n':
			temp.Langue = temp1
			temp1 = ""
			res = append(res, temp)
			temp = game.User{}
			temp2 = 0
		case ' ':
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
			temp1 = ""
		default:
			if element != '\r' {
				temp1 += string(element)
			}
		}
	}
	//fin de lecture users.txt
	Find := false
	for _, ele := range res {
		if ele.Pseudo == u.Pseudo {
			Find = true
			ele.NbPartieJoué++
			ele.Score += u.Score
			break
		}
	}
	if !Find {
		u.NbPartieJoué++
		res = append(res, game.User{
			Pseudo:       u.Pseudo,
			Score:        u.Score,
			NbPartieJoué: u.NbPartieJoué,
			Level:        u.Level,
			Langue:       u.Langue,
		})
	}
	//preparation du contenu
	contenu := ""
	for _, element := range res {
		contenu += element.Pseudo + " " + strconv.Itoa(element.NbPartieJoué) + " " + strconv.Itoa(element.Score) + " " + strconv.Itoa(element.Level) + " " + element.Langue + "\n"
	}
	//ecriture dans le ficher
	file, err := os.OpenFile("../Data/users.txt", os.O_WRONLY|os.O_TRUNC, 0644)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	_, err = file.WriteString(contenu)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier:", err)
		return
	}
}
