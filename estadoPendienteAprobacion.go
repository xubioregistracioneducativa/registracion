package main

import (
	"errors"
	"fmt"
)

type estadoPendienteAprobacion struct {

}

func (estado estadoPendienteAprobacion ) ingresarNuevosDatos (registracion *Registracion) (string, error) {
  return "", errors.New("ERROR_PENDIENTE_INGRESAR")
}

func (estado estadoPendienteAprobacion ) rechazarPorCS (registracion *Registracion) (string, error) {
  var err error
  fmt.Println("Se reinicia la registracion a 0")
  registracion.estado = estadoRechazadoID
  err = GetDBHelper().updateRegistracion(registracion)
  if err != nil {
    return "", err
  }
  err = enviarMailRechazoAlumno(registracion)
  if err != nil {
    return "", err
  }
  return "EXITO_RECHAZAR", nil
}

func (estado estadoPendienteAprobacion ) aceptarPorCS (registracion *Registracion) (string, error) {
  var err error
  fmt.Println("Se envía mail al profesor")
  registracion.estado = estadoAprobadoID
  err = GetDBHelper().updateRegistracion(registracion)
  if err != nil {
    return "", err
  }
  err = enviarMailProfesor(registracion)
  if err != nil {
    return "", err
  }
  return "EXITO_ACEPTAR", nil
}

func (estado estadoPendienteAprobacion ) anularPorCS (registracion *Registracion) (string, error) {
  fmt.Println("Se anula la Registracion")
  registracion.estado = estadoAnuladoID
  err := GetDBHelper().reingresarRegistracion(registracion)
  if err != nil {
    return "", err
  }
  err = enviarMailAnulacionAlumno(registracion)
  if err != nil {
    return "", err
  }
  return "EXITO_ANULAR", nil
}

func (estado estadoPendienteAprobacion ) confirmarPorProfesor (registracion *Registracion) (string, error) {
  return "", errors.New("ERROR_PENDIENTE_CONFIRMAR")
}

func (estado estadoPendienteAprobacion ) consultarEstado () string {
  return "ESTADO_PENDIENTE"
}

func (estado estadoPendienteAprobacion ) vencerRegistracion (registracion *Registracion) (string, error) {
  return "", errors.New("ERROR_REGISTRACIONINCOMPLETA")
}