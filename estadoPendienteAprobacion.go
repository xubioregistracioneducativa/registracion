package main

import (
	"errors"
	"fmt"
)

type estadoPendienteAprobacion struct {

}

func (estado estadoPendienteAprobacion ) ingresarNuevosDatos (registracion *Registracion) error{
  return errors.New("No se puede enviar una nueva registracion mientras esta se encuentra pendiente")
}

func (estado estadoPendienteAprobacion ) rechazarPorCS (registracion *Registracion) error {
  fmt.Println("Se reinicia la registracion a 0")
  registracion.estado = estadoInicioRegistracionID
  updateRegistracion(registracion)
  return nil
}

func (estado estadoPendienteAprobacion ) aceptarPorCS (registracion *Registracion) error {
  fmt.Println("Se envía mail al alumno y al profesor")
  registracion.estado = estadoAprobadoID
  updateRegistracion(registracion)
  return nil
}

func (estado estadoPendienteAprobacion ) confirmarPorProfesor (registracion *Registracion)error {
  return errors.New("Esta registracion esta todavía pendiente")
}