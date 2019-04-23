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

func responderMono(writer http.ResponseWriter, r *http.Request, code int, codigoDeMensaje string){
	responderJSON(writer, http.StatusOK, map[string]string{"codigo": codigoDeMensaje, "mensaje": getMensaje(codigoDeMensaje)})
}


func responderExito(w http.ResponseWriter, r *http.Request, code int, message string) {
	http.Redirect(w, r, configuracion.UrlMono() + configuracion.PathExito() + message, http.StatusSeeOther)
}

func responderError(w http.ResponseWriter, r *http.Request, code int, message string) {
	http.Redirect(w, r, configuracion.UrlMono() + configuracion.PathError() + message, http.StatusSeeOther)
}

func responderEstado(w http.ResponseWriter, r *http.Request, code int, message string) {
	http.Redirect(w, r, configuracion.UrlMono() + configuracion.PathConsultarEstado() + message, http.StatusSeeOther)
}

func ModificarRegistracion(writer http.ResponseWriter, request *http.Request){
	defer handlePanic(writer, request)

	params := mux.Vars(request)

	registracion, err := GetDBHelper().obtenerRegistracionPorEmail(params["email"])

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
		if err != nil {
			responderMono(writer, request, http.StatusBadRequest, err.Error())
			return
		}
		responderMono(writer, request, http.StatusAccepted, mensajeEstado)
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
		responderNuevaRegistracion(writer, request, http.StatusBadRequest,("ERROR_DEFAULT"))
	}
}

func NuevaRegistracion(writer http.ResponseWriter, request *http.Request){

	defer handlePanic( writer, request)

	var err error

	datosRegistracion := DecodificarDatosRegistracion(request)

	mensajeExito, err := ingresarNuevaRegistracion(datosRegistracion)

	if err != nil {
		responderNuevaRegistracion(writer, request, http.StatusBadRequest, err.Error())
		return
	}

	responderNuevaRegistracion(writer, request, http.StatusBadRequest, mensajeExito)

}

//Pisa el codigo porque el ajax solo acepta status 200 TODO
func responderNuevaRegistracion(writer http.ResponseWriter, request *http.Request, code int, codigoDeMensaje string) {
	responderJSON(writer, http.StatusOK, map[string]string{"codigo": codigoDeMensaje, "mensaje": getMensaje(codigoDeMensaje)})
}

func ingresarNuevaRegistracion(datosRegistracion DatosRegistracion) (string, error){
	err := verificarDatosValidos(&datosRegistracion)

	if err != nil {
		return "", err
	}

	datosRegistracion.Registracion.estado, err = GetDBHelper().obtenerEstadoIDPorEmail(datosRegistracion.Registracion.Email)

	if err != nil {
		return "", err
	}

	estado, err := nuevoEstado(datosRegistracion.Registracion.estado)

	if err != nil {
		return "", err
	}

	mensajeEstado, err := estado.ingresarNuevosDatos(&datosRegistracion.Registracion)

	if err != nil {
		return "", err
	}

	err = GetDBHelper().eliminarLinksPorID(datosRegistracion.Registracion.IDRegistracion)

	if err != nil {
		return "", err
	}

	err = generarLinks(datosRegistracion.Registracion.IDRegistracion)

	if err != nil {
		return "", err
	}

	err = enviarMailBienvenidaAlumno(&datosRegistracion.Registracion)
	if err != nil {
		return "", err
	}
	err = enviarMailCS(&datosRegistracion.Registracion)
	if err != nil {
		return "", nil
	}

	return mensajeEstado, nil
}