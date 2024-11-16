# 🌟 Hangman Web 🌟
<div>
  <img src="/assets/image/logo.webp" alt="Logo" width="198">
  <img src="/assets/image/game.webp" alt="Logo" width="198">
  <img src="/assets/image/game/game12.webp" alt="Logo" width="198">
  <img src="/assets/image/game/game6.webp" alt="Logo" width="198">
  <img src="/assets/image/game/game0.webp" alt="Logo" width="198">
</div>

# ⚙️Prérequis
- Avoir Git ([Download Git](https://git-scm.com/downloads))
- Avoir goland (1.23.2) ([Download Goland](https://go.dev/dl))

# 🛠️ Installation
- Ouvrez un terminal dans le dossier où vous voulez installer le projet (Allez dans votre Explorateur de fichiers, ouvrez le dossier dans lequel vous voulez stockez le projet, faites un clic droit, Afficher plus d'options, Ouvrir dans le terminal.)
- Taper la commande ci-dessous
```bash
git clone https://github.com/ZdarkBlackShadow/Hangman-Web.git
```

# 🚀 Lancer le projet
- Ouvrez un terminal dans le dossier où vous stocker le projet (Allez dans votre Explorateur de fichiers, ouvrez le dossier dans lequel vous stockez le projet, faites un clic droit, Afficher plus d'options, Ouvrir dans le terminal.)
- Puis tapez les commandes ci'dessous
```bash
 cd Go
```
```bash
 go run main.go
```
- puis dans votre navigateur internet, tapez l'URl suivante :
```bash
 http://localhost:8000
```
# 🌐 Routes de l'application
- Routes qui distribuent des vues :

| Méthode  | Route          | Description |
| :---------------: |:---------------:| :----------------:|
| Get  | /acceil | Page d'accueil qui propose de s'identifier, choisir les langues et la difficulté |
| Get  | /game | Page du jeu qui n'est qu'accesible que lorsqu'on est identifié |
| Get  | /score | Page qui affiche un tableau des scores |
| Get  | /ending | Page qui affiche le résulatat à la fin du jeu |
| Get  | /temporisation | Page qui s'affiche quand on as tapé le mauvais url ou que l'on esseye d'acceder à une page et qu'on y a pas le droit |

- Routes qui traitent des données :

| Méthode  | Route          | Description |
| :---------------: |:---------------:| :----------------:|
| Post  | /traitement | Traitement des données de l'utilisateur |
| Post  | /game/traitement | Traitement des données du jeu |
| Post  | /score/traitement | Traitement des données après la fin du jeu |
| Post  | /score/traitement/replay | Traitement des données si l'utilisateur veut rejouer |
| post  | /score/traitement/acceil | Traitement des données si l'utilsateur veus retourner à l'acceuil |

# 👥 Équipe du projet
- Adrien Lecomte  :  [Github](https://github.com/ZdarkBlackShadow)
- Jonathan Perez  :  [Github](https://github.com/Jonathan-p-z)
