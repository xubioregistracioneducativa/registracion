package main

import (
	"errors"
	"fmt"
)

type estadoConfirmado struct {

}

func (estado estadoConfirmado ) ingresarNuevosDatos (registracion *Registracion) error {
  return errors.New("Esta registracion ya fue confirmada")
}

func (estado estadoConfirmado ) rechazarPorCS (registracion *Registracion) error {
  return errors.New("Esta registracion ya fue confirmada")
}

func (estado estadoConfirmado ) aceptarPorCS (registracion *Registracion) error {
  return errors.New("Esta registracion ya fue confirmada")
}

func (estado estadoConfirmado ) anularPorCS (registracion *Registracion) error {
	return errors.New("Esta registracion ya fue confirmada, por lo tanto no puede anularse")
}

func (estado estadoConfirmado ) confirmarPorProfesor (registracion *Registracion) error {
  return errors.New("Esta registracion ya fue confirmada")

}

func (estado estadoConfirmado ) consultarEstado () string {
	return fmt.Sprint("Esta registracion ya fue completada, podes entrar a Xubio.com e ingresar con tu email y contraseña")
}