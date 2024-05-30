package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

type User struct {
	ID       int
	Email    string
	Username string
	Password string
}

type Post struct {
	ID                 int
	UserID             int
	Title              string
	Content            string
	Categories         []string
	CreatedAt          time.Time
	CreatedAtFormatted string
	LikeCount          int
	DislikeCount       int
	Username           string
}

type Comment struct {
	ID                 int
	PostID             int
	UserID             int
	Content            string
	CreatedAt          time.Time
	CreatedAtFormatted string
	LikeCount          int
	DislikeCount       int
	Username           string // Kullanıcı adı
}

type Like struct {
	ID        int
	UserID    int
	PostID    sql.NullInt64
	CommentID sql.NullInt64
}

type Session struct {
	ID     string
	UserID int
	Expiry time.Time
}

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close()

	createTables()

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/createPost", createPostHandler)
	http.HandleFunc("/createComment", createCommentHandler)
	http.HandleFunc("/likeHandler", likeHandler)
	http.HandleFunc("/DislikeHandler", dislikeHandler)
	http.HandleFunc("/vote", voteHandler)
	http.HandleFunc("/filter", filterHandler)
	http.HandleFunc("/viewPost", viewPostHandler)

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

func createTables() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            email TEXT UNIQUE,
            username TEXT UNIQUE,
            password TEXT
        );`,
		`CREATE TABLE IF NOT EXISTS posts (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER,
            title TEXT,
            content TEXT,
            categories TEXT,
            created_at TIMESTAMP
        );`,
		`CREATE TABLE IF NOT EXISTS comments (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            post_id INTEGER,
            user_id INTEGER,
            content TEXT,
            created_at TIMESTAMP
        );`,
		`CREATE TABLE IF NOT EXISTS likes (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER,
            post_id INTEGER,
            comment_id INTEGER
        );`,
		`CREATE TABLE IF NOT EXISTS Dislikes (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER,
            post_id INTEGER,
            comment_id INTEGER
        );`,
		`CREATE TABLE IF NOT EXISTS sessions (
            id TEXT PRIMARY KEY,
            user_id INTEGER,
            expiry TIMESTAMP
        );`,
		`CREATE TABLE IF NOT EXISTS votes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			post_id INTEGER,
			comment_id INTEGER,
			vote_type INTEGER CHECK(vote_type IN (1, -1))
		);`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal("Query failed: ", err)
		}
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := getSession(r)

	query := `SELECT posts.id, posts.user_id, posts.title, posts.content, posts.created_at, users.username,
				COALESCE(SUM(CASE WHEN votes.vote_type = 1 THEN 1 ELSE 0 END), 0) AS like_count,
				COALESCE(SUM(CASE WHEN votes.vote_type = -1 THEN 1 ELSE 0 END), 0) AS dislike_count
				FROM posts
				JOIN users ON posts.user_id = users.id
				LEFT JOIN votes ON votes.post_id = posts.id
				GROUP BY posts.id
				ORDER BY posts.created_at DESC`

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		var username string
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt, &username, &post.LikeCount, &post.DislikeCount); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		post.CreatedAtFormatted = post.CreatedAt.Format("2006-01-02 15:04")
		post.Username = username
		posts = append(posts, post)
	}

	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Posts    []Post
		LoggedIn bool
	}{
		Posts:    posts,
		LoggedIn: session != nil,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		username := r.FormValue("username")
		password := r.FormValue("password")

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("INSERT INTO users (email, username, password) VALUES (?, ?, ?)", email, username, hashedPassword)
		if err != nil {
			http.Error(w, "Email or username already taken", http.StatusBadRequest)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	tmpl, _ := template.ParseFiles("templates/register.html")
	tmpl.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		var user User
		err := db.QueryRow("SELECT id, password FROM users WHERE email = ?", email).Scan(&user.ID, &user.Password)
		if err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		sessionID := uuid.NewString()
		expiry := time.Now().Add(24 * time.Hour)

		_, err = db.Exec("INSERT INTO sessions (id, user_id, expiry) VALUES (?, ?, ?)", sessionID, user.ID, expiry)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "session_id",
			Value:   sessionID,
			Expires: expiry,
		})

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmpl, _ := template.ParseFiles("templates/login.html")
	tmpl.Execute(w, nil)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	_, err = db.Exec("DELETE FROM sessions WHERE id = ?", cookie.Value)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		MaxAge: -1,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	session, err := getSession(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		content := r.FormValue("content")
		categories := r.Form["categories"]

		categoriesJSON, err := json.Marshal(categories)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("INSERT INTO posts (user_id, title, content, categories, created_at) VALUES (?, ?, ?, ?, ?)",
			session.UserID, title, content, categoriesJSON, time.Now())
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmpl, _ := template.ParseFiles("templates/createPost.html")
	tmpl.Execute(w, nil)
}

func createCommentHandler(w http.ResponseWriter, r *http.Request) {
	session, err := getSession(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		postID := r.FormValue("post_id")
		content := r.FormValue("content")

		_, err := db.Exec("INSERT INTO comments (post_id, user_id, content, created_at) VALUES (?, ?, ?, ?)",
			postID, session.UserID, content, time.Now())
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/viewPost?id=%s", postID), http.StatusSeeOther)
		return
	}
}

func likeHandler(w http.ResponseWriter, r *http.Request) {
	session, err := getSession(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	postID := r.FormValue("post_id")
	commentID := r.FormValue("comment_id")

	if postID != "" {
		_, err = db.Exec("INSERT INTO likes (user_id, post_id) VALUES (?, ?)", session.UserID, postID)
	} else if commentID != "" {
		_, err = db.Exec("INSERT INTO likes (user_id, comment_id) VALUES (?, ?)", session.UserID, commentID)
	}

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
}

func dislikeHandler(w http.ResponseWriter, r *http.Request) {
	session, err := getSession(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	postID := r.FormValue("post_id")
	commentID := r.FormValue("comment_id")

	if postID != "" {
		_, err = db.Exec("INSERT INTO Dislikes (user_id, post_id) VALUES (?, ?)", session.UserID, postID)
	} else if commentID != "" {
		_, err = db.Exec("INSERT INTO Dislikes (user_id, comment_id) VALUES (?, ?)", session.UserID, commentID)
	}

	if err != nil {
		log.Println("Error disliking:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
}

func voteHandler(w http.ResponseWriter, r *http.Request) {
	session, err := getSession(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	postID := r.FormValue("post_id")
	commentID := r.FormValue("comment_id")
	voteTypeStr := r.FormValue("vote_type")

	// Convert voteType from string to integer
	voteType, err := strconv.Atoi(voteTypeStr)
	if err != nil || (voteType != 1 && voteType != -1) {
		http.Error(w, "Invalid vote type", http.StatusBadRequest)
		return
	}

	var existingVoteType sql.NullInt64
	var query string

	if postID != "" {
		query = "SELECT vote_type FROM votes WHERE user_id = ? AND post_id = ?"
		err = db.QueryRow(query, session.UserID, postID).Scan(&existingVoteType)
	} else if commentID != "" {
		query = "SELECT vote_type FROM votes WHERE user_id = ? AND comment_id = ?"
		err = db.QueryRow(query, session.UserID, commentID).Scan(&existingVoteType)
	}

	if err != nil && err != sql.ErrNoRows {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if existingVoteType.Valid {
		if existingVoteType.Int64 == int64(voteType) {
			if postID != "" {
				query = "DELETE FROM votes WHERE user_id = ? AND post_id = ?"
				_, err = db.Exec(query, session.UserID, postID)
			} else if commentID != "" {
				query = "DELETE FROM votes WHERE user_id = ? AND comment_id = ?"
				_, err = db.Exec(query, session.UserID, commentID)
			}
		} else {
			if postID != "" {
				query = "UPDATE votes SET vote_type = ? WHERE user_id = ? AND post_id = ?"
				_, err = db.Exec(query, voteType, session.UserID, postID)
			} else if commentID != "" {
				query = "UPDATE votes SET vote_type = ? WHERE user_id = ? AND comment_id = ?"
				_, err = db.Exec(query, voteType, session.UserID, commentID)
			}
		}
	} else {
		if postID != "" {
			query = "INSERT INTO votes (user_id, post_id, vote_type) VALUES (?, ?, ?)"
			_, err = db.Exec(query, session.UserID, postID, voteType)
		} else if commentID != "" {
			query = "INSERT INTO votes (user_id, comment_id, vote_type) VALUES (?, ?, ?)"
			_, err = db.Exec(query, session.UserID, commentID, voteType)
		}
	}

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
}

func filterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	category := r.URL.Query().Get("category")
	query := "SELECT id, user_id, title, content, categories, created_at FROM posts"
	args := []interface{}{}

	if category != "" {
		query += " WHERE categories LIKE ?"
		args = append(args, "%"+category+"%")
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		var categories string
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &categories, &post.CreatedAt); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		json.Unmarshal([]byte(categories), &post.Categories)
		posts = append(posts, post)
	}

	tmpl, _ := template.ParseFiles("templates/filter.html")
	tmpl.Execute(w, posts)
}

func viewPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	postID := r.URL.Query().Get("id")
	if postID == "" {
		http.Error(w, "Post ID required", http.StatusBadRequest)
		return
	}

	var post Post
	var categories string
	err := db.QueryRow(`SELECT posts.id, posts.user_id, posts.title, posts.content, posts.created_at, users.username,
						COALESCE(SUM(CASE WHEN votes.vote_type = 1 THEN 1 ELSE 0 END), 0) AS like_count,
						COALESCE(SUM(CASE WHEN votes.vote_type = -1 THEN 1 ELSE 0 END), 0) AS dislike_count
						FROM posts
						JOIN users ON posts.user_id = users.id
						LEFT JOIN votes ON votes.post_id = posts.id
						WHERE posts.id = ?
						GROUP BY posts.id`, postID).Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt, &post.Username, &post.LikeCount, &post.DislikeCount)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	json.Unmarshal([]byte(categories), &post.Categories)
	post.CreatedAtFormatted = post.CreatedAt.Format("2006-01-02 15:04")

	rows, err := db.Query(`SELECT comments.id, comments.post_id, comments.user_id, comments.content, comments.created_at, users.username,
							COALESCE(SUM(CASE WHEN votes.vote_type = 1 THEN 1 ELSE 0 END), 0) AS like_count,
							COALESCE(SUM(CASE WHEN votes.vote_type = -1 THEN 1 ELSE 0 END), 0) AS dislike_count
							FROM comments
							JOIN users ON comments.user_id = users.id
							LEFT JOIN votes ON votes.comment_id = comments.id
							WHERE comments.post_id = ?
							GROUP BY comments.id`, postID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Content, &comment.CreatedAt, &comment.Username, &comment.LikeCount, &comment.DislikeCount); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		comment.CreatedAtFormatted = comment.CreatedAt.Format("2006-01-02 15:04")
		comments = append(comments, comment)
	}

	data := struct {
		Post     Post
		Comments []Comment
	}{
		Post:     post,
		Comments: comments,
	}

	tmpl, err := template.ParseFiles("templates/viewPost.html")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func getSession(r *http.Request) (*Session, error) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return nil, err
	}

	var session Session
	err = db.QueryRow("SELECT id, user_id, expiry FROM sessions WHERE id = ?", cookie.Value).Scan(&session.ID, &session.UserID, &session.Expiry)
	if err != nil {
		return nil, err
	}

	if session.Expiry.Before(time.Now()) {
		_, _ = db.Exec("DELETE FROM sessions WHERE id = ?", session.ID)
		return nil, fmt.Errorf("session expired")
	}

	return &session, nil
}
