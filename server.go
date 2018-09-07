package base

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/VG-Tech-Dojo/treasure2018/mid/NAKKA-K/cook-do/controller"
	"github.com/VG-Tech-Dojo/treasure2018/mid/NAKKA-K/cook-do/db"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/context"
	"github.com/gorilla/csrf"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// Serverはベースアプリケーションのserverを示します
//
// TODO: dbxをstructから分離したほうが複数人数開発だと見通しがよいかもしれない
type Server struct {
	dbx    *sqlx.DB
	router *mux.Router
}

func (s *Server) Close() error {
	return s.dbx.Close()
}

// InitはServerを初期化する
func (s *Server) Init(dbconf, env string) {
	cs, err := db.NewConfigsFromFile(dbconf)
	if err != nil {
		log.Fatalf("cannot open database configuration. exit. %s", err)
	}
	dbx, err := cs.Open(env)
	if err != nil {
		log.Fatalf("db initialization failed: %s", err)
	}
	s.dbx = dbx
	s.router = s.Route()
}

// Newはベースアプリケーションを初期化します
func New() *Server {
	return &Server{}
}

// csrfProtectKey should have 32 byte length.
var csrfProtectKey = []byte("32-byte-long-auth-key")

func (s *Server) Run(addr string) {
	log.Printf("start listening on %s", addr)
	// NOTE: when you serve on TLS, make csrf.Secure(true)
	extendsCorsMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	extendsCorsHeader := handlers.AllowedHeaders([]string{"X-Requested-With", "X-CSRF-Token", "Accept", "Accept-Language", "Content-Language", "Origin", "Content-Type"})
	http.ListenAndServe(addr, handlers.CORS(extendsCorsMethods, extendsCorsHeader)(context.ClearHandler(s.router)))
}

// Routeはベースアプリケーションのroutingを設定します
func (s *Server) Route() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "pong")
	}).Methods("GET")
	router.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"token": csrf.Token(r),
		})
	}).Methods("GET")

	todo := &controller.Todo{DB: s.dbx}

	// TODO ng?
	router.Handle("/api/todos", handler(todo.Get)).Methods("GET")
	router.Handle("/api/todos", handler(todo.Put)).Methods("PUT")
	router.Handle("/api/todos", handler(todo.Post)).Methods("POST")
	router.Handle("/api/todos", handler(todo.Delete)).Methods("DELETE")
	router.Handle("/api/todos/toggle", handler(todo.Toggle)).Methods("PUT")
	router.Handle("/api/todos/all", handler(todo.DeleteAll)).Methods("DELETE")
	router.Handle("/api/todos/search", handler(todo.Search)).Methods("GET")
	router.Handle("/api/recipe/scraping", handler(todo.Scraping)).Methods("POST")
	return router
}