package main

import (
  "errors"
  "fmt"
  "github.com/xubioregistracioneducativa/registracion/configuracion"
)

type Link struct {
  RegistracionID  int    		`json:"RegistracionID"`
  Accion           string   	`json:"Accion"`
  ValidationCode  string        `json:"ValidationCode"`

}

func generarLinks(registracionID int) error {
  err := generarLinkRandom(registracionID, "AceptarCS")
  err = generarLinkRandom(registracionID, "RechazarCS")
  err = generarLinkRandom(registracionID, "ConfirmarProfesor")
  err = generarLinkRandom(registracionID, "ConsultarEstado")
  err = generarLinkRandom(registracionID, "AnularCS")
  err = generarLink(registracionID, "VencerRegistracion", "Mkj0WEW1iWJvJGKWXAWG8HkWng4R0maRwxNl2_QOpu8=")
  if err != nil {
    return err
  }
  return nil

}

func generarLinkRandom(registracionID int, accion string) error {
  validationCode, err := GenerateRandomString(32)
  if err != nil {
    return err
  }
  err = generarLink(registracionID, accion, validationCode)
  if err != nil {
    return err
  }
  return nil
}

func generarLink(registracionID int, accion string, validationCode string) error{

  link := Link{registracionID, accion, validationCode}
  err := insertarNuevoLink(&link)
  if err != nil {
    return err
  }
  return nil
}



func validarLink(registracionID int, accion string, validationCode string) error {
  link, err := obtenerLink(registracionID, accion)
  if err != nil {
    return err
  }
  if link.ValidationCode != validationCode {
    return errors.New("El codigo de validacion es incorrecto")
  }
  return nil
}

func obtenerUrlLink(link *Link, email string) string {
  return fmt.Sprintf("%s/%s/%s/%s", configuracion.UrlStudent(), link.Accion, email, link.ValidationCode)
}

func obtenerUrlXubioNuevaRegistracion() string {
  return fmt.Sprintf("%s/%s", configuracion.UrlStudent(), "NuevaRegistracion")
}

func obtenerUrlRegistracionIngresada() string {
  return fmt.Sprintf("%s/%s", configuracion.UrlStudent(), "SolicitudIngresada")
}