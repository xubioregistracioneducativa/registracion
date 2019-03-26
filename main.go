package main

//import "github.com/jinzhu/gorm"
//import _ "github.com/jinzhu/gorm/dialects/postgres"

//var db gorm.DB
//var registracionPrueba = Registracion{1, "Walter", "Schmidt", "sebastian.taka@hotmail.com", "1167188487", "Ingenieria en Sistemas de Informacion", "q1q1qq", "Marina", "Olivella", "molivella@xubio.com" ,"Analisis de Sistemas", "", "Facultad Regional Buenos Aires", "Universidad tecnologica nacional", estadoAprobadoID}

func main() {
  //CrearTablas()
  var registracion *Registracion
  registrarTenant(registracion)
  //router := newRouter()
  
  //fmt.Println("Se empieza a escuchar el puerto")
  //server := http.ListenAndServe(":8081", router)

  //log.Fatal(server)

  //fmt.Println(err)
	//defer db.Close()

}
