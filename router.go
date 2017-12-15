package liquidmarket

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", a.getAccounts).Methods("GET")
	a.Router.HandleFunc("/users", a.updateUser).Methods("PUT")
	a.Router.HandleFunc("/accounts", a.getAccounts).Methods("GET")
	a.Router.HandleFunc("/accounts", a.createAccount).Methods("POST")
	a.Router.HandleFunc("/accounts/{uuid}", a.getAccount).Methods("GET")
	a.Router.HandleFunc("/accounts/{uuid}", a.updateAccount).Methods("PUT")
	a.Router.HandleFunc("/accounts/{uuid}", a.deleteAccount).Methods("DELETE")
	a.Router.HandleFunc("/prices", a.getPrices).Methods("GET")
	a.Router.HandleFunc("/accounts/{uuid}/shareholdings", a.getShareholdings).Methods("GET")
}