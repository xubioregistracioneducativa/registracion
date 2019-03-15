package main

import "fmt"
import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/postgres"

var db gorm.DB

type Registracion struct {
  gorm.Model
  DatosRegistracion  string 
  tipo  Tipo `gorm:"polymorphic:Owner;"`
}


type Tipo int

const (
  TipoA Tipo = 0
  TipoB Tipo = 2
)

func main() {

	db, err := gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")
	fmt.Println(err)
	defer db.Close()

}
