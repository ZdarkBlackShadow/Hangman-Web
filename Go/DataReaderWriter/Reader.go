package datareaderwriter

func Reader(filename string) {
	package datareaderwriter

import (
    "bufio"
    "fmt"
    "os"
    "path/filepath"
)

// Reader lit le contenu d'un fichier et l'affiche dans la console.
func Reader(filename string) {
    // Créer le chemin complet vers le fichier
    path := filepath.Join("data", "français", filename)

    file, err := os.Open(path)
    if err != nil {
        fmt.Printf("Erreur lors de l'ouverture du fichier '%s': %v\n", path, err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("Erreur lors de la lecture du fichier '%s': %v\n", path, err)
    }
}

}
