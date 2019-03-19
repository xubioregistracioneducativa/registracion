package main

import "fmt"

type estadoInicioRegistracion struct {

}

func (estado estadoInicioRegistracion ) ingresarNuevosDatos (registracion *Registracion) {
	fmt.Println("Se guarda la Registracion")
	registracion.estado = estadoPendienteAprobacionID
	if mailDeRegistroLibre((*registracion).Email) {
    	insertarNuevaRegistracion(*registracion)
  	} else {
  		reingresarRegistracion(*registracion)
  	}
}

func (estado estadoInicioRegistracion ) rechazarPorCS (registracion *Registracion) {
  fmt.Println("ERROR: Esta registracion aun no fue completada o fue rechazada anteriormente")
}

func (estado estadoInicioRegistracion ) aceptarPorCS (registracion *Registracion) {
  fmt.Println("ERROR: Esta registracion aun no fue completada o fue rechazada anteriormente")
}

func (estado estadoInicioRegistracion ) confirmarPorProfesor (registracion *Registracion) {
  fmt.Println("ERROR: Esta registracion aun no fue completada o fue rechazada anteriormente")
}