package app

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/jinzhu/gorm"
    "github.com/donaderoyan/gomovie/app/model"
    "github.com/donaderoyan/gomovie/conf/configuration"
)

type App struct {
  Router *mux.Router
  DB     *gorm.DB
}

func (a *App) Initialize(config *config.Config) {
  dbURI := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
    config.DB.Host,
    config.DB.Port,
    config.DB.User,
    config.DB.Dbname,
    config.DB.Password
    )

  db, err := gorm.Open(config.DB.Dialect, dbURI)
  if err != nil {
    log.Fatal("Could not connect database")
  }

  a.DB = model.Migration(db)
  a.Router = mux.NewRouter()
  a.setRouters()
}

func (a *App) setRouters() {
  a.Get("/film", a.handleRequest(handler.GetAllfilm))
  a.Post("/film", a.handleRequest(handler.CreateFilm))
}


// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}


//Run app on it's router
func (a *App) Run(host string) {
  log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandleFunc {
  return func (w http.ResponseWriter, r *http.Request) {
    handler(a.DB, w, r)
  }
}
