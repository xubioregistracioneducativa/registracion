package main

import (
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "strconv"
)

func responderJSON(writer http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write([]byte(response))
}
 
func responderError(w http.ResponseWriter, code int, message string) {
	responderJSON(w, code, map[string]string{"error": message})
}


func NuevaRegistracion(writer http.ResponseWriter, request *http.Request){

	var err error

	datosRegistracion := DecodificarRegistracion(request)

	datosRegistracion.estado, err = obtenerEstadoIDPorEmail(datosRegistracion.Email)

	if err != nil {
		responderError(writer, 400, err.Error())
		return
	}

    err = nuevoEstado(datosRegistracion.estado).ingresarNuevosDatos(&datosRegistracion)

    if err != nil {
		responderError(writer, 400, err.Error())
		return
	}

	responderJSON(writer, 202, datosRegistracion)

}

func AceptarCS(writer http.ResponseWriter, request *http.Request){

	id := obtenerID(request)

	registracion, err := obtenerRegistracionPorID(id)

	if err != nil {
		responderError(writer, 400, err.Error())
		return
	}

  	err = nuevoEstado(registracion.estado).aceptarPorCS(&registracion)

  	if err != nil {
		responderError(writer, 400, err.Error())
		return
	}

	responderJSON(writer, 202, registracion)

	return

}

func RechazarCS(writer http.ResponseWriter, request *http.Request){

	id := obtenerID(request)

	registracion, err := obtenerRegistracionPorID(id)

	if err != nil {
		responderError(writer, 400, err.Error())
		return
	}

  	err = nuevoEstado(registracion.estado).rechazarPorCS(&registracion)

  	if err != nil {
		responderError(writer, 400, err.Error())
		return
	}

	responderJSON(writer, 202, registracion)

	return

}

func ConfirmarProfesor(writer http.ResponseWriter, request *http.Request){

	id := obtenerID(request)

	registracion, err := obtenerRegistracionPorID(id)

	if err != nil {
		responderError(writer, 400, err.Error())
		return
	}

  	err = nuevoEstado(registracion.estado).confirmarPorProfesor(&registracion)

  	if err != nil {
		responderError(writer, 400, err.Error())
		return
	}

	responderJSON(writer, 202, registracion)

	return

}



func obtenerID(request *http.Request) int {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	if(err != nil){
		panic(err)
	}

	return id
}