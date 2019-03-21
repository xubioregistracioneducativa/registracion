package main

import (
	"errors"
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

func (estado estadoConfirmado ) confirmarPorProfesor (registracion *Registracion) error {
  return errors.New("Esta registracion ya fue confirmada")

}