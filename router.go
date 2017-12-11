package liquidmarket

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", handler)
	a.Router.HandleFunc("/users", a.getUsers)
}