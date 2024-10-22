package datareaderwriter

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

// Reader lit le contenu d'un fichier dans la langue spécifiée (français ou anglais) et l'affiche dans la console.
func DifferentLanguages(filename, langue string) {
	var path string

	// Différenciation entre les chemins pour le français et l'anglais
	if strings.ToLower(langue) == "français" {
		path = filepath.Join("data", "français", filename)
	} else if strings.ToLower(langue) == "anglais" || strings.ToLower(langue) == "english" {
		path = filepath.Join("data", "english", filename)
	} else {
		fmt.Printf("Erreur : Langue inconnue '%s'. Choisissez entre 'français' et 'anglais'.\n", langue)
		return
	}

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

func AnalyseLangue(motCache, langue string) {
	var path string

	if strings.ToLower(langue) == "français" {
		path = filepath.Join("data", "français", "Files.txt")
	} else if strings.ToLower(langue) == "anglais" || strings.ToLower(langue) == "english" {
		path = filepath.Join("data", "english", "File1")
	} else {
		fmt.Printf("Erreur : Langue inconnue '%s. Choisissez entre 'français' et 'anglais'.\n", langue)
		return
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Erreur lors de l'ouverture du fichier '%v': %v\n", path, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	found := false
	for scanner.Scan() {
		if strings.ToLower(scanner.Text()) == strings.ToLower(motCache) {
			found = true
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Erreurr llrs de la lecture du fichier '%s' : %v\n", path, err)
	}

	if found {
		fmt.Printf("Le mot caché '%s' est reconnu comme étant en %s.\n", motCache, langue)
	} else {
		fmt.Printf("Le mt caché '%s' n'est pas trouvé dans la liste des mots en %s.\n", motCache, langue)
	}
}
