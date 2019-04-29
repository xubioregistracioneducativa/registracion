package main

import (
	"fmt"
	"github.com/mattbaird/gochimp"
	"github.com/xubioregistracioneducativa/registracion/configuracion"
	"log"
)

//ENVIAR UN MAIL GENERAL (CANDIDATO A SER UN PACKAGE APARTE QUE SE ENCARGUE DE LOS MAILS)

func enviarOGuardarMail(email string, asunto string, html string) error {
	err := enviarMail(email, asunto, html)
	if err != nil{
		err = DBHelper.guardarMail(email, asunto, html)
		if err != nil {
			log.Panicln(err)
		}
	}
	return nil
}

func enviarMail(email string, asunto string, html string) error {
	//apiKey := os.Getenv("MANDRILL_KEY")
	apiKey := configuracion.MandrillKey()
	mandrillApi, err := gochimp.NewMandrill(apiKey)

	if err != nil {
		return err
	}

	var recipient gochimp.Recipient
	if configuracion.UsaEmailPrueba() {
		recipient = gochimp.Recipient{Email: configuracion.EmailPrueba()}
	} else {
		recipient = gochimp.Recipient{Email: email}
	}
	recipients := []gochimp.Recipient{recipient}

	message := gochimp.Message{
		Html:      html,
		Subject:   asunto,
		FromEmail: "noresponder@xubiomail.com",
		FromName:  "Xubio",
		To:        recipients,
	}

	if configuracion.EnviaEmails(){
		_, err = mandrillApi.MessageSend(message, false)
	} else {
		fmt.Println(asunto)
	}

	if err != nil {
		return err
	}

	return nil
}

//PARA LOS BOTONES DE LOS LINKS

func getButton(value string , link string ) string {
	var button = "<table>"
	button += "<tr>"
	button += "<a href=\"" + link + "\""
	button += "<td align=\"center\" width=\"250\" height=\"20\" bgcolor=\"#cf142b\""
	button += "style=\"font-weight: 700; font-family: 'Open Sans', Arial, Helvetica, sans-serif; "
	button += "font-size: 14px; margin-top: 5px;	padding: 7px 31px;	text-transform: uppercase;	"
	button += "display: block;	margin: 25px auto 0px auto;	text-transform: uppercase;  border-bottom: 3px solid #005e86;  "
	button += "border-top: 0;  border-right: 0;  border-left: 0;  text-decoration: none;  background-color: #f4f4f4; "
	button += "border-radius: 4px !important;  min-width: 55px;  color: #FFF;  background: #0193e1; "
	button += "background: -moz-linear-gradient(top, #00abeb 0%, #027cd8 100%); "
	button += "background: -webkit-gradient(linear, left top, left bottom, color-stop(0%,#2da9dc ), color-stop(100%,#027cd8)); "
	button += "background: -webkit-linear-gradient(top, #00abeb 0%,#027cd8 100%); background: -o-linear-gradient(top, #00abeb 0%,#027cd8 100%); "
	button += "background: -ms-linear-gradient(top, #00abeb 0%,#027cd8 100%); background: linear-gradient(to bottom, #00abeb 0%,#027cd8 100%); "
	button += "filter: progid:DXImageTransform.Microsoft.gradient( startColorstr='#00abeb ', endColorstr='#027cd8',GradientType=0 ); "
	button += "border: 1px solid #00abeb !important;  box-sizing: initial;\">"
	button += value + "</td>"
	button += "</a>"
	button += "</tr>"
	button += "</table>"

	button += "<div style='font-size: 11px; padding-top: 10px; padding-bottom: 10px;'>"
	button += "Si el botón no funciona copiá y pegá el siguiente link en tu navegador: <a href=\"" + link + "\">"+link+"</a>"
	button += "</div>"


	if !configuracion.EnviaEmails() {
		fmt.Println(link)
	}

	return button;
}

