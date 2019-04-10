package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
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

func responderExito(w http.ResponseWriter, code int, message string) {
	responderJSON(w, code, map[string]string{"Exito": message})
}


func ModificarRegistracion(writer http.ResponseWriter, request *http.Request){
	defer handlePanic(writer)

	params := mux.Vars(request)

	registracion, err := obtenerRegistracionPorEmail(params["email"])

	if err != nil {
		responderError(writer, http.StatusBadRequest, err.Error())
		return
	}

	err = validarLink(registracion.IDRegistracion, params["accion"], params["validationCode"]);

	if err != nil {
		responderError(writer, http.StatusBadRequest, err.Error())
		return
	}

	estado, err := nuevoEstado(registracion.estado)

	if err != nil {
		responderError(writer, http.StatusBadRequest, err.Error())
		return
	}

	var mensajeEstado string

	switch(params["accion"]) {
    case "AceptarCS":
		mensajeEstado, err = estado.aceptarPorCS(&registracion)
    case "RechazarCS":
		mensajeEstado, err = estado.rechazarPorCS(&registracion)
	case "AnularCS":
		mensajeEstado, err = estado.anularPorCS(&registracion)
    case "ConfirmarProfesor":
		mensajeEstado, err = estado.confirmarPorProfesor(&registracion)
	case "ConsultarEstado":
		mensajeEstado = estado.consultarEstado()
		//http.Redirect(writer, request, "https://localhost:8443/mantenimiento.jsp", http.StatusSeeOther)
		responderJSON(writer, http.StatusAccepted, mensajeEstado)
		return
	case "VencerRegistracion":
		err = estado.vencerRegistracion(&registracion)
    default:
    	err = errors.New(getMensaje("ERROR_ACCIONINCORRECTA"))
 	 }

  	if err != nil {
		return
	}

	responderExito(writer, http.StatusAccepted, mensajeEstado)

}

func handlePanic(writer http.ResponseWriter) {
	if recoveredError := recover(); recoveredError != nil{
		responderError(writer, http.StatusBadRequest,getMensaje("ERROR_DEFAULT"))
	}
}

func NuevaRegistracion(writer http.ResponseWriter, request *http.Request){
	defer handlePanic( writer)

	var err error

	datosRegistracion := DecodificarDatosRegistracion(request)

	err = verificarDatosValidos(&datosRegistracion)

	if err != nil {
		responderError(writer, http.StatusBadRequest, err.Error())
		return
	}

	datosRegistracion.Registracion.estado, err = obtenerEstadoIDPorEmail(datosRegistracion.Registracion.Email)

	if err != nil {
		responderError(writer, http.StatusBadRequest, err.Error())
		return
	}

	estado, err := nuevoEstado(datosRegistracion.Registracion.estado)

	if err != nil {
		responderError(writer, http.StatusBadRequest, err.Error())
		return
	}

	err = estado.ingresarNuevosDatos(&datosRegistracion.Registracion)

	if err != nil {
		responderError(writer, http.StatusBadRequest, err.Error())
		return
	}

	err = eliminarLinksPorID(datosRegistracion.Registracion.IDRegistracion)

	if err != nil {
		responderError(writer, http.StatusBadRequest, err.Error())
		return
	}

	err = generarLinks(datosRegistracion.Registracion.IDRegistracion)

	if err != nil {
		responderError(writer, http.StatusBadRequest, err.Error())
		return
	}

	err = enviarMailBienvenidaAlumno(&datosRegistracion.Registracion)
	if err != nil {
		responderError(writer, http.StatusBadRequest, err.Error())
		return
	}
	err = enviarMailCS(&datosRegistracion.Registracion)
	if err != nil {
		responderError(writer, http.StatusBadRequest, err.Error())
		return
	}

	responderJSON(writer, http.StatusCreated, datosRegistracion)

}



