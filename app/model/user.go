package model

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
  gorm.Model
  FirstName     string `gorm:"size:255"`
  LastName      string `gorm:"size:255"`
  Emails        []Email

  BillingAddress      Address         //one-to-one relationship (belongs to - use BillingAddressID as foreign Key)
  BillingAddressID    int

  ShippingAddress     Address         // one-to-one relationship (belongs to - use ShippingAddressID as foreign key)
  ShippingAddressID   int

  Languages           []Language  `gorm:"many2many:user_languages;"`
}

type Email struct {
  ID      int
  UserID  int     `gorm:"index"`
  Email   string  `gorm:"type:varchar(100);unique_index"`
  Subscribed  bool
}

type Address struct {
  ID        int
  Address1  string          `gorm:"not null;unique"`
  Address2  string          `gorm:"type:varchar(100);unique"`
  Post      string  `gorm:"not null"`
}

type  Language struct {
  ID    int
  Name  string  `gorm:"index:idx_name_code"`  // Create index with name, and will create combined index if find other fields defined same name
  Code  string  `gorm:"index:idx_name_code"`  // `unique_index` also works
}

type CreditCard struct {
  gorm.Model
  UserID  uint
  Number  string
}
