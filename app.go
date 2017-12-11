package liquidmarket

import (
	"net/http"
	// "log"
    "database/sql"
	"github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
    // "fmt"
)

type App struct {
    Router *mux.Router
    DB     *sql.DB
}

func (a *App) Initialize(user, password, socket, dbname string) {
    // connectionString := fmt.Sprintf("%s:%s@cloudsql(%s)/liquidmarket", user, password, dbname)
	// var err error
    // a.DB, err = sql.Open("mysql", connectionString)
    // if err != nil {
    //     log.Fatal(err)
    // }
    a.Router = mux.NewRouter()
    a.initializeRoutes()
}

func (a *App) Run() {
	http.Handle("/", corsWrapper(a.Router))
}
