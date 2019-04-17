package main

import "github.com/xubioregistracioneducativa/registracion/configuracion"

//MAIL DE INICIO REGISTRACION CS

func enviarMailCS(registracion *Registracion) error{
	linkAceptado, err := GetDBHelper().obtenerLink(registracion.IDRegistracion, "AceptarCS")
	if err != nil {
		return err
	}
	linkRechazado, err := GetDBHelper().obtenerLink(registracion.IDRegistracion, "RechazarCS")
	if err != nil {
		return err
	}
	linkAnulado, err := GetDBHelper().obtenerLink(registracion.IDRegistracion, "AnularCS")
	if err != nil {
		return err
	}

	cuerpo := obtenerCuerpoMailCS(registracion, obtenerUrlLink(&linkAceptado, registracion.Email), obtenerUrlLink(&linkRechazado, registracion.Email), obtenerUrlLink(&linkAnulado, registracion.Email))
	asunto := "Nueva Registracion Educativa " + registracion.Email
	err = enviarMail(configuracion.EmailCS(), asunto , cuerpo)
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
	linkConsultarEstado, err := GetDBHelper().obtenerLink(registracion.IDRegistracion, "ConsultarEstado")
	if err != nil {
		return err
	}
	cuerpo := bienvenidaStudent(registracion, obtenerUrlLink(&linkConsultarEstado, registracion.Email))
	err = enviarMail(registracion.Email, getMensaje("EMAIL_ASUNTO_BIENVENIDA"), cuerpo)
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
	err := enviarMail(registracion.Email, getMensaje("EMAIL_ASUNTO_RECHAZO"), cuerpo)
	if err != nil {
		return err
	}
	return nil
}

func mailDeRechazoAlumnos() string{
	cuerpo := "Tu registracion ha sido rechazada por el equipo de Xubio por no poder corroborar los datos enviados."
	cuerpo += "<br>"
	cuerpo += "Podes reintentar usando otro profesor en "
	cuerpo += "<br>"
	cuerpo += getButton("Xubio Educativo - Nueva Registracion",obtenerUrlXubioNuevaRegistracion() )
	cuerpo += "<br><br>"
	return cuerpo
}

//MAIL A PROFESOR

func enviarMailProfesor(registracion *Registracion) error {

	linkConfirmado, err := GetDBHelper().obtenerLink(registracion.IDRegistracion, "ConfirmarProfesor")
	if err != nil {
		return err
	}
	cuerpo := mailProfesor(registracion, obtenerUrlLink(&linkConfirmado, registracion.Email))
	err = enviarMail(registracion.EmailProfesor, getMensaje("EMAIL_ASUNTO_ALUMNONUEVOPROFESOR"), cuerpo)
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

//REENVIO MAIL PROFESOR

func enviarMailProfesorReenviado(registracion *Registracion) error {

	linkConfirmado, err := GetDBHelper().obtenerLink(registracion.IDRegistracion, "ConfirmarProfesor")
	if err != nil {
		return err
	}
	cuerpo := mailProfesor(registracion, obtenerUrlLink(&linkConfirmado, registracion.Email))
	err = enviarMail(registracion.EmailProfesor, getMensaje("EMAIL_ASUNTO_REINTENTARALUMNOPROFESOR" ), cuerpo)
	if err != nil {
		return err
	}
	return nil
}

//MAIL DE ANULACION

func enviarMailAnulacionAlumno(registracion *Registracion) error {

	cuerpo := mailDeAnulacionAlumnos()
	err := enviarMail(registracion.Email, getMensaje("EMAIL_ASUNTO_ANULACION"), cuerpo)
	if err != nil {
		return err
	}
	return nil
}

func mailDeAnulacionAlumnos() string{
	cuerpo := "Tu registracion ha sido anulada porque tu profesor no confirmó tu registracion."
	cuerpo += "<br>"
	cuerpo += "Podes reintentar la registracion en: "
	cuerpo += "<br>"
	cuerpo += getButton("Xubio Educativo - Nueva Registracion", obtenerUrlXubioNuevaRegistracion() )
	cuerpo += "<br><br>"
	return cuerpo
}

//MAIL ACEPTADO

func enviarMailAceptacionAlumno(registracion *Registracion) error {

	cuerpo := mailDeAceptacionAlumnos()
	err := enviarMail(registracion.Email, getMensaje("EMAIL_ASUNTO_ACEPTACION"), cuerpo)
	if err != nil {
		return err
	}
	return nil
}

func mailDeAceptacionAlumnos() string{
	cuerpo := "Tu registracion ha sido aceptada por el equipo de Xubio y se envió un mail a tu profesor para que confirme que sos su alumno."
	cuerpo += "<br>"
	cuerpo += "Si en dos semanas tu profesor no acepta, tu registracion quedará anulada."
	cuerpo += "<br><br>"
	return cuerpo
}

//MAIL REGISTRACION

func enviarMailRegistracionAlumno(registracion *Registracion) error {

	cuerpo := mailDeRegistracionAlumnos()
	err := enviarMail(registracion.Email, getMensaje("EMAIL_ASUNTO_REGISTROTENANT"), cuerpo)
	if err != nil {
		return err
	}
	return nil
}

func mailDeRegistracionAlumnos() string {
	cuerpo := "Tu registracion ha sido exitosa."
	cuerpo += "Ingresa con tu email y contraseña en "
	cuerpo += "<br>"
	cuerpo += getButton("Xubio", configuracion.UrlMono())
	cuerpo += "<br><br>"
	cuerpo += "En el caso que no recuerdes tu contraseña, podes recuperarla usando Recuperar contraseña. "
	return cuerpo
}
