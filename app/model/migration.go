package model

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

func Migration(db *gorm.DB) *gorm.DB {
  db.AutoMigrate(
    &User{},
    &Email{},
    &Address{},
    &Language{},
    &CreditCard{},
  )
  return db
}
