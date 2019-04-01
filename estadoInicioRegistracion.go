package main

import (
	"fmt"
	"errors"	
)

type estadoInicioRegistracion struct {

}

func (estado estadoInicioRegistracion ) ingresarNuevosDatos (registracion *Registracion) error {
	fmt.Println("Se guarda la Registracion")
	registracion.estado = estadoPendienteAprobacionID
	var err error
	if emailDeRegistroLibre((*registracion).Email) {
    	err = insertarNuevaRegistracion(registracion)
  	} else {
  		err = reingresarRegistracion(registracion)
  	}
	if err != nil {
		return err
	}

  	return nil
}

func (estado estadoInicioRegistracion ) rechazarPorCS (registracion *Registracion) error {
  return errors.New("Esta registracion aun no fue completada o fue rechazada anteriormente")
}

func (estado estadoInicioRegistracion ) aceptarPorCS (registracion *Registracion) error{
  return errors.New("Esta registracion aun no fue completada o fue rechazada anteriormente")
}

func (estado estadoInicioRegistracion ) anularPorCS (registracion *Registracion) error {
	return errors.New("Esta registracion ya esta vencida, por lo tanto no puede anularse")
}

func (estado estadoInicioRegistracion ) confirmarPorProfesor (registracion *Registracion) error{
  return errors.New("Esta registracion aun no fue completada o fue rechazada anteriormente")
}

func (estado estadoInicioRegistracion ) consultarEstado () string {
	return fmt.Sprint("Esta registracion se encuentra en el inicio de registracion")
}

func (estado estadoInicioRegistracion ) vencerRegistracion (registracion *Registracion) error{
	return errors.New("No se puede vencer una registracion que no esta completa")
}