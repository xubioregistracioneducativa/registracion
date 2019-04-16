package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/xubioregistracioneducativa/registracion/configuracion"
	"log"
	"net/http"
)

func responderJSON(writer http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, err = writer.Write([]byte(err.Error()))
		if err != nil {
			log.Panic(err)
		}
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	_ , err = writer.Write([]byte(response))
	if err != nil {
		log.Panic(err)
	}
}
 
func responderExito(w http.ResponseWriter, r *http.Request, code int, message string) {
	//responderJSON(w, code, map[string]string{"error": message})
	http.Redirect(w, r, configuracion.UrlMono() + configuracion.UrlExito() + message, http.StatusSeeOther)
}

func responderError(w http.ResponseWriter, r *http.Request, code int, message string) {
	//responderJSON(w, code, map[string]string{"error": message})
	http.Redirect(w, r, configuracion.UrlMono() + configuracion.UrlError() + message, http.StatusSeeOther)
}

func responderEstado(w http.ResponseWriter, r *http.Request, code int, message string) {
	//responderJSON(w, code, map[string]string{"error": message})
	http.Redirect(w, r, configuracion.UrlMono() + configuracion.UrlConsultarEstado() + message, http.StatusSeeOther)
}

func responderExitoCreado(w http.ResponseWriter, r *http.Request, code int, message string) {
	responderJSON(w, code, map[string]string{"Exito": message})
}

func ModificarRegistracion(writer http.ResponseWriter, request *http.Request){
	defer handlePanic(writer, request)

	params := mux.Vars(request)

	registracion, err := obtenerRegistracionPorEmail(params["email"])

	if err != nil {
		responderError(writer, request, http.StatusBadRequest, err.Error())
		return
	}

	err = validarLink(registracion.IDRegistracion, params["accion"], params["validationCode"]);

	if err != nil {
		responderError(writer, request, http.StatusBadRequest, err.Error())
		return
	}

	estado, err := nuevoEstado(registracion.estado)

	if err != nil {
		responderError(writer, request, http.StatusBadRequest, err.Error())
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
		responderEstado(writer, request, http.StatusAccepted, mensajeEstado)
		return
	case "VencerRegistracion":
		mensajeEstado, err = estado.vencerRegistracion(&registracion)
    default:
    	err = errors.New("ERROR_ACCIONINCORRECTA")
    	log.Println(err)
 	 }

  	if err != nil {
		responderError(writer, request, http.StatusBadRequest, err.Error())
		return
	}

	responderExito(writer, request, http.StatusSeeOther, mensajeEstado)

}

func handlePanic(writer http.ResponseWriter, request *http.Request) {
	if recoveredError := recover(); recoveredError != nil{
		log.Println(recoveredError)
		responderError(writer, request, http.StatusBadRequest,("ERROR_DEFAULT"))
	}
}

func NuevaRegistracion(writer http.ResponseWriter, request *http.Request){

	defer handlePanic( writer, request)

	var err error

	datosRegistracion := DecodificarDatosRegistracion(request)

	err = verificarDatosValidos(&datosRegistracion)

	if err != nil {
		responderError(writer, request, http.StatusBadRequest, err.Error())
		return
	}

	datosRegistracion.Registracion.estado, err = obtenerEstadoIDPorEmail(datosRegistracion.Registracion.Email)

	if err != nil {
		responderError(writer, request, http.StatusBadRequest, err.Error())
		return
	}

	estado, err := nuevoEstado(datosRegistracion.Registracion.estado)

	if err != nil {
		responderError(writer, request, http.StatusBadRequest, err.Error())
		return
	}

	mensajeEstado, err := estado.ingresarNuevosDatos(&datosRegistracion.Registracion)

	if err != nil {
		responderError(writer, request, http.StatusBadRequest, err.Error())
		return
	}

	err = eliminarLinksPorID(datosRegistracion.Registracion.IDRegistracion)

	if err != nil {
		responderError(writer, request, http.StatusBadRequest, err.Error())
		return
	}

	err = generarLinks(datosRegistracion.Registracion.IDRegistracion)

	if err != nil {
		responderError(writer, request, http.StatusBadRequest, err.Error())
		return
	}

	err = enviarMailBienvenidaAlumno(&datosRegistracion.Registracion)
	if err != nil {
		responderError(writer, request, http.StatusBadRequest, err.Error())
		return
	}
	err = enviarMailCS(&datosRegistracion.Registracion)
	if err != nil {
		responderError(writer, request, http.StatusBadRequest, err.Error())
		return
	}

	responderExito(writer, request, http.StatusCreated, mensajeEstado)

}



