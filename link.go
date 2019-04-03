package main

import (
  "fmt"
)

type Link struct {
  Url             string    	    `json:"Url"`
  Input           string   	    `json:"Input"`
  ValidationCode  string        `json:"ValidationCode"`
  RegistracionID  int    		`json:"RegistracionID"`
}

func generarLinks(registracionID int, email string) error {
  err := generarLinkRandom("AceptarCS", registracionID)
  err = generarLinkRandom("RechazarCS", registracionID)
  err = generarLinkRandom("ConfirmarProfesor", registracionID)
  err = generarLinkRandom("ConsultarEstado", registracionID)
  err = generarLinkRandom("AnularCS", registracionID)
  err = generarLinkEmail("VencerRegistracion", email, "Mkj0WEW1iWJvJGKWXAWG8HkWng4R0maRwxNl2_QOpu8=", registracionID)
  if err != nil {
    return err
  }
  return nil

}

func generarLinkRandom(input string, registracionID int) error {
  validationCode, err := GenerateRandomString(32)
  if err != nil {
    return err
  }
  err = generarLink(input, registracionID, validationCode)
  if err != nil {
    return err
  }
  return nil
}

func generarLink(input string, registracionID int, validationCode string) error{

  url := obtenerUrl(input, fmt.Sprint(registracionID), validationCode)
  link := Link{url, input, validationCode, registracionID}
  err := insertarNuevoLink(&link)
  if err != nil {
    return err
  }
  return nil
}

func generarLinkEmail(input string, email string, validationCode string, registracionID int) error{

  url := obtenerUrl(input, email, validationCode)
  link := Link{url, input, validationCode, registracionID}
  err := insertarNuevoLink(&link)
  if err != nil {
    return err
  }
  return nil
}

func obtenerUrl(input string, identificador string, validationCode string) string {
  paginaBase := "http://localhost:8081" //Como obtener esto dinamicamente?
  return fmt.Sprintf("%s/%s/%s/%s", paginaBase, input, identificador, validationCode)
}

func obtenerUrlLink(link *Link) string {
  paginaBase := "http://localhost:8081" //Como obtener esto dinamicamente?
  return fmt.Sprintf("%s/%s/%d/%s", paginaBase, link.Input, link.RegistracionID, link.ValidationCode)
}

