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
	if GetDBHelper().emailDeRegistroLibre((*registracion).Email) {
		err = GetDBHelper().insertarNuevaRegistracion(registracion)
	} else {
		err = GetDBHelper().reingresarRegistracion(registracion)
	}
	if err != nil {
		return "", err
	}

	return "EXITO_INGRESAR", nil
}

func (estado estadoRechazado ) rechazarPorCS (registracion *Registracion) (string, error) {
	return "", errors.New("ERROR_RECHAZADA_RECHAZAR")
}

func (estado estadoRechazado ) aceptarPorCS (registracion *Registracion) (string, error){
	return "", errors.New("ERROR_RECHAZADA_ACEPTAR")
}

func (estado estadoRechazado ) anularPorCS (registracion *Registracion) (string, error) {
	return "", errors.New("ERROR_RECHAZADA_ANULAR")
}

func (estado estadoRechazado ) confirmarPorProfesor (registracion *Registracion) (string, error) {
	return "", errors.New("ERROR_RECHAZADA_CONFIRMAR")
}

func (estado estadoRechazado ) consultarEstado () string {
	return "ESTADO_RECHAZADO"
}

func (estado estadoRechazado ) vencerRegistracion (registracion *Registracion) (string, error) {
	return "", errors.New("ERROR_REGISTRACIONINCOMPLETA")
}