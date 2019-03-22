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
		_, err = writer.Write([]byte(err.Error()))
		if err != nil {
			panic(err)
		}
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	_ , err = writer.Write([]byte(response))
	if err != nil {
		panic(err)
	}
}
 
func responderError(w http.ResponseWriter, code int, message string) {
	responderJSON(w, code, map[string]string{"error": message})
}


func ModificarRegistracion(writer http.ResponseWriter, request *http.Request){

	params := mux.Vars(request)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responderError(writer, 400, err.Error())
		return
	}

	registracion, err := obtenerRegistracionPorID(id)

	if err != nil {
		responderError(writer, 400, err.Error())
		return
	}

	estado, err := nuevoEstado(registracion.estado)

	if err != nil {
		responderError(writer, 400, err.Error())
		return
	}

	input := params["input"]
	
	switch(input) {
    case "AceptarCS":
      err = estado.aceptarPorCS(&registracion)
    case "RechazarCS":
      err = estado.rechazarPorCS(&registracion)
    case "ConfirmarProfesor":
      err = estado.confirmarPorProfesor(&registracion)
    default:
    	responderError(writer, 400, "No se reconoce el input")
 	 }

  	if err != nil {
		responderError(writer, 400, err.Error())
		return
	}

	responderJSON(writer, 202, registracion)
}

func NuevaRegistracion(writer http.ResponseWriter, request *http.Request){

	var err error

	datosRegistracion := DecodificarRegistracion(request)

	datosRegistracion.estado, err = obtenerEstadoIDPorEmail(datosRegistracion.Email)

	if err != nil {
		responderError(writer, 400, err.Error())
		return
	}

    estado, err := nuevoEstado(datosRegistracion.estado)

    if err != nil {
		responderError(writer, 400, err.Error())
		return
	}

    err = estado.ingresarNuevosDatos(&datosRegistracion)

    if err != nil {
		responderError(writer, 400, err.Error())
		return
	}

	responderJSON(writer, 202, datosRegistracion)

}
