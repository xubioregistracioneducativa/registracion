package main

import (
	"log"
	"time"
)

type Mail struct {
	IDMail		int
	Email		string
	Asunto		string
	Cuerpo		string
}

func EnviarMailsNoEnviados() {
	for{
		time.Sleep(1 * time.Hour)
		mails, err := DBHelper.obtenerMailsNoEnviados()
		if err != nil {
			continue
		}
		var sliceIDMails []int;

		for i := 0; i < len(mails); i++{
			err = enviarMail(mails[i].Email, mails[i].Asunto, mails[i].Cuerpo)
			if err != nil {
				log.Printf("No se envio el mail: %d\n", mails[i].IDMail)
			} else {
				sliceIDMails = append(sliceIDMails, mails[i].IDMail)
			}
		}
		DBHelper.eliminarMails(sliceIDMails)
	}
}