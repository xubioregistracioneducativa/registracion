package main

type Registracion struct {
  IDRegistracion int
  nombre       string
  apellido     string
  email        string
  telefono     string
  carrera      string
  clave        string
  nombreProfesor string
  apellidoProfesor string
  emailProfesor string
  materia string
  catedra string
  facultad string
  universidad string
  estado int
  //estado  Estado `gorm:"polymorphic:Owner;"`
}