# 🌟 Hangman Web 🌟
<div>
  <img src="/assets/image/logo.webp" alt="Logo" width="198">
  <img src="/assets/image/game.webp" alt="Image du jeu" width="198">
  <img src="/assets/image/game/game12.webp" alt="Image quand on à toute les PV" width="198">
  <img src="/assets/image/game/game6.webp" alt="Image quand on a la moitié des PV" width="198">
  <img src="/assets/image/game/game0.webp" alt="Image quand on est mort dans le jeu" width="198">
</div>

# ⚙️Prérequis
- Avoir Git ([Télécharger Git](https://git-scm.com/downloads))
- Avoir goland (1.23.2) ([Télécharger Goland](https://go.dev/dl))

# 🛠️ Installation
- Ouvrez un terminal dans le dossier où vous voulez installer le projet (Allez dans votre Explorateur de fichiers, ouvrez le dossier dans lequel vous voulez stocker le projet, faite un clic droit, Afficher plus d'options, Ouvrir dans le terminal.)
- Taper la commande ci-dessous
```bash
git clone https://github.com/ZdarkBlackShadow/Hangman-Web.git
```

# 🚀 Lancer le projet
- Ouvrez un terminal dans le dossier où vous stockez le projet (Allez dans votre Explorateur de fichiers, ouvrez le dossier dans lequel vous stockez le projet, faites un clic droit, Afficher plus d'options, Ouvrir dans le terminal.)
- Puis tapez les commandes ci-dessous
```bash
 cd Go
```
```bash
 go run main.go
```
- puis dans votre navigateur internet, tapez l'URl suivante :
```bash
 http://localhost:8000/acceil
```
# 🌐 Routes de l'application
- Routes qui distribuent des vues :

| Méthode  | Route          | Description |
| :---------------: |:---------------:| :----------------:|
| Get  | /acceil | Page d'accueil qui propose de s'identifier, choisir les langues et la difficulté |
| Get  | /game | Page du jeu qui n'est accessible que lorsqu'on est identifié |
| Get  | /score | Page qui affiche un tableau des scores |
| Get  | /ending | Page qui affiche le résultat à la fin du jeu |
| Get  | /temporisation | Page qui s'affiche quand on a tapé le mauvais URL ou que l'on essaie d'accéder à une page et qu'on n'y a pas le droit |

- Routes qui traitent des données :

| Méthode  | Route          | Description |
| :---------------: |:---------------:| :----------------:|
| Post  | /traitement | Traitement des données de l'utilisateur |
| Post  | /game/traitement | Traitement des données du jeu |
| Post  | /score/traitement | Traitement des données après la fin du jeu |
| Post  | /score/traitement/replay | Traitement des données si l'utilisateur veut rejouer |
| post  | /score/traitement/acceil | Traitement des données si l'utilisateur veut retourner à l'accueil |

# 👥 Équipe du projet
- Adrien Lecomte  :  [Github](https://github.com/ZdarkBlackShadow)
- Jonathan Perez  :  [Github](https://github.com/Jonathan-p-z)
