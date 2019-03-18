package main

import "fmt"

type estadoAprobado struct {

}

func (estado estadoAprobado ) ingresarNuevosDatos (registracion *Registracion) {
	fmt.Println("Se guarda la Registracion")
  registracion.estado = estadoPendienteAprobacionID
}

func (estado estadoAprobado ) rechazarPorCS (registracion *Registracion) {
  fmt.Println("ERROR: Esta registracion ya fue aceptada")
}

func (estado estadoAprobado ) aceptarPorCS (registracion *Registracion) {
  fmt.Println("Se reenvía mail al alumno y al profesor")
  registracion.estado = estadoAprobadoID
}

func (estado estadoAprobado ) aceptarPorProfesor (registracion *Registracion) {
  fmt.Println("Se Registra el Tenant en Xubio y se avisa al alumno")
  registracion.estado = estadoConfirmadoID
}