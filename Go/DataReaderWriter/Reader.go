package datareaderwriter

import (
	"bufio"
	"fmt"
	game "game/Game"
	"os"
	"path/filepath"
	"strconv"
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

func ReaderUser() game.Tableau {

	filePath := "Data/users.txt"
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Errorf("erreur d'ouverture du fichier : %v\n", err)
	}
	defer file.Close()
	var users game.Tableau = game.Tableau{
		Pseudos: []game.User{},
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 5 {
			fmt.Println("Ligne invalide dans le fichier :", line)
			continue
		}
		user := game.User{
			Pseudo:       parts[0],
			NbPartieJoué: atoi(parts[1]),
			Score:        atoi(parts[2]),
			Level:        atoi(parts[3]),
			Langue:       parts[4],
		}
		users.Pseudos = append(users.Pseudos, user)
		fmt.Printf("Erreur de lecture^du fichier: %v\n", err)

	}
	return users

}

func atoi(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("Erreur de conversion de '%s' en entier : %v\n", s, err)
		return 0
	}
	return value
}
