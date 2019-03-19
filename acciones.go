package main

import (
	"fmt"
    "net/http"
    "log"
    "encoding/json"
)

func responderRegistracion(writer http.ResponseWriter, status int, results Registracion){

	writer.Header().Set("Content-Type", "application-json")
	writer.WriteHeader(status)

	json.NewEncoder(writer).Encode(results)
}


func NuevaRegistracion(writer http.ResponseWriter, request *http.Request){
 

	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields ()
	var datosRegistracion Registracion

	//&nombre_var para decirle que es la var que no tiene datos y va a tener que rellenar
	var err = decoder.Decode(&datosRegistracion)

	if(err != nil){
		panic(err)
		writer.Header().Set("Content-Type", "application-json")
		writer.WriteHeader(400)
	}

	//Para cerrar la lectura de algo
	defer request.Body.Close()

	log.Println(datosRegistracion)

	fmt.Println(datosRegistracion.estado)

	datosRegistracion.estado = estadoInicioRegistracionID

    nuevoEstado(registracionPrueba.estado).ingresarNuevosDatos(&registracionPrueba)

	responderRegistracion(writer, 202, datosRegistracion)

}

func AceptarCS(writer http.ResponseWriter, request *http.Request){

  	nuevoEstado(registracionPrueba.estado).aceptarPorCS(&registracionPrueba)

	responderRegistracion(writer, 202, registracionPrueba)

}
