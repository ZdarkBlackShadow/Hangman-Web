package game

// Structure pour représenter un utilisateur
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
	EssaisRestants        int  // Ajout pour gérer les essais restants dans la partie
	Score                 int  // Ajout du score directement dans GameAffichage pour suivre les points du joueur
	MotActuel             string // Ajout pour stocker le mot à deviner dans la partie
}

type Tableau struct {
	Pseudos []User
}
