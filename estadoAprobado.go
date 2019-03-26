package main

import (
	"errors"
	"fmt"
)

type estadoAprobado struct {

}

func (estado estadoAprobado ) ingresarNuevosDatos (registracion *Registracion) error {
	var err error
	fmt.Println("Se guarda la Registracion")
  	registracion.estado = estadoPendienteAprobacionID
  	err = reingresarRegistracion(registracion)
  	if err != nil {
  		return err
  	}
  	return nil
}

func (estado estadoAprobado ) rechazarPorCS (registracion *Registracion) error{
  	return errors.New("Esta registracion ya fue aceptada")
}

func (estado estadoAprobado ) aceptarPorCS (registracion *Registracion) error {
	  fmt.Println("Se reenvía mail al alumno y al profesor")
	  return nil
}

func (estado estadoAprobado ) confirmarPorProfesor (registracion *Registracion) error {
	  var err error
	  fmt.Println("Se Registra el Tenant en Xubio y se avisa al alumno")
	  registrarTenant(registracion)
	  registracion.estado = estadoConfirmadoID
	  err = updateRegistracion(registracion)
	  if err != nil {
			return err
	  }

	  return nil
}