package main

import (
  "fmt"
  "github.com/gorilla/handlers"
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


  //#####################LOGGER#####################
  /*logger, err := os.OpenFile("errors.log", os.O_CREATE | os.O_RDWR | os.O_APPEND, 0666)
  if err != nil {
    log.Println("Error opening/creating database log file")
  }


  defer logger.Close()
  */
  log.SetOutput(os.Stdout)
  //###################CargarConfiguracion###########
  configuracion.CargarConfiguracion()

  //###################CrearTablas###################
  GetDBHelper().crearTablas()
  //#############ServicioDeEnviarMailsColgados#######
  go EnviarMailsNoEnviados()
  //###################LevantarServer################

  router := newRouter()
 // headers:= handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Access-Control-Allow-Origin", "application/x-www-form-urlencoded", "application/json", "application/x-www-form-urlencoded;charset=UTF-8"})
  //methods :=  handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
  //origins := handlers.AllowedOrigins([]string{"*"})
  fmt.Println("Se empieza a escuchar el puerto")
  server := http.ListenAndServeTLS(configuracion.Puerto(),"server.crt", "server.key", handlers.CORS()(router))

  log.Fatal(server)

}
