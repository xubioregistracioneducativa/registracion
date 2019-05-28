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
	asunto := getMensaje("EMAIL_ASUNTO_NUEVASOLICITUDCS") + registracion.Email
	err = enviarOGuardarMail(configuracion.EmailCS(), asunto , cuerpo)
	if err != nil {
		return err
	}
	return nil
}

func obtenerCuerpoMailCS(registracion *Registracion, linkAceptado string, linkRechazado string, linkAnulado string ) string {
	cuerpo := "Se recibio una nueva registración:"
	cuerpo += mostrarDatosRegistracion(registracion)
	cuerpo += botonesMailCS(linkAceptado, linkRechazado, linkAnulado)
	return cuerpo
}

func mostrarDatosRegistracion(registracion *Registracion) string {
	result := "<br><br>"
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
	result += "Apellido Profesor: " + registracion.ApellidoProfesor
	result += "<br>"
	result += "Email Profesor: " + registracion.EmailProfesor
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
	err = enviarOGuardarMail(registracion.Email, getMensaje("EMAIL_ASUNTO_BIENVENIDA"), cuerpo)
	if err != nil {
		return err
	}
	return nil
}
func bienvenidaStudent(registracion *Registracion, linkConsultarEstado string) string{
	cuerpo := "Ya recibimos su solicitud y estaremos analizandola en los proximos dias."
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

	cuerpo := mailDeRechazoAlumnos(registracion.Nombre)
	err := enviarOGuardarMail(registracion.Email, getMensaje("EMAIL_ASUNTO_RECHAZO"), cuerpo)
	if err != nil {
		return err
	}
	return nil
}

func mailDeRechazoAlumnos(nombreAlumno string) string{
	cuerpo := "Hola " + nombreAlumno + ". Lamentablemente tu solicitud fue rechazada. Los motivos habituales son: el email de tu profesor no " +
		"existe o fue rechazado, o tu profesor no confirmó que seas su alumno, o no pudimos verificar que sea profesor universitario."
	cuerpo += "<br>"
	cuerpo += "<br>"
	cuerpo += getButton("Volver a crear cuenta universitaria", obtenerUrlXubioNuevaRegistracion() )
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
	err = enviarOGuardarMail(registracion.EmailProfesor, getMensaje("EMAIL_ASUNTO_ALUMNONUEVOPROFESOR"), cuerpo)
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
	result := "Hola " + registracion.NombreProfesor + " " + registracion.ApellidoProfesor + ".<br><br> Te escribimos de Xubio," +
		" la Solución de Gestión online.<br><br> Recibimos una solicitud para crear una cuenta Estudiante en Xubio." +
		" Para verificar que el solicitante es estudiante universitario te pedimos que confirmes si es alumno tuyo. <br><br>"
	result += "\"El alumno " + registracion.Apellido + ", " + registracion.Nombre
	result += " con email " + registracion.Email
	result += " indicó que es su alumno en la materia " + registracion.Materia + " de la cátedra " + registracion.Catedra + " en la carrera " + registracion.Carrera + ".\""
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
	err = enviarOGuardarMail(registracion.EmailProfesor, getMensaje("EMAIL_ASUNTO_REINTENTARALUMNOPROFESOR" ), cuerpo)
	if err != nil {
		return err
	}
	return nil
}

//MAIL DE ANULACION

func enviarMailAnulacionAlumno(registracion *Registracion) error {

	cuerpo := mailDeRechazoAlumnos(registracion.Nombre)
	err := enviarOGuardarMail(registracion.Email, getMensaje("EMAIL_ASUNTO_ANULACION"), cuerpo)
	if err != nil {
		return err
	}
	return nil
}

//MAIL REGISTRACION

func enviarMailRegistracionAlumno(registracion *Registracion) error {

	cuerpo := mailDeRegistracionAlumnos(registracion)
	err := enviarOGuardarMail(registracion.Email, getMensaje("EMAIL_ASUNTO_REGISTROTENANT"), cuerpo)
	if err != nil {
		return err
	}
	return nil
}

func mailDeRegistracionAlumnos(registracion *Registracion) string {
	cuerpo := "Hola " + registracion.Nombre + "! Tu profesor confirmó que eres su alumno, así que ya puedes comenzar a utilizar el plan Estudiantes! " +
		"<br><br> " +
		"Muy importante: " +
		"El plan Estudiantes es gratuito y fue diseñado para que te capacites por tu cuenta en nuestro Centro de Ayuda. " +
		"No cuenta con atención por mail, chat, ni teléfono. " +
		"La cuenta estará activa por 365 días, luego de ese tiempo se eliminará junto con toda la información que ingresaste. " +
		"<br>"
	cuerpo += "Ingresa con tu email y contraseña en Xubio"
	cuerpo += "<br>"
	cuerpo += getButton("Ir a Xubio", configuracion.UrlMono())
	cuerpo += "<br><br>"
	cuerpo += "En el caso que no recuerdes tu contraseña, podes recuperarla usando Recuperar contraseña. "
	return cuerpo
}

// MAIL REGISTRACION CS

func enviarMailConfirmadoCS(registracion *Registracion) error{

	cuerpo := obtenerCuerpoMailConfirmadoCS(registracion)
	asunto := getMensaje("EMAIL_ASUNTO_NUEVACUENTACS") + registracion.Email
	err := enviarOGuardarMail(configuracion.EmailCS(), asunto , cuerpo)
	if err != nil {
		return err
	}
	return nil
}

func obtenerCuerpoMailConfirmadoCS(registracion *Registracion) string {
	cuerpo := "Se creo una nueva cuenta: "
	cuerpo += mostrarDatosRegistracion(registracion)
	return cuerpo
}
