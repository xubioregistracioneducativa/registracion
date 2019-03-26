package main

import (
  "fmt"
)

type Url string

type Link struct {
  Url             Url    	    `json:"Url"`
  Input           string   	    `json:"Input"`
  ValidationCode  string        `json:"ValidationCode"`
  RegistracionID  int    		`json:"RegistracionID"`
}

func generarLinks(registracionID int) error {
  err := generarLink("AceptarCS", registracionID)
  err = generarLink("RechazarCS", registracionID)
  err = generarLink("ConfirmarProfesor", registracionID)
  if err != nil {
    return err
  }
  return nil

}

func generarLink(input string, registracionID int) error{

  validationCode, err := GenerateRandomString(32)
  if err != nil {
      return err
  }
  url := obtenerUrl(input, registracionID, validationCode)
  link := Link{url, input, validationCode, registracionID}
  err = insertarNuevoLink(&link)
  if err != nil {
    return err
  }
  return nil
}

func obtenerUrl(input string, registracionID int, validationCode string) Url {
  paginaBase := "http://localhost:8081" //Como obtener esto dinamicamente?
  return Url(fmt.Sprintf("%s/%s/%d/%s", paginaBase, input, registracionID, validationCode))
}

