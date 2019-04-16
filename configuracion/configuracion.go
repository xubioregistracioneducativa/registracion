package configuracion

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"sync"
)

type ConfiguracionStruct struct {
	Puerto 					string		`json:"puerto"`
	UrlMono 				string		`json:"urlMono"`
	UrlStudent 				string		`json:"urlStudent"`
	EnviaEmails				bool		`json:"enviaEmails"`
	EmailPrueba 			string		`json:"emailPrueba"`
	ValidaCaptcha 			bool		`json:"validaCaptcha"`
	CodigoSalteaCaptcha 	string 		`json:"codigoSalteaCaptcha"`
	NombreArchivoMensajes 	string		`json:"nombreArchivoMensajes"`
	DBHost					string		`json:"dbhost"`
	DBPort					int			`json:"dbport"`
	DBUser					string		`json:"dbuser"`
	DBPassword				string		`json:"dbpassword"`
	DBName					string		`json:"dbname"`
	UrlExito				string		`json:"urlExito"`
	UrlError				string		`json:"urlError"`
	UrlConsultarEstado		string		`json:"urlConsultarEstado"`
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
			log.Panic(errors.New("La variable de entorno no esta configurada correctamente"))
		}
		file, _ := os.Open(config_filename)
		defer file.Close()
		decoder := json.NewDecoder(file)

		err := decoder.Decode(&configuracion)
		if err != nil {
			log.Panic(err)
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

func NombreArchivoMensajes() string {
	return configuracion.NombreArchivoMensajes
}

func DBHost() string {
	return configuracion.DBHost
}

func DBPort() int {
	return configuracion.DBPort
}

func DBUser() string{
	return configuracion.DBUser
}

func DBPassword() string{
	return configuracion.DBPassword
}

func DBName() string  {
	return configuracion.DBName
}

func UrlExito() string {
	return configuracion.UrlExito
}

func UrlError() string {
	return configuracion.UrlError
}

func UrlConsultarEstado() string {
	return configuracion.UrlConsultarEstado
}
