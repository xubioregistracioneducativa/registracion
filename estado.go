package main

import "errors"

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
  ingresarNuevosDatos(*Registracion) error
  rechazarPorCS(*Registracion) error
  aceptarPorCS(*Registracion) error
  anularPorCS(*Registracion) error
  confirmarPorProfesor(*Registracion) error
  consultarEstado() string
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
    return nil, errors.New("Esta registracion se encuentra en un estado desconocido")
  }
}