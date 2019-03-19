package main

import 	(
	"encoding/json"
	"net/http"
) 


func DecodificarRegistracion(request *http.Request) Registracion {

	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields ()
	var datosRegistracion Registracion

	//&nombre_var para decirle que es la var que no tiene datos y va a tener que rellenar
	var err = decoder.Decode(&datosRegistracion)

	if(err != nil){
		panic(err)
	}

	//Para cerrar la lectura de algo
	defer request.Body.Close()

	return datosRegistracion

}