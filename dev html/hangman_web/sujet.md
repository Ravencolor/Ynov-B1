## hangman-web

Le projet `hangman-web` consiste à créer et exécuter un serveur dans lequel il sera possible d'utiliser une interface utilisateur graphique (GUI) web de votre dernier projet, `hangman-classic`.

### Go Module

Vous devez utiliser un module Go pour appeler vos fonctions depuis `hangman-classic`. Vous ne devez pas copier/coller ou les réécrire dans ce nouveau projet `hangman-web`.

Vous devrez créer un dépôt privé avec le nom `hangman-web`.

### Notions

- Documentation Golang : net
- Documentation Golang : ioutil
- Documentation Golang : rand
- Documentation de l'exemple web Go : templates
- Documentation Golang : templates

### Objectifs

Créez un programme `hangman-web` qui prendra un fichier `words.txt` en paramètre. Créez un fichier `words.txt` qui contient plusieurs mots avec lesquels le programme jouera. Chaque mot est séparé par un retour à la ligne.

Le comportement du jeu est le même que le projet `hangman`, référez-vous à celui-ci pour plus de détails.

Dans ce projet, vous devrez implémenter au moins les points d'accès suivants :

1. `GET /` : Envoie une réponse HTML - la page principale, elle affichera essentiellement votre interface.
   1.1. Conseil pour GET : utilisez les modèles Go pour recevoir et afficher les données du serveur.

2. `POST /hangman` : Envoie des données au serveur Golang (la lettre que vous voulez trouver).
   2.1. Conseil pour POST : utilisez les balises de formulaire et d'autres types de balises pour effectuer la requête POST. Le formulaire doit rediriger vers `/hangman`.

La page principale doit comporter au moins :

- Un texte représentant le mot à révéler.
- Une zone de texte.
- Un bouton qui envoie une requête POST à `/hangman` et affiche le résultat sur la page.

### Paquets autorisés

Seuls les paquets standard de Go sont autorisés.

### Pas d'utilisation de Framework HTML/CSS

### Instructions

- Le serveur HTTP doit être écrit en Go.
- Les modèles HTML doivent être dans le répertoire de la racine du projet `templates`.
- Le code doit respecter les bonnes pratiques.
- Utilisez le projet `Hangman` en tant que package importé.

### Utilisation

```bash
# Exécutez le serveur hangman-web avec le fichier words.txt en paramètre
$> go run hangman-web.go words.txt
```
Puis aller sur localhost8080