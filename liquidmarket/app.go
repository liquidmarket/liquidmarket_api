package liquidmarket

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, socket, dbname string, local bool) {
	var connectionString string
	if local {
		connectionString = fmt.Sprintf("%s:%s@/%s", user, password, dbname)
	} else {
		connectionString = fmt.Sprintf("%s:%s@cloudsql(%s)/%s", user, password, socket, dbname)
	}
	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run() {
	http.Handle("/", corsWrapper(a.Router))
}

func (a *App) RunLocal(addr string) {
	fmt.Println("running on localhost" + addr + " (hopefully)")
	log.Fatal(http.ListenAndServe(addr, corsWrapper(a.Router)))
}
