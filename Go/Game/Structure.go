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
	Mot                   []string
	ListeLettre           []string
	ListeMots             []string
	PvRestant             int
	Finie                 bool
	Victoire              bool
	EssaisRestants        int    // Ajout pour gérer les essais restants dans la partie
	Score                 int    // Ajout du score directement dans GameAffichage pour suivre les points du joueur
	MotActuel             string // Ajout pour stocker le mot à deviner dans la partie
	Phrasefinal           string
	Pack                  LangueText
}

type Tableau struct {
	Pseudos []User
	Pack    LangueText
	GameEnd bool
}

type Erreur struct {
	Message string
	BackTo  string
	Pack    LangueText
}

type TestLangue struct {
	Pack        LangueText
	Description string
}

type LangueText struct {
	//Header
	Jeu   string
	Score string
	//acceuil
	Inscription string
	Pseudo      string
	Level       string
	Langue      string
	Valider     string
	//Game
	PhraseStart string
	BienJoue    string
	Dommage     string
	Valider2    string
	ListeLettre string
	ListeMots   string
	//Score
	Tableau  string
	Pseudo2  string
	Score2   string
	NbPartie string
	Rejouer  string
	Retour   string
	//temporisation
	Retour2    string
	RetourGame string
	//Ending
	TabelauScore string
	//erreur
	WrongUrl string
	//pour les aveugles
	ImageGame     string //à modifier en fonction des hp restants
	Imagedefaite  string
	Imagevictoire string
	ImageTemp     string
	Logo          string
}
