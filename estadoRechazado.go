package main

import (
	"errors"
	"fmt"
)

type estadoRechazado struct {

}

func (estado estadoRechazado ) ingresarNuevosDatos (registracion *Registracion) (string, error) {
	fmt.Println("Se guarda la Registracion")
	registracion.estado = estadoPendienteAprobacionID
	var err error
	if emailDeRegistroLibre((*registracion).Email) {
		err = insertarNuevaRegistracion(registracion)
	} else {
		err = reingresarRegistracion(registracion)
	}
	if err != nil {
		return "", err
	}

	return getMensaje("EXITO_INGRESAR"), nil
}

func (estado estadoRechazado ) rechazarPorCS (registracion *Registracion) (string, error) {
	return "", errors.New(getMensaje("ERROR_RECHAZADA_RECHAZAR"))
}

func (estado estadoRechazado ) aceptarPorCS (registracion *Registracion) (string, error){
	return "", errors.New(getMensaje("ERROR_RECHAZADA_ACEPTAR"))
}

func (estado estadoRechazado ) anularPorCS (registracion *Registracion) (string, error) {
	return "", errors.New(getMensaje("ERROR_RECHAZADA_ANULAR"))
}

func (estado estadoRechazado ) confirmarPorProfesor (registracion *Registracion) (string, error) {
	return "", errors.New(getMensaje("ERROR_RECHAZADA_CONFIRMAR"))
}

func (estado estadoRechazado ) consultarEstado () string {
	return getMensaje("ESTADO_RECHAZADO")
}

func (estado estadoRechazado ) vencerRegistracion (registracion *Registracion) (string, error) {
	return "", errors.New(getMensaje("ERROR_REGISTRACIONINCOMPLETA"))
}