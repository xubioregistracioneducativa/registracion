package configuracion

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

type ConfiguracionStruct struct {
	Puerto string				`json:"puerto"`
	UrlMono string				`json:"urlMono"`
	UrlStudent string			`json:"urlStudent"`
	EnviaEmails	bool			`json:"enviaEmails"`
	EmailPrueba string			`json:"emailPrueba"`
	ValidaCaptcha bool			`json:"validaCaptcha"`
	CodigoSalteaCaptcha string 	`json:"codigoSalteaCaptcha"`
}

var configuracion ConfiguracionStruct
var once sync.Once

func CargarConfiguracion() {
	once.Do( func() {
		var config_filename string
		entorno := os.Getenv("RECENV")
		switch entorno {
		case "D":
			config_filename = "config-desarrollo.json"
		case "P":
			config_filename = "config-production.json"
		case "T":
			config_filename = "config-testing.json"
		default:
			panic(errors.New("La variable de entorno no esta configurada correctamente"))
		}
		file, _ := os.Open(config_filename)
		defer file.Close()
		decoder := json.NewDecoder(file)

		err := decoder.Decode(&configuracion)
		if err != nil {
			panic(err)
		}
	})
}

func Puerto() string{
	return configuracion.Puerto
}

func UrlMono() string{
	return configuracion.UrlMono
}

func UrlStudent() string{
	return configuracion.UrlStudent
}

func EmailPrueba() string {
	return configuracion.EmailPrueba
}

func ValidaCaptcha() bool {
	return configuracion.ValidaCaptcha
}

func EnviaEmails() bool {
	return configuracion.EnviaEmails

}

func CodigoSalteaCaptcha() string {
	return configuracion.CodigoSalteaCaptcha
}