package main

import (
	"errors"
	"fmt"
)

type estadoAprobado struct {

}

func (estado estadoAprobado ) ingresarNuevosDatos (registracion *Registracion) (string, error) {
	return "", errors.New("ERROR_APROBADA_INGRESAR")
}

func (estado estadoAprobado ) rechazarPorCS (registracion *Registracion) (string, error) {
  	return "", errors.New("ERROR_APROBADA_RECHAZAR")
}

func (estado estadoAprobado ) aceptarPorCS (registracion *Registracion) (string, error)  {
	  fmt.Println("Se reenvía mail al alumno y al profesor")
	err := enviarMailProfesorReenviado(registracion)
	if err != nil {
		return "", err
	}
	return "EXITO_ACEPTAR", nil
}

func (estado estadoAprobado ) anularPorCS (registracion *Registracion) (string, error) {
	fmt.Println("Se anula la Registracion")
	registracion.estado = estadoAnuladoID
	err := GetDBHelper().reingresarRegistracion(registracion)
	if err != nil {
		return "", err
	}
	err = enviarMailAnulacionAlumno(registracion)
	if err != nil {
		return "", err
	}
	return "EXITO_ANULAR", nil
}

func (estado estadoAprobado ) confirmarPorProfesor (registracion *Registracion) (string, error) {
	var err error
	fmt.Println("Se Registra el Tenant en Xubio y se avisa al alumno")
	err = registrarTenant(registracion)
	if err != nil {
		return "", err
	}
	registracion.estado = estadoConfirmadoID
	err = GetDBHelper().updateRegistracion(registracion)
	if err != nil {
		return "", err
	}
	err = enviarMailRegistracionAlumno(registracion)
	if err != nil {
		return "", err
	}
	err = enviarMailConfirmadoCS(registracion)
	if err != nil {
		return "", err
	}

	return "EXITO_CONFIRMAR", nil
}

func (estado estadoAprobado ) consultarEstado () string {
	return "ESTADO_APROBADO"
}

func (estado estadoAprobado ) vencerRegistracion (registracion *Registracion) (string, error) {
	return "", errors.New("ERROR_REGISTRACIONINCOMPLETA")
}