package datareaderwriter

import (
	"fmt"
	"io/ioutil"
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
}
