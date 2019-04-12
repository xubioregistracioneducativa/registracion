package main

import (
	"net/http"
)

/*
func DecodificarDatosRegistracion(request *http.Request) DatosRegistracion {

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
 */

func DecodificarDatosRegistracion(request *http.Request) DatosRegistracion {

	err := request.ParseForm()

	if err != nil {
		panic(err)
	}
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
	datosRegistracion.Registracion.NombreProfesor = request.FormValue("nombreProfesor")
	datosRegistracion.Registracion.ApellidoProfesor = request.FormValue("apellidoProfesor")
	datosRegistracion.Registracion.EmailProfesor = request.FormValue("emailProfesor")
	datosRegistracion.CaptchaValue = "TODO"
	datosRegistracion.LeiTerminos = true


	//Para cerrar la lectura de algo
	defer request.Body.Close()

	return datosRegistracion

}