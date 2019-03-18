package main

import "fmt"

type estadoConfirmado struct {

}

func (estado estadoConfirmado ) ingresarNuevosDatos (registracion *Registracion) {
  fmt.Println("ERROR: Esta registracion ya fue confirmada")
}

func (estado estadoConfirmado ) rechazarPorCS (registracion *Registracion) {
  fmt.Println("ERROR: Esta registracion ya fue confirmada")
}

func (estado estadoConfirmado ) aceptarPorCS (registracion *Registracion) {
  fmt.Println("ERROR: Esta registracion ya fue confirmada")
}

func (estado estadoConfirmado ) aceptarPorProfesor (registracion *Registracion) {
  fmt.Println("ERROR: Esta registracion ya fue confirmada")

}