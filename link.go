package main

import (
  "fmt"
)

type Link struct {
  Url             string    	`json:"Url"`
  Input           string   	    `json:"Input"`
  ValidationCode  string        `json:"ValidationCode"`
  RegistracionID  int    		`json:"RegistracionID"`
}

func generarLinks(registracionID int){
  generarLink("AceptarCS", registracionID)
  generarLink("RechazarCS", registracionID)
  generarLink("ConfirmarProfesor", registracionID)
}

func generarLink(input string, registracionID int){
  paginaBase := "http://localhost:8081" //Como obtener esto dinamicamente?
  validationCode, err := GenerateRandomString(32)
  if err != nil {
        panic(err)
  }
  url := fmt.Sprintf("%s/%s/%s/%d", paginaBase, validationCode, input, registracionID)
  link := Link{url, input, validationCode, registracionID}
  fmt.Println(link)
  //link.guardarse()
}

