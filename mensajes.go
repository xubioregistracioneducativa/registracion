package main

import (
	"encoding/json"
	"github.com/xubioregistracioneducativa/registracion/configuracion"
	"os"
	"sync"
)

var mensajes map[string]string;
var once sync.Once


func cargarMensajes() {
	once.Do( func() {
		println("Se cargan los mensajes")
		mensajes = make(map[string]string)
		file, _ := os.Open(configuracion.NombreArchivoMensajes())
		defer file.Close()

		decoder := json.NewDecoder(file)
		err := decoder.Decode(&mensajes)
		if err != nil {
			panic(err)
		}
	})
}

func getMensaje(clave string) string {
	cargarMensajes()
	return mensajes[clave]
}
