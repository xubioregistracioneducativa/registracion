package main

type DatosRegistracion struct {
  LeiTerminos bool `json:"leiterminos"`
  CaptchaValue string `json:"captchavalue"`
  Registracion Registracion `json:"registracion"`
}

type Registracion struct {
  IDRegistracion int    	`json:"idregistracion"`
  Nombre       string   	`json:"nombre"`
  Apellido     string 		`json:"apellido"`
  Email        string 		`json:"email"`
  Telefono     string 		`json:"telefono"`
  Carrera      string 		`json:"carrera"`
  Clave        string 		`json:"clave"`
  NombreProfesor string 	`json:"nombreprofesor"`
  ApellidoProfesor string 	`json:"apellidoprofesor"`
  EmailProfesor string 		`json:"emailprofesor"`
  Materia string  			`json:"materia"`
  Catedra string 			`json:"catedra"`
  Facultad string 			`json:"facultad"`
  Universidad string 		`json:"universidad"`
  Utm_source    string      `json:"utm_source"`
  Utm_medium    string      `json:"utm_medium"`
  Utm_term      string      `json:"utm_term"`
  Utm_content   string      `json:"utm_content"`
  Utm_campaign  string      `json:"utm_campaign"`
  estado estadoID    			
  //estado  estado `gorm:"polymorphic:Owner;"`
}