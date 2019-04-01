package main

import (
	"encoding/json"
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


func ModificarRegistracion(writer http.ResponseWriter, request *http.Request){

	params := mux.Vars(request)

	link, err := obtenerLinkPorUrl(obtenerUrl(params["input"], params["id"], params["validationCode"]))

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
	case "AnularCS":
		err = estado.anularPorCS(&registracion)
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

	err = generarLinks(datosRegistracion.IDRegistracion, datosRegistracion.Email)

	if err != nil {
		responderError(writer, 400, err.Error())
		return
	}
	responderJSON(writer, 201, datosRegistracion)

}

func VencerRegistracion(writer http.ResponseWriter, request *http.Request){

	params := mux.Vars(request)

	registracion, err := obtenerRegistracionPorEmail(params["email"])

	if err != nil {
		responderError(writer, 400, err.Error())
		return
	}

	_ , err = obtenerLinkPorUrl(obtenerUrl("VencerRegistracion", registracion.Email, "Mkj0WEW1iWJvJGKWXAWG8HkWng4R0maRwxNl2_QOpu8="))

	if err != nil {
		responderError(writer, 400, err.Error())
		return
	}

	estado, err := nuevoEstado(registracion.estado)

	if err != nil {
		responderError(writer, 400, err.Error())
		return
	}

	err = estado.vencerRegistracion(&registracion)

	if err != nil {
		responderError(writer, 400, err.Error())
		return
	}

	responderJSON(writer, 202, registracion)

}

