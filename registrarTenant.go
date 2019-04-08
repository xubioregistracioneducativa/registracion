package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
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
	var parametros string = fmt.Sprintf("newDominio=%s&newUserName=%s&newPassword=%s&newNombre=%s&newPlan=inicial&newTelefono=%s&newCodigoPromocional=&soyContador=false&soyEmpresa=true&leiTerminos=true&g-recaptcha-response=03AOLTBLS8pKEId5Nr9_ZSiDl_3OuYqivWEbxtFgGaL-gHbpy6IURAso2e6-eWRYssPWYt4gZ8IT9clVrde8Xsm5j0X_HcCmQhlX34NyWWdUJRQd0OqQvJeZW-D04yBV7obhj7awOyzw8qNLRN3llSW8-87roJE9ixewym9FpYme30QLMxEKOzTdtIlPRjzOtrit4RyY8L4XiznRWhukL6T1fR3W2RB3FUxVMrcCJsgS1mDtILjz6Fc9b317d2oum-NN-FzfSKE8c9SokumoUDi3xCU3iBOO7KSHiWkVRMQX8CKwm3rIJ9irdZRvFZEknmOnvyO91QldaU&pais=AR&ip=190.55.124.44&ip_pais=AR&locationPathname=/NXV/inicio/createAccount.jsp", registracion.Email, registracion.Email, registracion.Clave, registracion.Nombre, registracion.Telefono)
	var urlencodedstring = []byte(parametros)
	req, err := http.NewRequest("POST", Direccion, bytes.NewBuffer(urlencodedstring))
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
		return err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil
}
