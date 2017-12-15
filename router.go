package liquidmarket

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", handler).Methods("GET")
	a.Router.HandleFunc("/users", a.getUsers).Methods("GET")
	a.Router.HandleFunc("/accounts", a.getAccounts).Methods("GET")
	a.Router.HandleFunc("/accounts", a.createAccount).Methods("POST")
	a.Router.HandleFunc("/accounts/{uuid}", a.getAccount).Methods("GET")
	a.Router.HandleFunc("/accounts/{uuid}", a.updateAccount).Methods("PUT")
	a.Router.HandleFunc("/accounts/{uuid}", a.deleteAccount).Methods("DELETE")
	a.Router.HandleFunc("/prices", a.getPrices).Methods("GET")
}