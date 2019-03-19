package main

import (
    "net/http"
    "encoding/json"
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

  	nuevoEstado(registracionPrueba.estado).aceptarPorCS(&registracionPrueba)

	responderRegistracion(writer, 202, registracionPrueba)

}
