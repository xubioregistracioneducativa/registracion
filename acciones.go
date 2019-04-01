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

	link, err := obtenerLinkPorUrl(obtenerUrl(params["input"], id, params["validationCode"]))

	if err != nil {
		responderError(writer, 400, err.Error())
		return
	}

	registracion, err := obtenerRegistracionPorID(link.RegistracionID)

	if err != nil {
		responderError(writer, 400, err.Error())
		return
	}

	estado, err := nuevoEstado(registracion.estado)

	if err != nil {
		responderError(writer, 400, err.Error())
		return
	}
	
	switch(link.Input) {
    case "AceptarCS":
      err = estado.aceptarPorCS(&registracion)
    case "RechazarCS":
      err = estado.rechazarPorCS(&registracion)
    case "ConfirmarProfesor":
      err = estado.confirmarPorProfesor(&registracion)
	case "ConsultarEstado":
		mensajeEstado := estado.consultarEstado()
		responderJSON(writer, 202, mensajeEstado)
		return
    default:
    	responderError(writer, 400, "No se reconoce el input")
		return
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

	err = eliminarLinksPorID(datosRegistracion.IDRegistracion)

	if err != nil {
		responderError(writer, 400, err.Error())
		return
	}

	err = generarLinks(datosRegistracion.IDRegistracion)

	if err != nil {
		responderError(writer, 400, err.Error())
		return
	}
	responderJSON(writer, 202, datosRegistracion)

}
