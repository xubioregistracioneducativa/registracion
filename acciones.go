package main

import (
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "strings"
)

func responderRegistracion(writer http.ResponseWriter, status int, results Registracion){

	writer.Header().Set("Content-Type", "application-json")
	writer.WriteHeader(status)

	json.NewEncoder(writer).Encode(results)
}


func NuevaRegistracion(writer http.ResponseWriter, request *http.Request){

	datosRegistracion := DecodificarRegistracion(request)

	datosRegistracion.estado = estadoInicioRegistracionID

    nuevoEstado(datosRegistracion.estado).ingresarNuevosDatos(&datosRegistracion)

	responderRegistracion(writer, 202, datosRegistracion)

}

func AceptarCS(writer http.ResponseWriter, request *http.Request){

	id := obtenerID(request)

	registracion := obtenerRegistracionPorID(id)

  	nuevoEstado(registracion.estado).aceptarPorCS(&registracion)

	responderRegistracion(writer, 202, registracion)

}

func obtenerID(request *http.Request) int {
	params := mux.Vars(request)
	return params["id"]
}