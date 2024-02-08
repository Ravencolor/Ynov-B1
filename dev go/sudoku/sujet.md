# Sudoku Resolver

## Instructions

Créez un programme qui résout un sudoku.

Un sudoku valide a une seule solution possible.

Assurez-vous de soumettre tous les fichiers nécessaires pour exécuter le programme.

## Utilisation

### Exemple 1:

Exemple de sortie pour un sudoku valide :

```bash
$ go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
3 9 6 2 4 5 7 8 1$
1 7 8 3 6 9 5 2 4$
5 2 4 8 1 7 3 9 6$
2 8 7 9 5 1 6 4 3$
9 3 1 4 8 6 2 7 5$
4 6 5 7 2 3 9 1 8$
7 1 2 6 3 8 4 5 9$
6 5 9 1 7 4 8 3 2$
8 4 3 5 9 2 1 6 7$
$
```

### Exemple 2:

Exemples de sorties attendues pour des entrées invalides ou des sudokus invalides :

```bash
$ go run . 1 2 3 4 | cat -e
Error$
$ go run . | cat -e
Error$
$ go run . ".96.4...1" "1...6.1.4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
Error$
$
```

## Commentaires

Placez les commentaires nécessaires dans votre code pour expliquer la logique et le fonctionnement du programme.