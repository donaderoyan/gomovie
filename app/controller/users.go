package controller

import (
  "encoding/json"
  "net/http"
  "fmt"

  // "github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
  "github.com/donaderoyan/gomovie/app/model"
)

func GetAllUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
  users := []model.User{}
  db.Debug().Find(&users)
  respondJSON(w, http.StatusOK, users)
}


func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
  user  := model.User{}

  decoder := json.NewDecoder(r.Body)
  if err := decoder.Decode(&user); err != nil {
    respondError(w, http.StatusBadRequest, err.Error())
    return
  }
  defer r.Body.Close()

  fmt.Printf("%+v\n", &user)

  if err := db.Debug().Save(&user).Error; err != nil {
    respondError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondJSON(w, http.StatusCreated, user)
}
