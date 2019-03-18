package main

import (
  "fmt"
  "log"
  "net/http"
)
//import "github.com/jinzhu/gorm"
//import _ "github.com/jinzhu/gorm/dialects/postgres"

//var db gorm.DB
var registracionPrueba = Registracion{0, "Walter", "Schmidt", "sebastian.taka@hotmail.com", "1167188487", "Ingenieria en Sistemas de Informacion", "q1q1qq", "Marina", "Olivella", "molivella@xubio.com" ,"Analisis de Sistemas", "", "Facultad Regional Buenos Aires", "Universidad tecnologica nacional", estadoPendienteAprobacionID}

func main() {
  
  router := newRouter()
  
  fmt.Println("Se empieza a escuchar el puerto")
  server := http.ListenAndServe(":8080", router)

  log.Fatal(server)

  //fmt.Println(err)
	//defer db.Close()

}
