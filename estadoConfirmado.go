package main

import (
	"errors"
	"fmt"
)

type estadoConfirmado struct {

}

func (estado estadoConfirmado ) ingresarNuevosDatos (registracion *Registracion) (string, error) {
  return "", errors.New(getMensaje("ERROR_CONFIRMADA_INGRESAR"))
}

func (estado estadoConfirmado ) rechazarPorCS (registracion *Registracion) (string, error) {
  return "", errors.New(getMensaje("ERROR_CONFIRMADA_RECHAZAR"))
}

func (estado estadoConfirmado ) aceptarPorCS (registracion *Registracion) (string, error) {
  return "", errors.New(getMensaje("ERROR_CONFIRMADA_ACEPTAR"))
}

func (estado estadoConfirmado ) anularPorCS (registracion *Registracion) (string, error) {
	return "", errors.New(getMensaje("ERROR_CONFIRMADA_ANULAR"))
}

func (estado estadoConfirmado ) confirmarPorProfesor (registracion *Registracion) (string, error) {
  return "", errors.New(getMensaje("ERROR_CONFIRMADA_CONFIRMAR"))

}

func (estado estadoConfirmado ) consultarEstado () string {
	return getMensaje("ESTADO_CONFIRMADO")
}

func (estado estadoConfirmado ) vencerRegistracion (registracion *Registracion) (string, error){
	fmt.Println("Se guarda la Registracion")
	registracion.estado = estadoInicioRegistracionID
	err := reingresarRegistracion(registracion)
	if err != nil {
		return "", err
	}
	return getMensaje("EXITO_CONFIRMAR"), nil
}