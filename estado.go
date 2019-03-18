package main

const (  // iota is reset to 0
        estadoInicioRegistracionID = iota  // c0 == 0
        estadoPendienteAprobacionID   // c1 == 1
        estadoAprobadoID // c2 == 2
        estadoConfirmadoID
)

type estado interface {
  ingresarNuevosDatos(*Registracion)
  rechazarPorCS(*Registracion)
  aceptarPorCS(*Registracion)
  confirmarPorProfesor(*Registracion)
}

func nuevoEstado (idEstado int) estado {
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