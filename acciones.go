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

	registracion, err := obtenerRegistracionPorEmail(params["email"])

	if err != nil {
		responderError(writer, 400, err.Error())
		return
	}

	err = validarLink(registracion.IDRegistracion, params["accion"], params["validationCode"]);

	if err != nil {
		responderError(writer, 400, err.Error())
		return
	}

	estado, err := nuevoEstado(registracion.estado)

	if err != nil {
		responderError(writer, 400, err.Error())
		return
	}
	
	switch(params["accion"]) {
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
	case "VencerRegistracion":
		err = estado.vencerRegistracion(&registracion)
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

	err = enviarMailBienvenidaAlumno(&datosRegistracion)
	if err != nil {
		responderError(writer, 400, err.Error())
		return
	}
	err = enviarMailCS(&datosRegistracion)
	if err != nil {
		responderError(writer, 400, err.Error())
		return
	}

	responderJSON(writer, 201, datosRegistracion)

}


