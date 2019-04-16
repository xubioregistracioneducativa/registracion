package main

import (
	"fmt"
	"errors"
)

type estadoAnulado struct {

}

func (estado estadoAnulado ) ingresarNuevosDatos (registracion *Registracion) (string, error) {
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

func (estado estadoAnulado ) rechazarPorCS (registracion *Registracion)(string, error) {
	return "", errors.New("ERROR_ANULADA_RECHAZAR")
}

func (estado estadoAnulado ) aceptarPorCS (registracion *Registracion) (string, error){
	return "", errors.New("ERROR_ANULADA_ACEPTAR")
}

func (estado estadoAnulado ) anularPorCS (registracion *Registracion) (string, error) {
	return "", errors.New("ERROR_ANULADA_ANULAR")
}

func (estado estadoAnulado ) confirmarPorProfesor (registracion *Registracion) (string, error) {
	return "", errors.New("ERROR_ANULADA_CONFIRMAR")
}

func (estado estadoAnulado ) consultarEstado () string {
	return "ESTADO_ANULADO"
}

func (estado estadoAnulado ) vencerRegistracion (registracion *Registracion) (string, error) {
	return "", errors.New("ERROR_REGISTRACIONINCOMPLETA")
}