package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Pages struct {
	Title        string
	ErrorMessage string
}

type Post struct {
	ID       int
	Title    string
	Body     string
	Comments []Comment
}

type Comment struct {
	ID     int
	PostID int
	Body   string
}

var database *sql.DB

func main() {
	// Open SQLite database
	dbPath := "./forum.db"
	var err error
	database, err = openDatabase(dbPath)
	if err != nil {
		log.Fatalf("Erreur lors de l'ouverture de la base de données: %v", err)
	}

	createAllTables(dbPath)

	// Create posts table if it doesn't exist
	_, err = database.Exec(`CREATE TABLE IF NOT EXISTS posts (
        id INTEGER PRIMARY KEY,
        title TEXT,
        body TEXT
    )`)
	if err != nil {
		log.Fatal(err)
	}

	// Create comments table if it doesn't exist
	_, err = database.Exec(`CREATE TABLE IF NOT EXISTS comments (
        id INTEGER PRIMARY KEY,
        post_id INTEGER,
        body TEXT
    )`)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	router := mux.NewRouter()

	assetsDir := http.Dir("../frontend/assets")
	assetsHandler := http.StripPrefix("/assets/", http.FileServer(assetsDir))
	router.PathPrefix("/assets/").Handler(assetsHandler)
	fs := http.FileServer(http.Dir("frontend"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))
	// Initialize routes
	// router.HandleFunc("/", indexHandler).Methods("GET")
	router.HandleFunc("/", successfulLoginHandler)
	router.HandleFunc("/Login", successfulLoginHandler)
	router.HandleFunc("/signup", successfulSignUp)
	router.HandleFunc("/home", indexHandler)
	router.HandleFunc("/archisubmit", archiSubmitHandler).Methods("POST")
	router.HandleFunc("/infosubmit", infoSubmitHandler)
	router.HandleFunc("/cybersubmit", cyberSubmitHandler)
	router.HandleFunc("/marketsubmit", marketSubmitHandler)
	router.HandleFunc("/techsubmit", techSubmitHandler)
	router.HandleFunc("/creasubmit", creaSubmitHandler)
	router.HandleFunc("/troisDsubmit", troisDSubmitHandler)
	router.HandleFunc("/audiosubmit", audiovisuelSubmitHandler)
	router.HandleFunc("/comment", archiCommentHandler).Methods("POST")
	router.HandleFunc("/infocomment", infoCommentHandler).Methods("POST")
	router.HandleFunc("/cybercomment", cyberCommentHandler).Methods("POST")
	router.HandleFunc("/marketcomment", marketCommentHandler).Methods("POST")
	router.HandleFunc("/creacomment", creaCommentHandler).Methods("POST")
	router.HandleFunc("/troisDcomment", troisDCommentHandler).Methods("POST")
	router.HandleFunc("/audiocomment", audioCommentHandler).Methods("POST")
	router.HandleFunc("/Informatique", InformatiqueHandler).Methods("GET")
	router.HandleFunc("/Cybersecurite", CybersecuriteHandler).Methods("GET")
	router.HandleFunc("/MarketCom", MarketComHandler).Methods("GET")
	router.HandleFunc("/TechAndBuisness", techHandler).Methods("GET")
	router.HandleFunc("/Architecture", ArchitectureHandler).Methods("GET")
	router.HandleFunc("/CreaDesign", CreaDesignHandler).Methods("GET")
	router.HandleFunc("/3DJeuxvideo", TroisDJeuxVideoHandler)
	router.HandleFunc("/Audiovisuel", AudiovisuelHandler)

	// Serve static files
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the server
	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func createAllTables(dbPath string) {
	DB, err := openDatabase(dbPath)
	if err != nil {
		log.Fatalf("%v cant be opened", err)
	}
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS archiposts (id INTEGER PRIMARY KEY, title TEXT, body TEXT)")
	fmt.Println("created table archiposts")

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS archicomments (id INTEGER PRIMARY KEY, post_id INTEGER, body TEXT)")
	fmt.Println("Created table archicomments")

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS infoposts (id INTEGER PRIMARY KEY, title TEXT, body TEXT)")
	fmt.Println("created table infoposts")

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS infocomments (id INTEGER PRIMARY KEY, post_id INTEGER, body TEXT)")
	fmt.Println("Created table infocomments")

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS cyberposts (id INTEGER PRIMARY KEY, title TEXT, body TEXT)")
	fmt.Println("created table cyberposts")

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS cybercomments (id INTEGER PRIMARY KEY, post_id INTEGER, body TEXT)")
	fmt.Println("Created table cybercomments")

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS marketposts (id INTEGER PRIMARY KEY, title TEXT, body TEXT)")
	fmt.Println("created table marketposts")

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS marketcomments (id INTEGER PRIMARY KEY, post_id INTEGER, body TEXT)")
	fmt.Println("Created table marketcomments")

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS techposts (id INTEGER PRIMARY KEY, title TEXT, body TEXT)")
	fmt.Println("created table techposts")

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS techcomments (id INTEGER PRIMARY KEY, post_id INTEGER, body TEXT)")
	fmt.Println("Created table techcomments")

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS creaposts (id INTEGER PRIMARY KEY, title TEXT, body TEXT)")
	fmt.Println("created table creaposts")

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS creacomments (id INTEGER PRIMARY KEY, post_id INTEGER, body TEXT)")
	fmt.Println("Created table creacomments")

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS troisDposts (id INTEGER PRIMARY KEY, title TEXT, body TEXT)")
	fmt.Println("created table troisDposts")

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS troisDcomments (id INTEGER PRIMARY KEY, post_id INTEGER, body TEXT)")
	fmt.Println("Created table troisDcomments")

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS audioposts (id INTEGER PRIMARY KEY, title TEXT, body TEXT)")
	fmt.Println("created table audioposts")

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS audiocomments (id INTEGER PRIMARY KEY, post_id INTEGER, body TEXT)")
	fmt.Println("Created table audiocomments")
}

func openDatabase(dbPath string) (*sql.DB, error) {
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		log.Printf("la base de donees %s n'existe pas, creation...", dbPath)
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	if err := createDatabase(db); err != nil {
		return nil, err
	}

	return sql.Open("sqlite3", dbPath)
}

// Will createa all the necessary tables to stock the informations
func createDatabase(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, email TEXT UNIQUE,password TEXT, username TEXT);")
	if err != nil {
		return err
	}

	return nil
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	loginPage := Pages{
		Title: "Forum",
	}

	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Construct the path to the template file relative to the current working directory
	templatePath := filepath.Join(cwd, "../", "frontend", "pages", "login.html")
	fmt.Println("Template Path:", templatePath) // Debugging

	parsedTemplate, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Println("Error parsing template:", err)
		return
	}

	err = parsedTemplate.Execute(w, loginPage)
	if err != nil {
		log.Println("Error executing template:", err)
		return
	}
}

func successfulLoginHandler(w http.ResponseWriter, r *http.Request) {
	password := r.FormValue("password")
	username := r.FormValue("username")

	log.Printf("Received login request for username: %s\n", username)

	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	correct, err := canConnect(db, username, password)
	if err != nil {
		log.Fatal(err)
	}

	if correct {
		log.Printf("Login successful for username: %s\n", username)
		http.SetCookie(w, &http.Cookie{
			Name:  "username",
			Value: username,
		})
		http.Redirect(w, r, "/home", http.StatusSeeOther) // Redirect to /test if login is successful
		return
	} else {
		log.Printf("Login failed for username: %s\n", username)
		loginPage := Pages{
			Title:        "Forum",
			ErrorMessage: "Incorrect username or password",
		}
		// Render the login page again with an error message
		renderLoginTemplate(w, loginPage)
		return
	}
}

func successfulSignUp(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email")
	password := r.FormValue("password")
	username := r.FormValue("username")

	if email == "" || password == "" || username == "" {
		errorMessage := "Please fill out all required fields."
		pageInfo := Pages{Title: "Sign Up", ErrorMessage: errorMessage}
		renderSignUpPage(w, pageInfo)
		return
	}

	log.Printf("Received sign up request for username: %s\n", username)

	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if the username already exists in the database
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", username).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count > 0 {
		log.Printf("Username already exists: %s\n", username)
		errorMessage := "Username already exists. Please choose a different username."
		pageInfo := Pages{Title: "Sign Up", ErrorMessage: errorMessage}
		renderSignUpPage(w, pageInfo)
		// Render the sign up page again with an error message
		renderSignUpPage(w, pageInfo)
		return
	}

	// Insert the new user into the database
	_, err = db.Exec("INSERT INTO users (email, password, username) VALUES (?, ?, ?)", email, password, username)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Sign up successful for username: %s\n", username)
	log.Println("Redirecting to login page...")
	http.Redirect(w, r, "/Login", http.StatusSeeOther) // Redirect to /login after successful sign up
}

func renderSignUpPage(w http.ResponseWriter, pageInfo Pages) {
	signUpTmpl, err := template.ParseFiles("../frontend/templates/signup.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = signUpTmpl.Execute(w, pageInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func renderLoginTemplate(w http.ResponseWriter, page Pages) {
	pageInfo := Pages{
		Title:        "Sign Up",
		ErrorMessage: "",
	}
	loginTmpl, err := template.ParseFiles("../frontend/templates/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	loginTmpl.Execute(w, pageInfo)

}

func homePage(w http.ResponseWriter, r *http.Request) {
	homePage := Pages{
		Title:        "Home Page",
		ErrorMessage: " ",
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	homeTemplate := filepath.Join(cwd, "../", "frontend", "templates", "home.html")
	fmt.Println("Home path is :", homeTemplate)

	parsedHome, err := template.ParseFiles(homeTemplate)
	if err != nil {
		log.Println("Error parsing home file :", err)
	}

	err = parsedHome.Execute(w, homePage)
	if err != nil {
		log.Println("Error executing home page", err)
		return
	}
}

func canConnect(db *sql.DB, username string, password string) (bool, error) {
	var correct bool
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE username = ? AND password = ?)", username, password).Scan(&correct)
	if err != nil {
		return false, err
	}
	return correct, nil
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Récupérer le nom d'utilisateur à partir du cookie
	cookie, err := r.Cookie("username")
	if err != nil {
		http.Error(w, "User not logged in", http.StatusUnauthorized)
		return
	}
	username := cookie.Value

	// Créer une structure pour le modèle
	data := struct {
		Username string
	}{
		Username: username,
	}

	// Analyser le fichier de modèle
	tmpl, err := template.ParseFiles("../frontend/templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Exécuter le modèle avec les données
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func archiSubmitHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.Form.Get("title")
	body := r.Form.Get("body")

	// Insert post into the database
	stmt, err := database.Prepare("INSERT INTO archiposts (title, body) VALUES (?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/Architecture", http.StatusSeeOther)
}

func infoSubmitHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.Form.Get("title")
	body := r.Form.Get("body")

	// Insert post into the database
	stmt, err := database.Prepare("INSERT INTO infoposts (title, body) VALUES (?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/Informatique", http.StatusSeeOther)
}

func cyberSubmitHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.Form.Get("title")
	body := r.Form.Get("body")

	// Insert post into the database
	stmt, err := database.Prepare("INSERT INTO cyberposts (title, body) VALUES (?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/Cybersecurite", http.StatusSeeOther)
}

func marketSubmitHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.Form.Get("title")
	body := r.Form.Get("body")

	// Insert post into the database
	stmt, err := database.Prepare("INSERT INTO marketposts (title, body) VALUES (?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/MarketCom", http.StatusSeeOther)
}

func techSubmitHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.Form.Get("title")
	body := r.Form.Get("body")

	// Insert post into the database
	stmt, err := database.Prepare("INSERT INTO techposts (title, body) VALUES (?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/TechAndBuisness", http.StatusSeeOther)
}

func creaSubmitHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.Form.Get("title")
	body := r.Form.Get("body")

	// Insert post into the database
	stmt, err := database.Prepare("INSERT INTO creaposts (title, body) VALUES (?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/CreaDesign", http.StatusSeeOther)
}

func troisDSubmitHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.Form.Get("title")
	body := r.Form.Get("body")

	// Insert post into the database
	stmt, err := database.Prepare("INSERT INTO troisDposts (title, body) VALUES (?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/3DJeuxvideo", http.StatusSeeOther)
}

func audiovisuelSubmitHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.Form.Get("title")
	body := r.Form.Get("body")

	// Insert post into the database
	stmt, err := database.Prepare("INSERT INTO audioposts (title, body) VALUES (?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/Audiovisuel", http.StatusSeeOther)
}

func ArchitectureHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve posts from the database
	rows, err := database.Query("SELECT id, title, body FROM archiposts ORDER BY id DESC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Body); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Retrieve comments for each post
		comments, err := getArchiComments(post.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		post.Comments = comments
		posts = append(posts, post)
	}

	// Render posts template
	tmpl, err := template.ParseFiles("../frontend/templates/archi.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, posts)
}

func archiCommentHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	postID, _ := strconv.Atoi(r.Form.Get("post_id"))
	body := r.Form.Get("body")

	// Insert comment into the database
	stmt, err := database.Prepare("INSERT INTO archicomments (post_id, body) VALUES (?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(postID, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "../", http.StatusSeeOther)
}

func infoCommentHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	postID, _ := strconv.Atoi(r.Form.Get("post_id"))
	body := r.Form.Get("body")

	// Insert comment into the database
	stmt, err := database.Prepare("INSERT INTO infocomments (post_id, body) VALUES (?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(postID, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/Informatique", http.StatusSeeOther)
}

func creaCommentHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	postID, _ := strconv.Atoi(r.Form.Get("post_id"))
	body := r.Form.Get("body")

	// Insert comment into the database
	stmt, err := database.Prepare("INSERT INTO creacomments (post_id, body) VALUES (?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(postID, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/CreaDesign", http.StatusSeeOther)
}

func cyberCommentHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	postID, _ := strconv.Atoi(r.Form.Get("post_id"))
	body := r.Form.Get("body")

	// Insert comment into the database
	stmt, err := database.Prepare("INSERT INTO cybercomments (post_id, body) VALUES (?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(postID, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/Cybersecurite", http.StatusSeeOther)
}

func marketCommentHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	postID, _ := strconv.Atoi(r.Form.Get("post_id"))
	body := r.Form.Get("body")

	// Insert comment into the database
	stmt, err := database.Prepare("INSERT INTO marketcomments (post_id, body) VALUES (?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(postID, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/MarketCom", http.StatusSeeOther)
}

func troisDCommentHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	postID, _ := strconv.Atoi(r.Form.Get("post_id"))
	body := r.Form.Get("body")

	// Insert comment into the database
	stmt, err := database.Prepare("INSERT INTO troisDcomments (post_id, body) VALUES (?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(postID, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/3DJeuxvideo", http.StatusSeeOther)
}

func audioCommentHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	postID, _ := strconv.Atoi(r.Form.Get("post_id"))
	body := r.Form.Get("body")

	// Insert comment into the database
	stmt, err := database.Prepare("INSERT INTO audiocomments (post_id, body) VALUES (?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(postID, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/Audiovisuel", http.StatusSeeOther)
}

func getArchiComments(postID int) ([]Comment, error) {
	rows, err := database.Query("SELECT id, post_id, body FROM archicomments WHERE post_id = ?", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.Body); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func getInfoComments(postID int) ([]Comment, error) {
	rows, err := database.Query("SELECT id, post_id, body FROM infocomments WHERE post_id = ?", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.Body); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func getCyberComments(postID int) ([]Comment, error) {
	rows, err := database.Query("SELECT id, post_id, body FROM cybercomments WHERE post_id = ?", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.Body); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func getMarketComments(postID int) ([]Comment, error) {
	rows, err := database.Query("SELECT id, post_id, body FROM marketcomments WHERE post_id = ?", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.Body); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func getTechComments(postID int) ([]Comment, error) {
	rows, err := database.Query("SELECT id, post_id, body FROM techcomments WHERE post_id = ?", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.Body); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func getCreaComments(postID int) ([]Comment, error) {
	rows, err := database.Query("SELECT id, post_id, body FROM creacomments WHERE post_id = ?", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.Body); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func getTroisDComments(postID int) ([]Comment, error) {
	rows, err := database.Query("SELECT id, post_id, body FROM troisDcomments WHERE post_id = ?", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.Body); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func getAudioComments(postID int) ([]Comment, error) {
	rows, err := database.Query("SELECT id, post_id, body FROM audiocomments WHERE post_id = ?", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.Body); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func techHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve posts from the database
	rows, err := database.Query("SELECT id, title, body FROM techposts ORDER BY id DESC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Body); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Retrieve comments for each post
		comments, err := getTechComments(post.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		post.Comments = comments
		posts = append(posts, post)
	}

	// Render posts template
	tmpl, err := template.ParseFiles("../frontend/templates/tech.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, posts)
}

func MarketComHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve posts from the database
	rows, err := database.Query("SELECT id, title, body FROM marketposts ORDER BY id DESC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Body); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Retrieve comments for each post
		comments, err := getMarketComments(post.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		post.Comments = comments
		posts = append(posts, post)
	}

	// Render posts template
	tmpl, err := template.ParseFiles("../frontend/templates/market.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, posts)
}

func CybersecuriteHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve posts from the database
	rows, err := database.Query("SELECT id, title, body FROM cyberposts ORDER BY id DESC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Body); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Retrieve comments for each post
		comments, err := getCyberComments(post.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		post.Comments = comments
		posts = append(posts, post)
	}

	// Render posts template
	tmpl, err := template.ParseFiles("../frontend/templates/cyber.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, posts)
}

func CreaDesignHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve posts from the database
	rows, err := database.Query("SELECT id, title, body FROM creaposts ORDER BY id DESC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Body); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Retrieve comments for each post
		comments, err := getCreaComments(post.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		post.Comments = comments
		posts = append(posts, post)
	}

	// Render posts template
	tmpl, err := template.ParseFiles("../frontend/templates/crea.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, posts)
}

func TroisDJeuxVideoHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve posts from the database
	rows, err := database.Query("SELECT id, title, body FROM troisDposts ORDER BY id DESC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Body); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Retrieve comments for each post
		comments, err := getTroisDComments(post.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		post.Comments = comments
		posts = append(posts, post)
	}

	// Render posts template
	tmpl, err := template.ParseFiles("../frontend/templates/troisd.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, posts)
}

func AudiovisuelHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve posts from the database
	rows, err := database.Query("SELECT id, title, body FROM audioposts ORDER BY id DESC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Body); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Retrieve comments for each post
		comments, err := getAudioComments(post.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		post.Comments = comments
		posts = append(posts, post)
	}

	// Render posts template
	tmpl, err := template.ParseFiles("../frontend/templates/audio.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, posts)
}

func InformatiqueHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve posts from the database
	rows, err := database.Query("SELECT id, title, body FROM infoposts ORDER BY id DESC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Body); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Retrieve comments for each post
		comments, err := getInfoComments(post.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		post.Comments = comments
		posts = append(posts, post)
	}

	// Render posts template
	tmpl, err := template.ParseFiles("../frontend/templates/info.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, posts)
}
