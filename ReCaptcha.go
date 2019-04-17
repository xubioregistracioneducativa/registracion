package main

import (
	"encoding/json"
	"fmt"
	"github.com/xubioregistracioneducativa/registracion/configuracion"
	"log"
	"net/http"
)

func ValidateCaptcha (captchaValue string) bool {
	if(!configuracion.ValidaCaptcha()){
		return true
	}
	urlCaptcha := fmt.Sprintf("https://www.google.com/recaptcha/api/siteverify?secret=%s&response=%s", configuracion.SecretKeyReCaptcha(), captchaValue)

	resp, err := http.Get(urlCaptcha)

	defer resp.Body.Close()

	if err != nil {
		log.Panic(err)
	}

	var respuestas map[string]interface{};

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&respuestas)
	if err != nil {
		log.Panic(err)
	}

	validoCaptcha := respuestas["success"].(bool)

	fmt.Println(validoCaptcha)

	return validoCaptcha

}
