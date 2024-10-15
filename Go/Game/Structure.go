package game

//Ficher où l'on va mettre toute les tructures

type User struct {
	Pseudo       string
	Level        int
	Langue       string
	NbPartieJoué int
	Score        int
}

type GameAffichage struct {
	Start                 bool
	DerniereEssaieReussie bool
	Mot                   string
	ListeLettre           []string
	ListeMots             []string
	PvRestant             int
	Finie                 bool
	Victoire              bool
}

type Tableau struct {
	Pseudos []User
}
