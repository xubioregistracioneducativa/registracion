package main

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/xubioregistracioneducativa/registracion/configuracion"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func registrarTenant(registracion *Registracion) error {
	//Direccion := "http://localhost:8080/NOSecure/PUReceiveFeed"
	Direccion := "https://localhost:8443/CreateCuenta"
	//Direccion := "https://xubio.com/CreateCuenta"
	fmt.Println("URL:>", Direccion)
	//var valores = url.Values{"key": {"Value"}, "id": {"123"}}
	//resp, err := http.PostForm(Direccion, valores)
	//var parametros string = fmt.Sprintf("newDominio=%s&newUserName=%s&newPassword=%s&newNombre=%s&newPlan=inicial&newTelefono=%s&newCodigoPromocional=&codigoSalteaCaptcha=ELCODIGOULTRASECRETO&soyContador=false&soyEmpresa=true&soyEstudiante=true&leiTerminos=true&g-recaptcha-response=&pais=AR&ip=100.153.156.32&ip_pais=US&locationPathname=/NXV/inicio/createAccount.jsp",
	//	"cuenta2@cuenta.com", "cuenta2@cuenta.com", "asd12345", "12345646", "12315616")
	var parametros string = fmt.Sprintf("newDominio=%s" +
		"&newUserName=%s" +
		"&newPassword=%s" +
		"&newNombre=%s" +
		"&newPlan=inicial" +
		"&newTelefono=%s" +
		"&newCodigoPromocional=" +
		"&codigoSalteaCaptcha=%s" +
		"&soyContador=false" +
		"&soyEmpresa=true" +
		"&soyEstudiante=true" +
		"&leiTerminos=true" +
		"&g-recaptcha-response=" +
		"&pais=AR" +
		"&ip=190.55.124.44" +
		"&ip_pais=AR" +
		"&locationPathname=/NXV/inicio/createAccount.jsp",
		registracion.Email,
		registracion.Email,
		registracion.Clave,
		registracion.Nombre,
		registracion.Telefono,
		configuracion.CodigoSalteaCaptcha())
	var urlencodedstring = []byte(parametros)
	req, err := http.NewRequest("POST", Direccion, bytes.NewBuffer(urlencodedstring))
	if err != nil {
		log.Panic(err)
	}
	//req.Header.Set("Accept", "*/*")
	//req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	//req.Header.Set("Accept-Language", "es-419,es;q=0.9")
	//req.Header.Set("Cache-Control", "no-cache")
	//req.Header.Set("Connection", "keep-alive")
	//req.Header.Set("Content-Length", "")

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   120 * time.Second,
				KeepAlive: 120 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("response Body: "+ string(body) )
	err = buscarError(string(body))
	if err != nil{
		log.Panic(err)
	}
	// BODY ESPERADO "<createCuenta><errorNum>0</errorNum><mensaje>La creación del usuario se realizó correctamente</mensaje><OK>1</OK></createCuenta>"
	// BODY ERROR response Body: <createCuenta><errorNum>20</errorNum><mensaje>Ya existe un usuario con el mail ingresado. Si no recuerda la clave utilice la opción de "Olvidé mi clave" en la pantalla de ingreso.</mensaje><OK>0</OK></createCuenta>
	return nil
}

func buscarError(body string) error {


	substrings := strings.Split(
		strings.ReplaceAll(
			strings.ReplaceAll(body, "</OK>", "#"),
			"<OK>", "#"),
			"#")
	fmt.Println(substrings)
	Ok, err := strconv.Atoi(substrings[1])
	if err != nil{
		return err
	}
	if Ok == 0 {
		substrings = strings.Split(strings.ReplaceAll(
			strings.ReplaceAll(
				body, "<mensaje>", "#"),
				"</mensaje>", "#"),
				"#")
		return errors.New(substrings[1])
	}
	return nil
}
