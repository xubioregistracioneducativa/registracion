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
  aceptarPorProfesor(*Registracion)
}

func nuevoEstado (idEstado int) Estado {
    switch(idEstado) {
    case estadoInicioRegistracionID:
      return EstadoInicioRegistracion{}
    case estadoPendienteAprobacionID:
      return EstadoPendienteAprobacion{}
    case estadoAprobadoID:
      return EstadoAprobado{}
    case estadoConfirmadoID:
      return EstadoConfirmado{}
    default:
    return nil
  }
}