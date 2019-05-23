package main

import (
  "errors"
  "fmt"
  "github.com/xubioregistracioneducativa/registracion/configuracion"
  "log"
)

type Link struct {
  RegistracionID    int
  Accion            string
  ValidationCode    string

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
    log.Println(err)
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
  err := GetDBHelper().insertarNuevoLink(&link)
  if err != nil {
    return err
  }
  return nil
}



func validarLink(registracionID int, accion string, validationCode string) error {
  link, err := GetDBHelper().obtenerLink(registracionID, accion)
  if err != nil {
    return err
  }
  if link.ValidationCode != validationCode {
    err = errors.New("ERROR_LINK_VALIDACION")
    log.Println(err)
    return err
  }
  return nil
}

func obtenerUrlLink(link *Link, email string) string {
  return fmt.Sprintf("%s%s/%s/%s", configuracion.UrlStudent(), link.Accion, email, link.ValidationCode)
}

func obtenerUrlXubioNuevaRegistracion() string {
  return fmt.Sprintf("%s%s", configuracion.UrlMono(), configuracion.PathEstudiantes())
}
