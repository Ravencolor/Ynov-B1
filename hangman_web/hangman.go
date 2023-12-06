package main

import (
	"bufio"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	motADeviner     string
	lettresDevinees = make(map[rune]bool)
	tentativesMax   = 10
	tentatives      = 0
	partieGagnee    = false
	listeMots       [][]string
	listeCourante   int
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	changer_difficulte()
	http.HandleFunc("/", gestionnaire)
	http.HandleFunc("/start", gestionnaire)
	http.HandleFunc("/game", gestionnaire)
	http.HandleFunc("/restart", restartHandler) // Added restart handler
	http.ListenAndServe(":8080", nil)
}

func gestionnaire(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		affichage_index(w)
	} else if r.Method == http.MethodPost && r.URL.Path == "/start" {
		listeIndex := r.FormValue("liste")
		if index, err := strconv.Atoi(listeIndex); err == nil && index >= 0 && index < len(listeMots) {
			listeCourante = index
			resetGame() // Reset the game state on starting a new game
			http.Redirect(w, r, "/game", http.StatusSeeOther)
			return
		}
		http.Error(w, "Liste invalide", http.StatusBadRequest)
	} else if r.URL.Path == "/game" {
		if r.Method == http.MethodPost {
			proposition_joueur(w, r)
		} else {
			affichage_pendu(w)
		}
	} else {
		http.NotFound(w, r)
	}
}

func restartHandler(w http.ResponseWriter, r *http.Request) {
	resetGame() // Reset the game state
	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

func resetGame() {
	lettresDevinees = make(map[rune]bool)
	tentatives = 0
	partieGagnee = false
	motADeviner = mot_du_pendu()
}

func affichage_index(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func proposition_joueur(w http.ResponseWriter, r *http.Request) {
	proposition := []rune(r.FormValue("proposition"))[0]
	lettresDevinees[proposition] = true

	if !proposition_correcte(proposition) {
		tentatives++
	}

	if mot_deviné(motADeviner, lettresDevinees) {
		partieGagnee = true
	}

	affichage_pendu(w)
}

func proposition_correcte(proposition rune) bool {
	for _, char := range motADeviner {
		if proposition == char {
			return true
		}
	}
	return false
}

func affichage_pendu(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("pendu.html")
	if err != nil {
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	motAffiche := affichage_mot()
	lettresDevineesStr := obtenir_lettres_trouvees()
	imageErreurURL := getImageURL(tentatives)

	tmpl.Execute(w, map[string]interface{}{
		"MotAffiche":          motAffiche,
		"LettresDevinees":     lettresDevineesStr,
		"TentativesRestantes": tentativesMax - tentatives,
		"PartiePerdue":        tentatives >= tentativesMax,
		"PartieGagnee":        partieGagnee,
		"MotADeviner":         motADeviner,
		"ImageErreurURL":      imageErreurURL,
	})
}

func mot_deviné(mot string, lettres map[rune]bool) bool {
	for _, lettre := range mot {
		if _, trouvee := lettres[lettre]; !trouvee {
			return false
		}
	}
	return true
}

func affichage_mot() string {
	affichage := ""
	for _, char := range motADeviner {
		if lettresDevinees[char] {
			affichage += string(char)
		} else {
			affichage += "_ "
		}
	}
	return affichage
}

func obtenir_lettres_trouvees() string {
	lettres := ""
	for lettre := range lettresDevinees {
		lettres += string(lettre) + " "
	}
	return lettres
}

func mot_du_pendu() string {
	return MotAleatoire(listeMots[listeCourante])
}

func changer_difficulte() {
	listeMots = make([][]string, 3)

	charger_liste_mot("dico/nom_personnages.txt", 0)
	charger_liste_mot("dico/nom_planetes.txt", 1)
	charger_liste_mot("dico/nom_races.txt", 2)
}

func charger_liste_mot(nomFichier string, indice int) {
	f, err := os.Open(nomFichier)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		listeMots[indice] = append(listeMots[indice], scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func MotAleatoire(mots []string) string {
	rand.Seed(time.Now().Unix())
	return mots[rand.Intn(len(mots))]
}

func getImageURL(tentatives int) string {
	imagePaths := []string{
		"https://sweezy-cursors.com/wp-content/uploads/cursor/star-wars-asoka-tano-lightsaber/star-wars-asoka-tano-lightsaber-custom-cursor.png",
		"https://sweezy-cursors.com/wp-content/uploads/cursor/star-wars-aayla-secura-lightsaber/star-wars-aayla-secura-lightsaber-custom-cursor.png",
		"https://sweezy-cursors.com/wp-content/uploads/cursor/star-wars-princess-leia-jabba/star-wars-princess-leia-jabba-custom-cursor.png",
		"https://sweezy-cursors.com/wp-content/uploads/cursor/star-wars-jango-fett-westar-34-blaster/star-wars-jango-fett-westar-34-custom-cursor.png",
		"https://sweezy-cursors.com/wp-content/uploads/cursor/star-wars-bb-8/star-wars-bb-8-custom-cursor.png",
		"https://sweezy-cursors.com/wp-content/uploads/cursor/star-wars-maz-kanata-dh-17-blaster-pistol/star-wars-masz-kanata-dh-17-blaster-pistol-custom-cursor.png",
		"https://cdn.custom-cursor.com/packs/1881/pack2508.png",
		"https://sweezy-cursors.com/wp-content/uploads/cursor/star-wars-general-grievous-lightsabers/star-wars-general-grievous-lightsabers-custom-cursor.png",
		"https://sweezy-cursors.com/wp-content/uploads/cursor/star-wars-darth-vader-lightsaber/star-wars-darth-vader-lightsaber-custom-cursor.png",
		"https://sweezy-cursors.com/wp-content/uploads/cursor/star-wars-luke-skywalker-lightsaber/star-wars-luke-skywalker-lightsaber-custom-cursor.png",
	}

	index := tentatives - 1
	if index < 0 || index >= len(imagePaths) {
		index = len(imagePaths) - 1
	}

	return imagePaths[index]
}
