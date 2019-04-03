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
  var err error
  fmt.Println("Se reinicia la registracion a 0")
  registracion.estado = estadoRechazadoID
  err = updateRegistracion(registracion)
  if err != nil {
    return err
  }
  err = enviarMailRechazoAlumno(registracion)
  if err != nil {
    return err
  }
  return nil
}

func (estado estadoPendienteAprobacion ) aceptarPorCS (registracion *Registracion) error {
  var err error
  fmt.Println("Se envía mail al alumno y al profesor")
  registracion.estado = estadoAprobadoID
  err = updateRegistracion(registracion)
  if err != nil {
    return err
  }
  err = enviarMailProfesor(registracion)
  if err != nil {
    return err
  }
  return nil
}

func (estado estadoPendienteAprobacion ) anularPorCS (registracion *Registracion) error {
  fmt.Println("Se anula la Registracion")
  registracion.estado = estadoAnuladoID
  err := reingresarRegistracion(registracion)
  if err != nil {
    return err
  }
  err = enviarMailAnulacionAlumno(registracion)
  if err != nil {
    return err
  }
  return nil
}

func (estado estadoPendienteAprobacion ) confirmarPorProfesor (registracion *Registracion)error {
  return errors.New("Esta registracion esta todavía pendiente")
}

func (estado estadoPendienteAprobacion ) consultarEstado () string {
  return fmt.Sprint("Esta registracion se encuentra pendiente de aprobacion por nuestro equipo")
}

func (estado estadoPendienteAprobacion ) vencerRegistracion (registracion *Registracion) error{
  return errors.New("No se puede vencer una registracion que no esta completa")
}