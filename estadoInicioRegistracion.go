package main

import (
	"fmt"
	"errors"	
)

type estadoInicioRegistracion struct {

}

func (estado estadoInicioRegistracion ) ingresarNuevosDatos (registracion *Registracion) (string, error) {
	fmt.Println("Se guarda la Registracion")
	registracion.estado = estadoPendienteAprobacionID
	var err error
	if emailDeRegistroLibre((*registracion).Email) {
    	err = insertarNuevaRegistracion(registracion)
  	} else {
  		err = reingresarRegistracion(registracion)
  	}
	if err != nil {
		return "", err
	}

  	return "EXITO_INGRESAR", nil
}

func (estado estadoInicioRegistracion ) rechazarPorCS (registracion *Registracion) (string, error) {
  return "", errors.New("ERROR_ESTADOINICIO")
}

func (estado estadoInicioRegistracion ) aceptarPorCS (registracion *Registracion) (string, error){
  return "", errors.New("ERROR_ESTADOINICIO")
}

func (estado estadoInicioRegistracion ) anularPorCS (registracion *Registracion) (string, error) {
	return "", errors.New("ERROR_ESTADOINICIO")
}

func (estado estadoInicioRegistracion ) confirmarPorProfesor (registracion *Registracion) (string, error){
  return "", errors.New("ERROR_ESTADOINICIO")
}

func (estado estadoInicioRegistracion ) consultarEstado () string {
	return "ESTADO_INICIO"
}

func (estado estadoInicioRegistracion ) vencerRegistracion (registracion *Registracion) (string, error) {
	return "", errors.New("ERROR_ESTADOINICIO")
}