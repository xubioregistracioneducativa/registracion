package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/xubioregistracioneducativa/registracion/configuracion"
	"log"
	"net/http"
)

type TipoDeRespuesta int

const (  // iota is reset to 0
	responderExito TipoDeRespuesta = iota  // c0 == 0
	responderError   // c1 == 1
	responderConsulta // c2 == 2
)

func ModificarRegistracion(writer http.ResponseWriter, request *http.Request){
	defer handlePanic(writer, request)

	params := mux.Vars(request)

	tipoDeRespuesta, mensaje := modificarRegistracion(params["email"], params["accion"], params["validationCode"])

	responderModificacion(writer, request, tipoDeRespuesta, mensaje, params["accion"])

}

func responderModificacion(writer http.ResponseWriter, r *http.Request, tipoDeRespuesta TipoDeRespuesta, codigoDeMensaje string, accion string){
	if accion != "VencerRegistracion" {
		switch tipoDeRespuesta {
		case responderExito:
			responderRedireccionExito(writer, r, http.StatusSeeOther, codigoDeMensaje)
		case responderError:
			responderRedireccionError(writer, r, http.StatusSeeOther, codigoDeMensaje)
		case responderConsulta:
			responderRedireccionConsultarEstado(writer, r, http.StatusSeeOther, codigoDeMensaje)
		default:
			log.Panicln("Se obtuvo un tipo de respuesta inesperada")
		}
	} else {
		switch tipoDeRespuesta {
		case responderExito:
			responderMensajeJson(writer, r, http.StatusOK, codigoDeMensaje)
		case responderError:
			responderMensajeJson(writer, r, http.StatusBadRequest, codigoDeMensaje)
		case responderConsulta:
			responderMensajeJson(writer, r, http.StatusOK, codigoDeMensaje)
		default:
			log.Panicln("Se obtuvo un tipo de respuesta inesperada")
		}
	}
}

func recuperarPanic(tipo *TipoDeRespuesta, mensaje *string){
	if recoveredError := recover(); recoveredError != nil{
		log.Println(recoveredError)
		*mensaje = "ERROR_DEFAULT"
		*tipo = responderError
	}
}

func modificarRegistracion(email string, accion string, validationCode string) (tipoDeRespuesta TipoDeRespuesta, mensaje string )  {

	var err error

	defer recuperarPanic(&tipoDeRespuesta, &mensaje)

	registracion, err := GetDBHelper().obtenerRegistracionPorEmail(email)

	if err != nil {
		return responderError, err.Error()
	}

	err = validarLink(registracion.IDRegistracion, accion , validationCode);

	if err != nil {
		return responderError, err.Error()
	}

	estado, err := nuevoEstado(registracion.estado)

	if err != nil {
		return responderError, err.Error()
	}

	var mensajeEstado string

	switch(accion) {
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
		return responderConsulta, mensajeEstado
	case "VencerRegistracion":
		mensajeEstado, err = estado.vencerRegistracion(&registracion)
	default:
		err = errors.New("ERROR_ACCIONINCORRECTA")
		log.Println(err)
	}

	if err != nil {
		return responderError, err.Error()
	}

	return responderExito, mensajeEstado

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

func handlePanic(writer http.ResponseWriter, request *http.Request) {
	if recoveredError := recover(); recoveredError != nil{
		log.Println(recoveredError)
		responderNuevaRegistracion(writer, request, http.StatusBadRequest,("ERROR_DEFAULT"))
	}
}

//Pisa el codigo porque el ajax solo acepta status 200 TODO
func responderNuevaRegistracion(writer http.ResponseWriter, request *http.Request, code int, codigoDeMensaje string) {
	responderMensajeJson(writer, request,  http.StatusOK, codigoDeMensaje)
}

//Responde un JSON como Codigo y Mensaje
func responderMensajeJson(writer http.ResponseWriter, r *http.Request, code int, codigoDeMensaje string){
	responderJSON(writer, code, map[string]string{"codigo": codigoDeMensaje, "mensaje": getMensaje(codigoDeMensaje)})
}

//Redirecciona a la pantalla de exito con el mensaje de exito
func responderRedireccionExito(w http.ResponseWriter, r *http.Request, code int, message string) {
	http.Redirect(w, r, configuracion.UrlMono() + configuracion.PathExito() + message, http.StatusSeeOther)
}

//Redirecciona a la pantalla de error con el mensaje de error en la mayoria de los casos, pero HANDLE PANIC esta devolviendo JSON TODO
func responderRedireccionError(w http.ResponseWriter, r *http.Request, code int, message string) {
	http.Redirect(w, r, configuracion.UrlMono() + configuracion.PathError() + message, http.StatusSeeOther)
}

//Redirecciona a la pantalla de consultar estado con el codigo del estado
func responderRedireccionConsultarEstado(w http.ResponseWriter, r *http.Request, code int, message string) {
	http.Redirect(w, r, configuracion.UrlMono() + configuracion.PathConsultarEstado() + message, http.StatusSeeOther)
}

//Responder un JSON de forma comun
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