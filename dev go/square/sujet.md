## quadA

### Instructions

Écrivez une fonction `QuadA` qui imprime un rectangle valide avec une largeur `x` et une hauteur `y` données.

La fonction doit dessiner les rectangles comme dans les exemples.

Si `x` et `y` sont des nombres positifs, le programme doit imprimer les rectangles comme indiqué dans les exemples. Sinon, la fonction ne doit rien imprimer.

### Fonction attendue

```go
func QuadA(x, y int) {

}
```

### Utilisation

Voici des programmes possibles pour tester votre fonction :

#### Programme #1

```go
package main

import "piscine"

func main() {
	piscine.QuadA(5, 3)
}
```

Et sa sortie :

```plaintext
$ go run .
o---o
|   |
o---o
$
```

#### Programme #2

```go
package main

import "piscine"

func main() {
	piscine.QuadA(5, 1)
}
```

Et sa sortie :

```plaintext
$ go run .
o---o
$
```

#### Programme #3

```go
package main

import "piscine"

func main() {
	piscine.QuadA(1, 1)
}
```

Et sa sortie :

```plaintext
$ go run .
o
$
```

#### Programme #4

```go
package main

import "piscine"

func main() {
	piscine.QuadA(1, 5)
}
```

Et sa sortie :

```plaintext
$ go run .
o
|
|
|
o
$
```

## quadB

### Instructions

Écrivez une fonction `QuadB` qui imprime un rectangle valide de largeur `x` et de hauteur `y`.

La fonction doit dessiner les rectangles comme dans les exemples.

Si `x` et `y` sont des nombres positifs, le programme doit imprimer les rectangles comme indiqué dans les exemples. Sinon, la fonction ne doit rien imprimer.

### Fonction attendue

```go
func QuadB(x, y int) {

}
```

### Utilisation

Voici des programmes possibles pour tester votre fonction :

#### Programme #1

```go
package main

import "piscine"

func main() {
	piscine.QuadB(5, 3)
}
```

Et sa sortie :

```plaintext
$ go run .
/***\
*   *
\***/
$
```

#### Programme #2

```go
package main

import "piscine"

func main() {
	piscine.QuadB(5, 1)
}
```

Et sa sortie :

```plaintext
$ go run .
/***\
$
```

#### Programme #3

```go
package main

import "piscine"

func main() {
	piscine.QuadB(1, 1)
}
```

Et sa sortie :

```plaintext
$ go run .
/
$
```

#### Programme #4

```go
package main

import "piscine"

func main() {
	piscine.QuadB(1, 5)
}
```

Et sa sortie :

```plaintext
$ go run .
/
*
*
*
\
$
```

## quadC

### Instructions

Écrivez une fonction `QuadC` qui imprime un rectangle valide de largeur `x` et de hauteur `y`.

La fonction doit dessiner les rectangles comme dans les exemples.

Si `x` et `y` sont des nombres positifs, le programme doit imprimer les rectangles comme indiqué dans les exemples. Sinon, la fonction ne doit rien imprimer.

### Fonction attendue

```go
func QuadC(x, y int) {

}
```

### Utilisation

Voici des programmes possibles pour tester votre fonction :

#### Programme #1

```go
package main

import "piscine"

func main() {
	piscine.QuadC(5, 3)
}
```

Et sa sortie :

```plaintext
$ go run .
ABBBA
B   B
CBBBC
$
```

#### Programme #2

```go
package main

import "piscine"

func main() {
	piscine.QuadC(5, 1)
}
```

Et sa sortie :

```plaintext
$ go run .
ABBBA
$
```

#### Programme #3

```go
package main

import "piscine"

func main() {
	piscine.QuadC(1, 1)
}
```

Et sa sortie :

```plaintext
$ go run .
A
$
```

#### Programme #4

```go
package main

import "piscine"

func main() {
	piscine.QuadC(1, 5)
}
```

Et sa sortie :

```plaintext
$ go run .
A
B
B
B
C
$
```

## quadD

### Instructions

Écrivez une fonction `QuadD` qui imprime un rectangle valide de largeur `x` et de hauteur `y`.

La fonction doit dessiner les rectangles comme dans les exemples.

Si `x` et `y` sont des nombres positifs, le programme doit imprimer les rectangles comme indiqué dans les exemples. Sinon, la fonction ne doit rien imprimer.

### Fonction attendue

```go
func QuadD(x, y int) {

}
```

### Utilisation

Voici des programmes possibles pour tester votre fonction :

#### Programme #1

```go
package main

import "piscine"

func main() {
	piscine.QuadD(5, 3)
}
```

Et sa sortie :

```plaintext
$ go run .
ABBBC
B   B
ABBBC
$
```

#### Programme #2

```go
package main

import "piscine"

func main() {
	piscine.QuadD(5, 1)
}
```

Et sa sortie :

```plaintext
$ go run .
ABBBC
$
```

#### Programme #3

```go
package main

import "piscine"

func main() {
	piscine.QuadD(1, 1)
}
```

Et sa sortie :

```plaintext
$ go run .
A
$
```

#### Programme #4

```go
package main

import "piscine"

func main() {
	piscine.QuadD(1, 5)
}
```

Et sa sortie :

```plaintext
$ go run .
A
B
B
B
A
$
```

## quadE

### Instructions

Écrivez une fonction `QuadE` qui imprime un rectangle valide de largeur `x` et de hauteur `y`.

La fonction doit dessiner les rectangles comme dans les exemples.

Si `x` et `y` sont des nombres positifs, le programme doit imprimer les rectangles comme indiqué dans les exemples. Sinon, la fonction ne doit rien imprimer.

### Fonction attendue

```go
func QuadE(x, y int) {

}
```

### Utilisation

Voici des programmes possibles pour tester votre fonction :

#### Programme #1

```go
package main

import "piscine"

func main() {
	piscine.QuadE(5, 3)
}


```

Et sa sortie :

```plaintext
$ go run .
ABBBC
B   B
CBBBA
$
```

#### Programme #2

```go
package main

import "piscine"

func main() {
	piscine.QuadE(5, 1)
}
```

Et sa sortie :

```plaintext
$ go run .
ABBBC
$
```

#### Programme #3

```go
package main

import "piscine"

func main() {
	piscine.QuadE(1, 1)
}
```

Et sa sortie :

```plaintext
$ go run .
A
$
```

#### Programme #4

```go
package main

import "piscine"

func main() {
	piscine.QuadE(1, 5)
}
```

Et sa sortie :

```plaintext
$ go run .
A
B
B
B
C
$
```