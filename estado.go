package main

import (
    "errors"
    "log"
)

type estadoID int

const (  // iota is reset to 0
        estadoInicioRegistracionID estadoID = iota  // c0 == 0
        estadoPendienteAprobacionID   // c1 == 1
        estadoAprobadoID // c2 == 2
        estadoConfirmadoID
        estadoRechazadoID
        estadoAnuladoID
)

type estado interface {
  ingresarNuevosDatos(*Registracion) (string, error)
  rechazarPorCS(*Registracion) (string, error)
  aceptarPorCS(*Registracion) (string, error)
  anularPorCS(*Registracion) (string, error)
  confirmarPorProfesor(*Registracion) (string, error)
  consultarEstado() string
  vencerRegistracion(*Registracion) (string, error)
}

func nuevoEstado (idEstado estadoID) (estado, error) {
    switch(idEstado) {
    case estadoInicioRegistracionID:
      return estadoInicioRegistracion{}, nil
    case estadoPendienteAprobacionID:
      return estadoPendienteAprobacion{}, nil
    case estadoAprobadoID:
      return estadoAprobado{}, nil
    case estadoConfirmadoID:
      return estadoConfirmado{}, nil
    case estadoRechazadoID:
        return estadoRechazado{}, nil
    case estadoAnuladoID:
        return estadoAnulado{}, nil
    default:
        log.Panic("Esta registracion se encuentra en un estado desconocido")
        return nil, errors.New("Esta registracion se encuentra en un estado desconocido")
  }
}