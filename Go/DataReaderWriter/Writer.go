package datareaderwriter

import (
    "fmt"
    "os"
    "path/filepath"
	"io/ioutil"
)

// Writer écrit les données du joueur dans un fichier spécifié, en fonction de la langue.
func Writer(filename, pseudo string, nbParties, score int, difficulte, langue string) {
    var Folder string

    // Choisir le bon dossier en fonction de la langue
    switch langue {
    case "Français":
        Folder = "français"
    case "Anglais":
        Folder = "english"
    default:
        fmt.Println("Langue non reconnue. Veuillez choisir entre 'Français' ou 'Anglais'.")
        return
    }

    // Créer le chemin complet vers le fichier dans le dossier correct
    path := filepath.Join("data", Folder + ".txt", filename)

    file, err := ioutil.ReadFile("./Data/" + Folder + "/File1.txt")
    if err != nil {
        fmt.Printf("Erreur lors de l'ouverture du fichier '%s': %v\n", path, err)
        return
    }
    // Construire les données à écrire
    data := fmt.Sprintf("Pseudo: %s, Parties jouées: %d, Score: %d, Difficulté: %s, Langue: %s",
        pseudo, nbParties, score, difficulte, langue)

    // Écrire les données dans le fichier
    if _, err := file.WriteString(data + "\n"); err != nil {
        fmt.Printf("Erreur lors de l'écriture dans le fichier '%s': %v\n", path, err)
    }
}
