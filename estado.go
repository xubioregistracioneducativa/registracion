package main

type estadoID int

const (  // iota is reset to 0
        estadoInicioRegistracionID estadoID = iota  // c0 == 0
        estadoPendienteAprobacionID   // c1 == 1
        estadoAprobadoID // c2 == 2
        estadoConfirmadoID
)

type estado interface {
  ingresarNuevosDatos(*Registracion) error
  rechazarPorCS(*Registracion) error
  aceptarPorCS(*Registracion) error
  confirmarPorProfesor(*Registracion) error
}

func nuevoEstado (idEstado estadoID) estado {
    switch(idEstado) {
    case estadoInicioRegistracionID:
      return estadoInicioRegistracion{}
    case estadoPendienteAprobacionID:
      return estadoPendienteAprobacion{}
    case estadoAprobadoID:
      return estadoAprobado{}
    case estadoConfirmadoID:
      return estadoConfirmado{}
    default:
    return nil
  }
}