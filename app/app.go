package app

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/donaderoyan/gomovie/app/controller"
    "github.com/donaderoyan/gomovie/app/model"
    "github.com/donaderoyan/gomovie/config"
)

type App struct {
  Router *mux.Router
  DB     *gorm.DB
}

func (a *App) Initialize(config *config.Config) {
  dbURI := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
    config.DB.Host,
    config.DB.Port,
    config.DB.User,
    config.DB.Dbname,
    config.DB.Password,
    )
  fmt.Printf("%+v\n", dbURI)
  db, err := gorm.Open(config.DB.Dialect, dbURI)
  if err != nil {
    log.Fatal("Could not connect database %s", err.Error())
  }

  a.DB = model.Migration(db)
  a.Router = mux.NewRouter()
  a.setRouters()
}

func (a *App) setRouters() {
  a.Get("/user", a.handleRequest(controller.GetAllUser))
  a.Post("/user", a.handleRequest(controller.CreateUser))
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

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
  return func (w http.ResponseWriter, r *http.Request) {
    handler(a.DB, w, r)
  }
}
