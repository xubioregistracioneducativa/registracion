package main

import (
	"github.com/mattbaird/gochimp"
)



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


func enviarMailCS(registracion *Registracion) error{
	linkAceptado, err := obtenerLink("AceptarCS", registracion.IDRegistracion)
	if err != nil {
		return err
	}
	linkRechazado, err := obtenerLink("RechazarCS", registracion.IDRegistracion)
	if err != nil {
		return err
	}
	linkAnulado, err := obtenerLink("AnularCS", registracion.IDRegistracion)
	if err != nil {
		return err
	}

	cuerpo := obtenerCuerpoMailCS(registracion, obtenerUrlLink(&linkAceptado), obtenerUrlLink(&linkRechazado), obtenerUrlLink(&linkAnulado))
	asunto := "Nueva Registracion Educativa " + registracion.Email
	err = enviarMail("info@xubio.com", asunto , cuerpo)
	if err != nil {
		return err
	}
	return nil
}

func enviarMailAlumno(registracion *Registracion) error {
	linkConsultarEstado, err := obtenerLink("ConsultarEstado", registracion.IDRegistracion)
	if err != nil {
		return err
	}
	cuerpo := bienvenidaStudent(registracion, obtenerUrlLink(&linkConsultarEstado))
	err = enviarMail(registracion.Email, "Bienvenido a Xubio Educativo", cuerpo)
	if err != nil {
		return err
	}
	return nil
}

func bienvenidaStudent(registracion *Registracion, linkConsultarEstado string) string{
	cuerpo := "<h3>¡Bienvenido a Xubio Student!</h3>"
	cuerpo += "<br>"
	cuerpo += "Ya recibimos su solicitud y estaremos analizandola en los proximos dias"
	cuerpo += "<br>"
	cuerpo += botonesMailAlumno(linkConsultarEstado)
	return cuerpo
}

func obtenerCuerpoMailCS(registracion *Registracion, linkAceptado string, linkRechazado string, linkAnulado string ) string {
	cuerpo := mostrarDatosAlumno(registracion)
	cuerpo += botonesMailCS(linkAceptado, linkRechazado, linkAnulado)
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

func botonesMailAlumno(linkConsultarEstado string) string {
	var result = "Haga click en Consultar Estado para consultar el estado de su Registracion en cualquier momento"
	result += "<br>"
	result += getButton("Consultar Estado", linkConsultarEstado)
	result += "<br><br>"
	result += "Muchas gracias!"
	return  result
}

func obtenerMailProfesor(linkConfirmado string) string {
	var result = "Haga click en Confirmar Alumno para confirmar que este es su alumno y darle acceso a Xubio con una cuenta de estudiante"
	result += "<br>"
	result += getButton("Confirmar Alumno", linkConfirmado)
	result += "<br>"
	result += "En caso de que no sea su alumno, ignore este mensaje."
	result += "<br><br>"
	result += "Muchas gracias!"
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