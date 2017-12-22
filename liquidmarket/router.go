package liquidmarket

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", a.getAccountsOrCreateUser).Methods("GET")
	a.Router.HandleFunc("/users", a.updateUser).Methods("PUT")
	a.Router.HandleFunc("/users/{google_id}", a.getUser).Methods("GET")
	a.Router.HandleFunc("/accounts", a.getAccounts).Methods("GET")
	a.Router.HandleFunc("/accounts", a.createAccount).Methods("POST")
	a.Router.HandleFunc("/accounts/{uuid}", a.getAccount).Methods("GET")
	a.Router.HandleFunc("/accounts/{uuid}", a.updateAccount).Methods("PUT")
	a.Router.HandleFunc("/accounts/{uuid}", a.deleteAccount).Methods("DELETE")
	a.Router.HandleFunc("/prices", a.getPrices).Methods("GET")
	a.Router.HandleFunc("/markets", a.getMarkets).Methods("GET")
	a.Router.HandleFunc("/accounts/{uuid}/shareholdings", a.getShareholdings).Methods("GET")
	a.Router.HandleFunc("/trades", a.getTrades).Methods("GET")
	a.Router.HandleFunc("/addresses", a.getAddresses).Methods("GET")
	a.Router.HandleFunc("/addresses", a.createAddress).Methods("POST")
	a.Router.HandleFunc("/addresses/{uuid}", a.updateAddress).Methods("PUT")
	a.Router.HandleFunc("/addresses/{uuid}", a.deleteAddress).Methods("DELETE")
	a.Router.HandleFunc("/offers", a.submitOffer).Methods("POST")
}
