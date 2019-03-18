package main

import "fmt"

type estadoPendienteAprobacion struct {

}

func (estado estadoPendienteAprobacion ) ingresarNuevosDatos (registracion *Registracion) {
  fmt.Println("ERROR: No se puede enviar una nueva registracion mientras esta se encuentra pendiente")
}

func (estado estadoPendienteAprobacion ) rechazarPorCS (registracion *Registracion) {
  fmt.Println("Se reinicia la registracion a 0")
  registracion.estado = estadoInicioRegistracionID
}

func (estado estadoPendienteAprobacion ) aceptarPorCS (registracion *Registracion) {
  fmt.Println("Se envía mail al alumno y al profesor")
  registracion.estado = estadoAprobadoID
}

func (estado estadoPendienteAprobacion ) confirmarPorProfesor (registracion *Registracion) {
  fmt.Println("ERROR: Esta registracion esta todavía pendiente")
}