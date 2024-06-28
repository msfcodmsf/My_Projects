package datahandlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
)

var DB *sql.DB

type Session struct {
	ID     string
	UserID int
	Expiry time.Time
}

// Veritabanına bağlantı açar.
func SetDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./database/forum.db")
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
}

// HTTP isteğinden (r) oturum çerezini alarak oturum bilgilerini döndürür.
func GetSession(r *http.Request) (*Session, error) {
	if DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			return nil, nil // Çerez bulunamadıysa, oturum yok olarak döndür
		}
		return nil, err // Başka bir hata varsa hata döndür
	}

	sessionToken := cookie.Value

	var session Session
	err = DB.QueryRow("SELECT id, user_id, expiry FROM sessions WHERE id = ?", sessionToken).Scan(&session.ID, &session.UserID, &session.Expiry)
	if err != nil {
		return nil, err
	}

	if session.Expiry.Before(time.Now()) {
		return nil, fmt.Errorf("session expired")
	}

	// Oturum açıldığında, kullanıcının diğer oturumlarını kapat
	_, err = DB.Exec("DELETE FROM sessions WHERE user_id = ? AND id <> ?", session.UserID, sessionToken)
	if err != nil {
		return nil, err
	}

	// Oturum süresini her kontrol ettiğimizde uzatalım
	newExpiry := time.Now().Add(10 * time.Minute)
	_, err = DB.Exec("UPDATE sessions SET expiry = ? WHERE id = ?", newExpiry, sessionToken)
	if err != nil {
		return nil, err
	}
	session.Expiry = newExpiry

	return &session, nil
}

// Gerekli veritabanı tablolarını oluşturur.
func CreateTables() {
	SessionTables(DB)
	PostTables(DB)
	UsersTables(DB)
	VoteTables(DB)
	CommentTables(DB)
}

func SessionTables(db *sql.DB) {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS sessions (
            id TEXT PRIMARY KEY,
            user_id INTEGER,
            expiry TIMESTAMP,
            FOREIGN KEY(user_id) REFERENCES users(id)
        );`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal("Query failed: ", err)
		}
	}
}

func PostTables(db *sql.DB) {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS posts (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER,
            title TEXT,
            content TEXT,
            categories TEXT,
            created_at TIMESTAMP
        );`,
	}
	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal("Query failed: ", err)
		}
	}
}

func UsersTables(db *sql.DB) {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            email TEXT UNIQUE NOT NULL,
            username TEXT UNIQUE NOT NULL,
            password TEXT NOT NULL
        );`,
	}
	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal("Query failed: ", err)
		}
	}
}

// Like Ve Dislike tablolarını oluştur
func VoteTables(db *sql.DB) { // Sayısını artırır
	queries := []string{
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

func CommentTables(db *sql.DB) {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS comments (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            post_id INTEGER,
            user_id INTEGER,
            content TEXT,
            created_at TIMESTAMP
        );`,
	}
	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal("Query failed: ", err)
		}
	}
}
