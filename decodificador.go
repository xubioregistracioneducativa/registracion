package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func DecodificarDatosRegistracion(request *http.Request) DatosRegistracion {

	var datosRegistracion DatosRegistracion
	tipoDeRequest := request.Header.Get("Content-Type")
	if tipoDeRequest == "application/x-www-form-urlencoded" || tipoDeRequest == "application/x-www-form-urlencoded;charset=UTF-8" {
		datosRegistracion = DecodificarForm(request)
	}else if tipoDeRequest == "application/json" {
		datosRegistracion = DecodificarJson(request)
	} else {
		log.Println(tipoDeRequest)
		datosRegistracion = DecodificarJson(request)
	}
	return datosRegistracion

}

func DecodificarJson(request *http.Request) DatosRegistracion {

	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields ()
	var datosRegistracion DatosRegistracion

	//&nombre_var para decirle que es la var que no tiene datos y va a tener que rellenar
	var err = decoder.Decode(&datosRegistracion)

	if(err != nil){
		panic(err)
	}

	//Para cerrar la lectura de algo
	defer request.Body.Close()

	return datosRegistracion

}

func DecodificarForm(request *http.Request) DatosRegistracion {

	//Para cerrar la lectura
	defer request.Body.Close()

	err := request.ParseForm()

	if err != nil {
		log.Panic(err)
	};
	var datosRegistracion DatosRegistracion
	datosRegistracion.Registracion.Nombre = request.FormValue("nombre")
	datosRegistracion.Registracion.Apellido = request.FormValue("apellido")
	datosRegistracion.Registracion.Email = request.FormValue("email")
	datosRegistracion.Registracion.Telefono = request.FormValue("telefono")
	datosRegistracion.Registracion.Clave = request.FormValue("clave")
	datosRegistracion.Registracion.Materia = request.FormValue("materia")
	datosRegistracion.Registracion.Catedra = request.FormValue("catedra")
	datosRegistracion.Registracion.Carrera = request.FormValue("carrera")
	datosRegistracion.Registracion.Facultad = request.FormValue("facultad")
	datosRegistracion.Registracion.Universidad = request.FormValue("universidad")
	datosRegistracion.Registracion.NombreProfesor = request.FormValue("nombreprofesor")
	datosRegistracion.Registracion.ApellidoProfesor = request.FormValue("apellidoprofesor")
	datosRegistracion.Registracion.EmailProfesor = request.FormValue("emailprofesor")
	datosRegistracion.Registracion.Utm_source = request.FormValue("utm_source")
	datosRegistracion.Registracion.Utm_medium = request.FormValue("utm_medium")
	datosRegistracion.Registracion.Utm_term = request.FormValue("utm_term")
	datosRegistracion.Registracion.Utm_content = request.FormValue("utm_content")
	datosRegistracion.Registracion.Utm_campaign = request.FormValue("utm_campaign")
	datosRegistracion.CaptchaValue = request.FormValue("g-recaptcha-response")
	datosRegistracion.LeiTerminos = parseBool(request.FormValue("terminosycondiciones"))


	return datosRegistracion

}

func parseBool(booleano string) bool {
	if booleano == "on"{
		return true
	}
	return false
}