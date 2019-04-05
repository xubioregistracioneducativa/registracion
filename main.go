package main

import (
  "fmt"
  "github.com/xubioregistracioneducativa/registracion/configuracion"
  "log"
  "net/http"
  "os"
)

//import "github.com/jinzhu/gorm"
//import _ "github.com/jinzhu/gorm/dialects/postgres"

//var db gorm.DB
//var registracionPrueba = Registracion{1, "Walter", "Schmidt", "sebastian.taka@hotmail.com", "1167188487", "Ingenieria en Sistemas de Informacion", "q1q1qq", "Marina", "Olivella", "molivella@xubio.com" ,"Analisis de Sistemas", "", "Facultad Regional Buenos Aires", "Universidad tecnologica nacional", estadoAprobadoID}

func main() {

  err := os.Setenv("RECENV", "D")
  if err != nil {
    panic(err)
  }
  configuracion.CargarConfiguracion()


    CrearTablas()
    router := newRouter()

    fmt.Println("Se empieza a escuchar el puerto")
    server := http.ListenAndServe(configuracion.Puerto(), router)

    log.Fatal(server)

}
