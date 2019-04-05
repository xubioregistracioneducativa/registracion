package main

import (
	"github.com/mattbaird/gochimp"
)

//ENVIAR UN MAIL GENERAL (CANDIDATO A SER UN PACKAGE APARTE QUE SE ENCARGUE DE LOS MAILS)

func enviarMail(email string, asunto string, html string) error {
	//apiKey := os.Getenv("MANDRILL_KEY")
	apiKey := "zd_43LHDhNTbyE89_70HtQ"
	mandrillApi, err := gochimp.NewMandrill(apiKey)

	if err != nil {
		return err
	}

	recipients := []gochimp.Recipient{
		//gochimp.Recipient{Email: email},
		gochimp.Recipient{Email: "sebastian.taka@hotmail.com"},//CONFIGURABLE
	}

	message := gochimp.Message{
		Html:      html,
		Subject:   asunto,
		FromEmail: "noresponder@xubiomail.com",
		FromName:  "Xubio",
		To:        recipients,
	}

	_, err = mandrillApi.MessageSend(message, false)

	if err != nil {
		return err
	}

	return nil
}

//PARA LOS BOTONES DE LOS LINKS

func getButton(value string , link string ) string {
	var button = "<table>"
	button += "<tr>"
	button += "<a href=\"" + link + "\""
	button += "<td align=\"center\" width=\"250\" height=\"20\" bgcolor=\"#cf142b\""
	button += "style=\"font-weight: 700; font-family: 'Open Sans', Arial, Helvetica, sans-serif; "
	button += "font-size: 14px; margin-top: 5px;	padding: 7px 31px;	text-transform: uppercase;	"
	button += "display: block;	margin: 25px auto 0px auto;	text-transform: uppercase;  border-bottom: 3px solid #005e86;  "
	button += "border-top: 0;  border-right: 0;  border-left: 0;  text-decoration: none;  background-color: #f4f4f4; "
	button += "border-radius: 4px !important;  min-width: 55px;  color: #FFF;  background: #0193e1; "
	button += "background: -moz-linear-gradient(top, #00abeb 0%, #027cd8 100%); "
	button += "background: -webkit-gradient(linear, left top, left bottom, color-stop(0%,#2da9dc ), color-stop(100%,#027cd8)); "
	button += "background: -webkit-linear-gradient(top, #00abeb 0%,#027cd8 100%); background: -o-linear-gradient(top, #00abeb 0%,#027cd8 100%); "
	button += "background: -ms-linear-gradient(top, #00abeb 0%,#027cd8 100%); background: linear-gradient(to bottom, #00abeb 0%,#027cd8 100%); "
	button += "filter: progid:DXImageTransform.Microsoft.gradient( startColorstr='#00abeb ', endColorstr='#027cd8',GradientType=0 ); "
	button += "border: 1px solid #00abeb !important;  box-sizing: initial;\">"
	button += value + "</td>"
	button += "</a>"
	button += "</tr>"
	button += "</table>"

	button += "<div style='font-size: 11px; padding-top: 10px; padding-bottom: 10px;'>"
	button += "Si el botón no funciona copiá y pegá el siguiente link en tu navegador: <a href=\"" + link + "\">"+link+"</a>"
	button += "</div>"

	return button;
}

//MAIL DE INICIO REGISTRACION CS

func enviarMailCS(registracion *Registracion) error{
	linkAceptado, err := obtenerLink(registracion.IDRegistracion, "AceptarCS")
	if err != nil {
		return err
	}
	linkRechazado, err := obtenerLink(registracion.IDRegistracion, "RechazarCS")
	if err != nil {
		return err
	}
	linkAnulado, err := obtenerLink(registracion.IDRegistracion, "AnularCS")
	if err != nil {
		return err
	}

	cuerpo := obtenerCuerpoMailCS(registracion, obtenerUrlLink(&linkAceptado, registracion.Email), obtenerUrlLink(&linkRechazado, registracion.Email), obtenerUrlLink(&linkAnulado, registracion.Email))
	asunto := "Nueva Registracion Educativa " + registracion.Email
	err = enviarMail("info@xubio.com", asunto , cuerpo)
	if err != nil {
		return err
	}
	return nil
}

func obtenerCuerpoMailCS(registracion *Registracion, linkAceptado string, linkRechazado string, linkAnulado string ) string {
	cuerpo := mostrarDatosRegistracion(registracion)
	cuerpo += botonesMailCS(linkAceptado, linkRechazado, linkAnulado)
	return cuerpo
}

func mostrarDatosRegistracion(registracion *Registracion) string {
	var result = "Se recibio una nueva registración:"
	result += "<br><br>"
	result += "Datos Alumno:"
	result += "<br>"
	result += "Nombre Alumno: " + registracion.Nombre
	result += "<br>"
	result += "Apellido Alumno: " + registracion.Apellido
	result += "<br>"
	result += "Email Alumno: " + registracion.Email
	result += "<br>"
	result += "Telefono Alumno " + registracion.Telefono
	result += "<br><br>"
	result += "Datos Profesor:"
	result += "<br>"
	result += "Nombre Profesor: " + registracion.NombreProfesor
	result += "<br>"
	result += "Apellido Alumno: " + registracion.ApellidoProfesor
	result += "<br>"
	result += "Email Alumno: " + registracion.EmailProfesor
	result += "<br>"
	result += "Materia " + registracion.Materia
	result += "<br>"
	result += "Catedra " + registracion.Catedra
	result += "<br>"
	result += "Carrera " + registracion.Carrera
	result += "<br>"
	result += "Facultad " + registracion.Facultad
	result += "<br>"
	result += "Universidad " + registracion.Universidad
	result += "<br><br>"
	return result
}

func botonesMailCS(linkAceptado string, linkRechazado string, linkAnulado string ) string {
	var result = "Haga click en Aceptar Registracion para verificar que el profesor es válido y enviarle el mail de confirmacion"
	result += "<br>"
	result += getButton("Aceptar Registracion", linkAceptado)
	result += "<br><br>"
	result += "Haga click en Rechazar Registracion para indicar que el profesor es invalido"
	result += "<br>"
	result += getButton("Rechazar Registracion", linkRechazado)
	result += "<br><br>"
	result += "Haga click en Anular Registracion para cancelar la registracion, incluso luego de haber sido aceptada, usar en caso de que el profesor no responda o por algun otro motivo"
	result += "<br>"
	result += getButton("Anular Registracion", linkAnulado)
	result += "<br><br>"
	result += "Muchas gracias!"
	return result;
}

//MAIL DE INICIO REGISTRACION ALUMNO

func enviarMailBienvenidaAlumno(registracion *Registracion) error {
	linkConsultarEstado, err := obtenerLink(registracion.IDRegistracion, "ConsultarEstado")
	if err != nil {
		return err
	}
	cuerpo := bienvenidaStudent(registracion, obtenerUrlLink(&linkConsultarEstado, registracion.Email))
	err = enviarMail(registracion.Email, "Bienvenido a Xubio Educativo", cuerpo)
	if err != nil {
		return err
	}
	return nil
}
func bienvenidaStudent(registracion *Registracion, linkConsultarEstado string) string{
	cuerpo := "<h3>¡Bienvenido a Xubio Student!</h3>"
	cuerpo += "<br>"
	cuerpo += "Ya recibimos su solicitud y estaremos analizandola en los proximos dias."
	cuerpo += "<br><br>"
	cuerpo += botonesMailAlumno(linkConsultarEstado)
	return cuerpo
}

func botonesMailAlumno(linkConsultarEstado string) string {
	var result = "Haga click en Consultar Estado para consultar el estado de su Registracion en cualquier momento"
	result += "<br>"
	result += getButton("Consultar Estado", linkConsultarEstado)
	result += "<br><br>"
	result += "Muchas gracias!"
	return  result
}

//MAIL DE RECHAZO

func enviarMailRechazoAlumno(registracion *Registracion) error {

	cuerpo := mailDeRechazoAlumnos()
	err := enviarMail(registracion.Email, "Su Registracion educativa fue rechazada", cuerpo)
	if err != nil {
		return err
	}
	return nil
}

func mailDeRechazoAlumnos() string{
	cuerpo := "Tu registracion ha sido rechazada por el equipo de Xubio por no poder corroborar los datos enviados."
	cuerpo += "Podes reintentar usando otro profesor en:"
	cuerpo += getButton("Xubio Educativo","www.xubio.com/educativo/NuevaRegistracion" )
	return cuerpo
}

//MAIL A PROFESOR

func enviarMailProfesor(registracion *Registracion) error {

	linkConfirmado, err := obtenerLink(registracion.IDRegistracion, "ConfirmarProfesor")
	if err != nil {
		return err
	}
	cuerpo := mailProfesor(registracion, obtenerUrlLink(&linkConfirmado, registracion.Email))
	err = enviarMail(registracion.EmailProfesor, "Un alumno suyo quiere una cuenta educativa en Xubio", cuerpo)
	if err != nil {
		return err
	}
	return nil
}

func mailProfesor(registracion *Registracion, linkConfirmado string) string{
	cuerpo := mostrarDatosAlumno(registracion)
	cuerpo += BotonesMailProfesor(linkConfirmado)
	return cuerpo
}

func mostrarDatosAlumno(registracion *Registracion) string {
	var result = "El alumno " + registracion.Apellido + ", " + registracion.Nombre
	result += " con email " + registracion.Email
	result += " indico que es su alumno en la materia " + registracion.Materia + " de la catedra " + registracion.Catedra + " en la carrera " + registracion.Carrera
	result += " para ingresar al programa de Xubio Estudiantil, que permite que acceda a una cuenta premium de Xubio por un año"
	result += "<br><br>"
	return result
}

func BotonesMailProfesor(linkConfirmado string) string {
	var result = "Haga click en Confirmar Alumno para confirmar que este es su alumno y darle acceso a Xubio con una cuenta de estudiante"
	result += "<br>"
	result += getButton("Confirmar Alumno", linkConfirmado)
	result += "<br>"
	result += "En caso de que no sea su alumno, ignore este mensaje."
	result += "<br><br>"
	result += "Muchas gracias!"
	return result
}

//MAIL DE ANULACION

func enviarMailAnulacionAlumno(registracion *Registracion) error {

	cuerpo := mailDeAnulacionAlumnos()
	err := enviarMail(registracion.Email, "Su Registracion educativa fue anulada", cuerpo)
	if err != nil {
		return err
	}
	return nil
}

func mailDeAnulacionAlumnos() string{
	cuerpo := "Tu registracion ha sido anulada porque tu profesor no confirmó tu registracion."
	cuerpo += "Podes reintentar la registracion en:"
	cuerpo += getButton("Xubio Educativo","www.xubio.com/educativo/NuevaRegistracion" )
	return cuerpo
}

//MAIL REGISTRACION

func enviarMailRegistracionAlumno(registracion *Registracion) error {

	cuerpo := mailDeRegistracionAlumnos()
	err := enviarMail(registracion.Email, "Su cuenta en Xubio Educativo fue aceptada", cuerpo)
	if err != nil {
		return err
	}
	return nil
}

func mailDeRegistracionAlumnos() string {
	cuerpo := "Tu registracion ha sido exitosa."
	cuerpo += "Ingresa con tu email y contraseña en "
	cuerpo += getButton("Xubio", "www.xubio.com/")
	cuerpo += "<br><br>"
	cuerpo += "En el caso que no recuerdes tu contraseña, podes recuperarla usando Recuperar contraseña. "
	return cuerpo
}