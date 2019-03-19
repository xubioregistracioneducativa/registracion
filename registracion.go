package main

type Registracion struct {
  iDRegistracion int    	`json:"IDRegistracion"`
  Nombre       string   	`json:"Nombre"`
  Apellido     string 		`json:"Apellido"`
  Email        string 		`json:"Email"`
  Telefono     string 		`json:"Telefono"`
  Carrera      string 		`json:"Carrera"`
  Clave        string 		`json:"Clave"`
  NombreProfesor string 	`json:"NombreProfesor"`
  ApellidoProfesor string 	`json:"ApellidoProfesor"`
  EmailProfesor string 		`json:"EmailProfesor"`
  Materia string  			`json:"Materia"`
  Catedra string 			`json:"Catedra"`
  Facultad string 			`json:"Facultad"`
  Universidad string 		`json:"Universidad"`
  estado estadoID    			
  //estado  estado `gorm:"polymorphic:Owner;"`
}