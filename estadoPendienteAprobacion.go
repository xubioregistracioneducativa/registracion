package main

import (
	"errors"
	"fmt"
)

type estadoPendienteAprobacion struct {

}

func (estado estadoPendienteAprobacion ) ingresarNuevosDatos (registracion *Registracion) (string, error) {
  return "", errors.New(getMensaje("ERROR_PENDIENTE_INGRESAR"))
}

func (estado estadoPendienteAprobacion ) rechazarPorCS (registracion *Registracion) (string, error) {
  var err error
  fmt.Println("Se reinicia la registracion a 0")
  registracion.estado = estadoRechazadoID
  err = updateRegistracion(registracion)
  if err != nil {
    return "", err
  }
  err = enviarMailRechazoAlumno(registracion)
  if err != nil {
    return "", err
  }
  return getMensaje("EXITO_RECHAZAR"), nil
}

func (estado estadoPendienteAprobacion ) aceptarPorCS (registracion *Registracion) (string, error) {
  var err error
  fmt.Println("Se envía mail al alumno y al profesor")
  registracion.estado = estadoAprobadoID
  err = updateRegistracion(registracion)
  if err != nil {
    return "", err
  }
  err = enviarMailProfesor(registracion)
  if err != nil {
    return "", err
  }
  err = enviarMailAceptacionAlumno(registracion)
  if err != nil {
    return "", err
  }
  return getMensaje("EXITO_ACEPTAR"), nil
}

func (estado estadoPendienteAprobacion ) anularPorCS (registracion *Registracion) (string, error) {
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

func (estado estadoPendienteAprobacion ) confirmarPorProfesor (registracion *Registracion) (string, error) {
  return "", errors.New(getMensaje("ERROR_PENDIENTE_CONFIRMAR"))
}

func (estado estadoPendienteAprobacion ) consultarEstado () string {
  return getMensaje("ESTADO_PENDIENTE")
}

func (estado estadoPendienteAprobacion ) vencerRegistracion (registracion *Registracion) (string, error) {
  return "", errors.New(getMensaje("ERROR_REGISTRACIONINCOMPLETA"))
}