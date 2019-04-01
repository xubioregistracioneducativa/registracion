package main

import (
	"fmt"
	"errors"
)

type estadoAnulado struct {

}

func (estado estadoAnulado ) ingresarNuevosDatos (registracion *Registracion) error {
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

func (estado estadoAnulado ) rechazarPorCS (registracion *Registracion) error {
	return errors.New("Esta registracion fue anulada por nuestro equipo, por lo tanto no puede ser rechazada")
}

func (estado estadoAnulado ) aceptarPorCS (registracion *Registracion) error{
	return errors.New("Esta registracion fue anulada por nuestro equipo, por lo tanto no puede ser aceptada")
}

func (estado estadoAnulado ) anularPorCS (registracion *Registracion) error {
	return errors.New("Esta registracion fue anulada por nuestro equipo, por lo tanto no puede ser anulada")
}

func (estado estadoAnulado ) confirmarPorProfesor (registracion *Registracion) error{
	return errors.New("Esta registracion fue anulada por nuestro equipo, por lo tanto no puede ser confirmada")
}

func (estado estadoAnulado ) consultarEstado () string {
	return fmt.Sprint("Esta registracion fue anulada por nuestro equipo, puede volver a cargar una nueva registracion")
}

func (estado estadoAnulado ) vencerRegistracion (registracion *Registracion) error{
	return errors.New("No se puede vencer una registracion que no esta completa")
}