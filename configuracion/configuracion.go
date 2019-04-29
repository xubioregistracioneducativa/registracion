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
	UrlMono 				string		`json:"urlmono"`
	UrlStudent 				string		`json:"urlstudent"`
	EnviaEmails				bool		`json:"enviaemails"`
	UsaEmailPrueba			bool		`json:"usaemailprueba"`
	EmailPrueba 			string		`json:"emailprueba"`
	ValidaCaptcha 			bool		`json:"validacaptcha"`
	CodigoSalteaCaptcha 	string 		`json:"codigosalteacaptcha"`
	NombreArchivoMensajes 	string		`json:"nombrearchivomensajes"`
	DBHost					string		`json:"dbhost"`
	DBPort					int			`json:"dbport"`
	DBUser					string		`json:"dbuser"`
	DBPassword				string		`json:"dbpassword"`
	DBName					string		`json:"dbname"`
	PathExito				string		`json:"pathexito"`
	PathError				string		`json:"patherror"`
	PathConsultarEstado		string		`json:"pathconsultarEstado"`
	PathEstudiantes			string		`json:"pathestudientes"`
	SecretKeyReCaptcha		string		`json:"secretkeyrecaptcha"`
	EmailCS					string		`json:"emailcs"`
	MandrillKey				string		`json:"mandrillkey"`
	TiempoMailNoEnviadoSeg 	int 		`json:"tiempomailnoenviadoseg"`
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

func PathExito() string {
	return configuracion.PathExito
}

func PathError() string {
	return configuracion.PathError
}

func PathConsultarEstado() string {
	return configuracion.PathConsultarEstado
}

func PathEstudiantes() string {
	return configuracion.PathEstudiantes
}

func SecretKeyReCaptcha() string {
	return configuracion.SecretKeyReCaptcha
}

func EmailCS() string {
	return configuracion.EmailCS
}

func UsaEmailPrueba() bool{
	return configuracion.UsaEmailPrueba
}

func MandrillKey() string {
	return configuracion.MandrillKey
}

func TiempoMailNoEnviadoSeg() int {
	return configuracion.TiempoMailNoEnviadoSeg
}