package datareaderwriter

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Writer(pseudo string, nbParties, score int, difficulte, langue string) {

	file, err := ioutil.ReadFile("./Data/users.txt")
	if err != nil {
		fmt.Printf("Erreur lors de l'ouverture du fichier %v\n", err)
		return
	}
	Lignes := []string{} //contient toutes les lignes du ficher users.txt
	temp := ""
	for _, element := range string(file) {
		if element != '\n' && element != '\r' {
			temp += string(element)
		} else {
			if temp != "" {
				Lignes = append(Lignes, temp)
				temp = ""
			}
		}
	}
	//retour a la ligne
	if temp != "" {
		Lignes = append(Lignes, temp)
	}
	//Recherche le pseudo en question
	userExists := false
	for _, line := range Lignes {
		data := strings.Split(line, ",")
		if len(data) > 0 && data[0] == pseudo {
			userExists = true
			fmt.Printf("Utilisateur %s trouvé. Mise à jour des informations.\n", pseudo)
			data[1] = fmt.Sprintf("%d", nbParties)
			data[2] = fmt.Sprintf("%d", score)
			updatedLine := strings.Join(data, ",")
			for i, l := range Lignes {
				if l == line {
					Lignes[i] = updatedLine
					break
				}
			}
			break
		}
	}

	if !userExists {
		fmt.Printf("Erreur : l'utilisateur %s est inconnu.\n", pseudo)
		//message d'erreur si l'utilisateur n'existe pas
		return
	}
	err = ioutil.WriteFile("./Data/users.txt", []byte(strings.Join(Lignes, "\n")), 0644)
	if err != nil {
		fmt.Printf("Erreur lors de l'écriture du fichier : %v\n", err)
		return
	}
}

//analyse ligne pour savoir si l'utlisateur existe déja
/*
	// Construire les données à écrire
	data := fmt.Sprintf("Pseudo: %s, Parties jouées: %d, Score: %d, Difficulté: %s, Langue: %s",
		pseudo, nbParties, score, difficulte, langue)
	// Écrire les données dans le fichier
	if _, err := file.WriteString(data + "\n"); err != nil {
		fmt.Printf("Erreur lors de l'écriture dans le fichier '%s': %v\n", path, err)
	}
*/
