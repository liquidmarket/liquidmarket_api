package liquidmarket

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", handler)
}