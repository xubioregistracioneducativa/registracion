package main

//import "fmt"
//import "github.com/jinzhu/gorm"
//import _ "github.com/jinzhu/gorm/dialects/postgres"

//var db gorm.DB


func main() {
  var registracion = Registracion{0, "Walter", "Schmidt", "sebastian.taka@hotmail.com", "1167188487", "Ingenieria en Sistemas de Informacion", "q1q1qq", "Marina", "Olivella", "molivella@xubio.com" ,"Analisis de Sistemas", "", "Facultad Regional Buenos Aires", "Universidad tecnologica nacional", estadoInicioRegistracionID}
  for i := 0; i < 4;i++{
    
    nuevoEstado(registracion.estado).ingresarNuevosDatos(&registracion)
    nuevoEstado(registracion.estado).aceptarPorCS(&registracion)
    nuevoEstado(registracion.estado).rechazarPorCS(&registracion)
    nuevoEstado(registracion.estado).aceptarPorProfesor(&registracion)
  }
  //fmt.Println(err)
	//defer db.Close()

}
