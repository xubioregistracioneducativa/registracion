package main

import (
	"errors"
	"fmt"
)

type estadoAprobado struct {

}

func (estado estadoAprobado ) ingresarNuevosDatos (registracion *Registracion) (string, error) {
	fmt.Println("Se guarda la Registracion")
  	registracion.estado = estadoPendienteAprobacionID
  	err := reingresarRegistracion(registracion)
  	if err != nil {
  		return "", err
  	}
	return getMensaje("EXITO_INGRESAR"), nil
}

func (estado estadoAprobado ) rechazarPorCS (registracion *Registracion) (string, error) {
  	return "", errors.New(getMensaje("ERROR_APROBADA_RECHAZAR"))
}

func (estado estadoAprobado ) aceptarPorCS (registracion *Registracion) (string, error)  {
	  fmt.Println("Se reenvía mail al alumno y al profesor")
	err := enviarMailProfesorReenviado(registracion)
	if err != nil {
		return "", err
	}
	return getMensaje("EXITO_ACEPTAR"), nil
}

func (estado estadoAprobado ) anularPorCS (registracion *Registracion) (string, error) {
	fmt.Println("Se anula la Registracion")
	registracion.estado = estadoAnuladoID
	err := reingresarRegistracion(registracion)
	if err != nil {
		return "", err
	}
	err = enviarMailAnulacionAlumno(registracion)
	if err != nil {
		return "", err
	}
	return getMensaje("EXITO_ANULAR"), nil
}

func (estado estadoAprobado ) confirmarPorProfesor (registracion *Registracion) (string, error) {
	var err error
	fmt.Println("Se Registra el Tenant en Xubio y se avisa al alumno")
	err = registrarTenant(registracion)
	if err != nil {
		return "", err
	}
	registracion.estado = estadoConfirmadoID
	err = updateRegistracion(registracion)
	if err != nil {
		return "", err
	}
	err = enviarMailRegistracionAlumno(registracion)
	if err != nil {
		return "", err
	}

	return getMensaje("EXITO_CONFIRMAR"), nil
}

func (estado estadoAprobado ) consultarEstado () string {
	return getMensaje("ESTADO_APROBADO")
}

func (estado estadoAprobado ) vencerRegistracion (registracion *Registracion) (string, error) {
	return "", errors.New(getMensaje("ERROR_REGISTRACIONINCOMPLETA"))
}