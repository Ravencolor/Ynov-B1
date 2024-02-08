# YTrack

Ce projet est basé sur le célèbre jeu du Pendu (HangMan).

Vous devez créer un dépôt privé avec le nom `hangman-classic`.

Notions :
- Documentation Golang : ioutil
- Documentation Golang : rand
- Exemple de lecture de fichier

## Instructions :
Créez un programme `hangman` qui prendra un fichier en paramètre. Créez un fichier `words.txt` qui contient plusieurs mots avec lesquels le programme jouera. Chaque mot est séparé par un retour à la ligne `\n`. Vous pouvez également trouver différents dictionnaires [ici](https://www.mit.edu/~ecprice/wordlist.10000).

### PARTIE 1 —
Vous aurez 10 tentatives pour terminer le jeu.

1. Tout d'abord, le programme choisira de manière aléatoire un mot dans le fichier.
2. Le programme révélera n lettres aléatoires dans le mot, où n est égal à `len(word) / 2 - 1`.
3. Le programme lira l'entrée standard pour proposer une lettre.
4. Si la lettre n'est pas présente, il affichera un message d'erreur et le nombre de tentatives diminuera (de 10 à 9, etc., jusqu'à 0).
5. Si la lettre est présente, il révélera toutes les lettres correspondantes dans le mot.
6. Le programme continue jusqu'à ce que le mot soit trouvé ou que le nombre de tentatives soit épuisé.

### PARTIE 2 —
Appelons José le pauvre homme qui sera pendu à cette corde si vous perdez.

Vous recevrez un fichier nommé `hangman.txt` qui contient toutes les positions de José. Ce fichier contient 10 positions correspondant aux 10 tentatives nécessaires pour terminer le jeu.

Vous devrez analyser ce fichier et afficher la position appropriée de José à mesure que le nombre de tentatives diminue.

Chaque position a :

- Une hauteur de 7 lignes, se terminant par un retour à la ligne (donc 8 au total)
- Une longueur de 9 caractères, se terminant par un retour à la ligne (donc 10 au total)

Voici un exemple de la septième position de José :

```plaintext
  +---+  
  |   |  
  O   |  
 /|   |  
      |  
      |  
=========
```

### Aide :
Si vous ne savez pas comment manipuler les données de votre programme, vous pouvez toujours essayer d'utiliser une structure et un pointeur vers une structure dans les paramètres de vos futures fonctions.

C'est juste un exemple, vous êtes libre de ne pas l'utiliser, de le modifier ou de créer le vôtre.

C'est comme vous le souhaitez, et n'hésitez pas à utiliser ce avec quoi vous êtes à l'aise pour réaliser ce projet.

```go
type HangManData struct {
	Word             string          // Mot composé de '_', ex : H_ll_
	ToFind           string          // Mot final choisi par le programme au début. C'est le mot à trouver.
	Attempts         int             // Nombre de tentatives restantes
	HangmanPositions [10]string      // Cela peut être le tableau où les positions analysées dans "hangman.txt" sont stockées.
}
```

Bonne chance !

Paquets autorisés :
Seuls les paquets standard de Go sont autorisés.

### Utilisation

```bash
# Exécutez le serveur hangman-web avec le fichier words.txt en paramètre
$> go run main.go

Bonne chance, vous avez 10 tentatives.
_ _ _ _ O

Choisissez : E
_ E _ _ O

Choisissez : A
Non présent dans le mot, 9 tentatives restantes

         
         
         
         
         
         
=========
Choisissez : L
_ E L L O 

Choisissez : B
Non présent dans le mot, 8 tentatives restantes
         
      |  
      |  
      |  
      |  
      |  
=========
Choisissez : Z
Non présent dans le mot, 7 tentatives restantes
  +---+  
      |  
      |  
      |  
      |  
      |  
=========
Choisissez : H
H E L L O

Félicitations !
```