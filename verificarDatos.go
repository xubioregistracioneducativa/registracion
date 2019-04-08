package main

import (
	"errors"
	"github.com/xubioregistracioneducativa/registracion/configuracion"
	"regexp"
	"strings"
)

func verificarDatosValidos(datosRegistracion *DatosRegistracion) error {
	var err error
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
		return err
	}
	err = verificarCampoVacio(datosRegistracion.Registracion.Apellido, "Apellido")
	if err != nil {
		return err
	}
	err = verificarEmailAlumno(datosRegistracion.Registracion.Email)
	if err != nil {
		return err
	}
	err = verificarCampoVacio(datosRegistracion.Registracion.Telefono, "Telefono")
	if err != nil {
		return err
	}
	err = verificarCampoVacio(datosRegistracion.Registracion.Carrera, "Carrera")
	if err != nil {
		return err
	}
	err = verificarCampoVacio(datosRegistracion.Registracion.Clave, "Clave")
	if err != nil {
		return err
	}
	//PAIS
	err = verificarCampoVacio(datosRegistracion.Registracion.NombreProfesor, "Nombre del profesor")
	if err != nil {
		return err
	}
	err = verificarCampoVacio(datosRegistracion.Registracion.ApellidoProfesor, "Apellido del profesor")
	if err != nil {
		return err
	}
	err = verificarEmailProfesor(datosRegistracion.Registracion.EmailProfesor)
	if err != nil {
		return err
	}
	err = verificarCampoVacio(datosRegistracion.Registracion.Materia, "Materia")
	if err != nil {
		return err
	}
	err = verificarCampoVacio(datosRegistracion.Registracion.Catedra, "Catedra")
	if err != nil {
		return err
	}
	err = verificarCampoVacio(datosRegistracion.Registracion.Facultad, "Facultad")
	if err != nil {
		return err
	}
	err = verificarCampoVacio(datosRegistracion.Registracion.Universidad, "Universidad")
	if err != nil {
		return err
	}
	return nil
}

func verificarCaptcha(captchaValue string) error {
	if !validarCaptcha(captchaValue) {
		return errors.New("La validacion captcha es incorrecta")
	}
	return nil
}

func validarCaptcha(captchaValue string) bool {
	if !configuracion.ValidaCaptcha(){
		return true
	}
	return false //TODO
}

func verificarLeiTerminos(leiTerminos bool) error {
	if !leiTerminos{
		return errors.New("Debe aceptar los terminos y condiciones para seguir adelante")
	}
	return nil
}




func verificarEmailProfesor(emailProfesor string) error {

	err := verificarEmail(emailProfesor)
	if err != nil{
		return errors.New("El email del profesor es invalido: " + err.Error())
	}
	dominio := emailProfesor[strings.LastIndexByte(emailProfesor, '@'):]
	err = verificarDominioEstudiantil(dominio)
	if err != nil {
		return errors.New("El Profesor es invalido: " + err.Error())
	}

	return nil
}

func verificarEmailAlumno(emailAlumno string) error {
	err := verificarEmail(emailAlumno)
	if err != nil{
		return errors.New("EL Alumno es invalido: " + err.Error())
	}
	return nil
}

func verificarEmail(email string) error {
	err := verificarCampoVacio(email, "Email")
	if err != nil  {
		return err
	}
	emailRegexp := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegexp.MatchString(email) {
		return errors.New("El formato del Email es incorrecto")
	}
	return nil
}

func verificarDominioEstudiantil(dominio string) error {
	dominiosProhibidos :=  []string{"gmail", "yahoo", "hotmail", "outlook", "live"}
	for i := 0; i < len(dominiosProhibidos); i++ {
		if strings.Contains(dominio, dominiosProhibidos[i]) {
			return errors.New("El email no es el de la facultad")
		}
	}
	return nil
}

func noEstaVacio(string string) error {
	if len(strings.TrimSpace(string)) == 0{
		return errors.New("El string solo tiene espacios")
	}
	return nil
}

func verificarCampoVacio(string string, campo string) error {
	err := noEstaVacio(string)
	if err != nil {
		return errors.New("El campo " + campo + " esta vacio.")
	}
	return nil
}