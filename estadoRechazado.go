package main

import (
	"errors"
	"fmt"
)

type estadoRechazado struct {

}

func (estado estadoRechazado ) ingresarNuevosDatos (registracion *Registracion) error {
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

func (estado estadoRechazado ) rechazarPorCS (registracion *Registracion) error {
	return errors.New("Esta registracion ya fue rechazada anteriormente, por lo tanto no puede ser rechazada")
}

func (estado estadoRechazado ) aceptarPorCS (registracion *Registracion) error{
	return errors.New("Esta registracion ya fue rechazada anteriormente, por lo tanto no puede ser aceptada")
}

func (estado estadoRechazado ) anularPorCS (registracion *Registracion) error {
	return errors.New("Esta registracion ya fue rechazada anteriormente, por lo tanto no puede ser anulada")
}

func (estado estadoRechazado ) confirmarPorProfesor (registracion *Registracion) error{
	return errors.New("Esta registracion ya fue rechazada anteriormente, por lo tanto no puede ser confirmada")
}

func (estado estadoRechazado ) consultarEstado () string {
	return fmt.Sprint("Esta registracion fue rechazada, puede volver a cargar una nueva registracion")
}

func (estado estadoRechazado ) vencerRegistracion (registracion *Registracion) error{
	return errors.New("No se puede vencer una registracion que no esta completa")
}