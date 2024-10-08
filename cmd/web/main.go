package main

import (
	"database/sql"
	"flag"
	"forum/pkg/models/mysql"
	"html/template"
	"log"
	"net/http"
	"os"

	"time"

	"github.com/golangcollege/sessions"

	_ "github.com/go-sql-driver/mysql"
)


type contextKey string

var contextKeyUser = contextKey("user")
type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	session       *sessions.Session
	posts         *mysql.PostModel
	users         *mysql.UserModel
	templateCache map[string]*template.Template
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network adress")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data storage")
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()
	templateCache, err := newTemplateCache("./ui/html")
	if err != nil {
		errorLog.Fatal(err)
	}
	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour
	
	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		session:       session,
		posts:         &mysql.PostModel{DB: db},
		templateCache: templateCache,
		users:&mysql.UserModel{DB:db},
	}

	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on http://localhost%s", *addr)

	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
