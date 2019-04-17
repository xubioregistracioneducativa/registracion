package main

import (
	"errors"
	"regexp"
	"strings"
)

func verificarDatosValidos(datosRegistracion *DatosRegistracion) error {
	var err error
	normalizarDatos(&datosRegistracion.Registracion)
	err = verificarLeiTerminos(datosRegistracion.LeiTerminos)
	if err != nil {
		return err
	}
	err = verificarCaptcha(datosRegistracion.CaptchaValue)
	if err != nil {
		return err
	}
	err = verificarCampoVacio(datosRegistracion.Registracion.Nombre, "Nombre")
	if err != nil {
		return errors.New("ERROR_VALIDACION_VACIO_NOMBRE")
	}
	err = verificarCampoVacio(datosRegistracion.Registracion.Apellido, "Apellido")
	if err != nil {
		return errors.New("ERROR_VALIDACION_VACIO_APELLIDO")
	}
	err = verificarEmailAlumno(datosRegistracion.Registracion.Email)
	if err != nil {
		return err
	}
	err = verificarCampoVacio(datosRegistracion.Registracion.Telefono, "Telefono")
	if err != nil {
		return errors.New("ERROR_VALIDACION_VACIO_TELEFONO")
	}
	err = verificarCampoVacio(datosRegistracion.Registracion.Carrera, "Carrera")
	if err != nil {
		return errors.New("ERROR_VALIDACION_VACIO_CARRERA")
	}
	err = verificarCampoVacio(datosRegistracion.Registracion.Clave, "Clave")
	if err != nil {
		return errors.New("ERROR_VALIDACION_VACIO_CLAVE")
	}
	//PAIS
	err = verificarCampoVacio(datosRegistracion.Registracion.NombreProfesor, "Nombre del profesor")
	if err != nil {
		return errors.New("ERROR_VALIDACION_VACIO_NOMBREPROFESOR")
	}
	err = verificarCampoVacio(datosRegistracion.Registracion.ApellidoProfesor, "Apellido del profesor")
	if err != nil {
		return errors.New("ERROR_VALIDACION_VACIO_APELLIDOPROFESOR")
	}
	err = verificarEmailProfesor(datosRegistracion.Registracion.EmailProfesor)
	if err != nil {
		return err
	}
	err = verificarCampoVacio(datosRegistracion.Registracion.Materia, "Materia")
	if err != nil {
		return errors.New("ERROR_VALIDACION_VACIO_MATERIA")
	}
	err = verificarCampoVacio(datosRegistracion.Registracion.Catedra, "Catedra")
	if err != nil {
		return errors.New("ERROR_VALIDACION_VACIO_CATEDRA")
	}
	err = verificarCampoVacio(datosRegistracion.Registracion.Facultad, "Facultad")
	if err != nil {
		return errors.New("ERROR_VALIDACION_VACIO_FACULTAD")
	}
	err = verificarCampoVacio(datosRegistracion.Registracion.Universidad, "Universidad")
	if err != nil {
		return errors.New("ERROR_VALIDACION_VACIO_UNIVERSIDAD")
	}
	return nil
}

func normalizarDatos(registracion *Registracion){
	registracion.Email = strings.TrimSpace(strings.ToUpper(registracion.Email))
	registracion.EmailProfesor = strings.TrimSpace(strings.ToUpper(registracion.EmailProfesor))
	registracion.ApellidoProfesor = strings.TrimSpace(registracion.ApellidoProfesor)
	registracion.NombreProfesor = strings.TrimSpace(registracion.NombreProfesor)
	registracion.Nombre = strings.TrimSpace(registracion.Nombre)
	registracion.Apellido = strings.TrimSpace(registracion.Apellido)
	registracion.Catedra = strings.TrimSpace(registracion.Catedra)
	registracion.Materia = strings.TrimSpace(registracion.Materia)
	registracion.Carrera = strings.TrimSpace(registracion.Carrera)
	registracion.Universidad = strings.TrimSpace(registracion.Universidad)
	registracion.Telefono = strings.TrimSpace(registracion.Telefono)
}

func verificarCaptcha(captchaValue string) error {
	if !ValidateCaptcha(captchaValue) {
		return errors.New("ERROR_VALIDACION_CAPTCHA")
	}
	return nil
}

func verificarLeiTerminos(leiTerminos bool) error {
	if !leiTerminos{
		return errors.New("ERROR_VALIDACION_LEITERMINOS")
	}
	return nil
}




func verificarEmailProfesor(emailProfesor string) error {

	err := verificarEmail(emailProfesor)
	if err != nil{
		return errors.New("ERROR_VALIDACION_INVALIDO_EMAILPROFESOR")
	}
	dominio := emailProfesor[strings.LastIndexByte(emailProfesor, '@'):]
	err = verificarDominioEstudiantil(dominio)
	if err != nil {
		return errors.New("ERROR_VALIDACION_INVALIDO_EMAILPROFESORNOESTUDIANTIL")
	}

	return nil
}

func verificarEmailAlumno(emailAlumno string) error {
	err := verificarEmail(emailAlumno)
	if err != nil{
		return errors.New("ERROR_VALIDACION_INVALIDO_EMAIL")
	}
	return nil
}

func verificarEmail(email string) error {
	err := verificarCampoVacio(email, "Email")
	if err != nil  {
		return errors.New("ERROR_VALIDACION_EMAILINVALIDO_VACIO")
	}
	emailRegexp := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegexp.MatchString(email) {
		return errors.New("ERROR_VALIDACION_EMAILINVALIDO")
	}
	return nil
}

func verificarDominioEstudiantil(dominio string) error {
		dominiosProhibidos :=  []string{"GMAIL", "YAHOO", "HOTMAIL", "OUTLOOK", "LIVE"}
	for i := 0; i < len(dominiosProhibidos); i++ {
		if strings.Contains(dominio, dominiosProhibidos[i]) {
			return errors.New("ErrorInterno: Dominio Invalido")
		}
	}
	return nil
}

func noEstaVacio(string string) error {
	if len(strings.TrimSpace(string)) == 0{
		return errors.New("ErrorInterno: El string solo tiene espacios")
	}
	return nil
}

func verificarCampoVacio(string string, campo string) error {
	err := noEstaVacio(string)
	if err != nil {
		return errors.New("ErrorInterno: El campo " + campo + " esta vacio.")
	}
	return nil
}