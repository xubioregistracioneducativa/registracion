package main

import (
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
)

const (
  host     = "192.168.30.111"
  port     = 5432
  user     = "postgres"
  password = "Post66MM/"
  dbname   = "DES_MULTITENANT_AR_1"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func CrearTablaXRERegistracion(){
	
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `CREATE TABLE IF NOT EXISTS XRERegistracion 
	(idregistracion SERIAL PRIMARY KEY, Nombre VARCHAR, Apellido VARCHAR, Email VARCHAR UNIQUE NOT NULL, Telefono VARCHAR,
		Carrera VARCHAR, Clave VARCHAR, NombreProfesor VARCHAR, ApellidoProfesor VARCHAR, EmailProfesor VARCHAR UNIQUE NOT NULL,
		Materia VARCHAR, Catedra VARCHAR, Facultad VARCHAR, Universidad VARCHAR, estado INT);`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}

func insertarNuevaRegistracion(registracion Registracion) int{

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
        panic(err)
    }

	sqlStatement := `INSERT INTO XRERegistracion (Nombre, Apellido, Email, Telefono, Carrera, Clave, NombreProfesor, ApellidoProfesor, EmailProfesor, Materia, Catedra, Facultad, Universidad, estado)
	 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING idregistracion`
	id := 0

	err = tx.QueryRow(sqlStatement, registracion.Nombre, registracion.Apellido, registracion.Email, registracion.Telefono, registracion.Carrera,
		registracion.Clave, registracion.NombreProfesor, registracion.ApellidoProfesor, registracion.EmailProfesor, registracion.Materia, registracion.Catedra,
		registracion.Facultad, registracion.Universidad, estadoPendienteAprobacionID).Scan(&id)
	if err != nil {
	  tx.Rollback()
	  panic(err)
	}
	tx.Commit()
	fmt.Println("New record ID is:", id)
	return id
}

func updateRegistracion(registracion Registracion) {

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
        panic(err)
    }

	sqlStatement := `
	UPDATE XRERegistracion
	SET Nombre = $2, Apellido = $3, Email = $4, Telefono = $5, Carrera = $6, Clave = $7, NombreProfesor = $8, ApellidoProfesor = $9, EmailProfesor = $10, Materia = $11, Catedra = $12, Facultad = $13, Universidad = $14, estado = $15
	WHERE idregistracion = $1;`
	_, err = tx.Exec(sqlStatement, registracion.iDRegistracion ,registracion.Nombre, registracion.Apellido, registracion.Email, registracion.Telefono, registracion.Carrera,
		registracion.Clave, registracion.NombreProfesor, registracion.ApellidoProfesor, registracion.EmailProfesor, registracion.Materia, registracion.Catedra,
		registracion.Facultad, registracion.Universidad, registracion.estado)
	if err != nil {
	  tx.Rollback()
	  panic(err)
	}
	tx.Commit()
	fmt.Println("Se updateo el registro con ID:", registracion.iDRegistracion)
}

func mailDeRegistroLibre(mail string) bool{
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

    var cantidadDeCuentas int

	sqlStatement := `select count(idregistracion) as cantidadDeCuentas from xreregistracion where email ilike $1;`
	
	row := db.QueryRow(sqlStatement, mail)
	err = row.Scan(&cantidadDeCuentas)
	if err != nil {
	  panic(err)
	}

	fmt.Println(cantidadDeCuentas)

	if cantidadDeCuentas > 0 {
		return false
	}
	return true

}

func reingresarRegistracion(registracion Registracion) {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `select idregistracion from xreregistracion where email ilike $1;`
	
	row := db.QueryRow(sqlStatement, registracion.Email)
	err = row.Scan(&(registracion.iDRegistracion))
	if err != nil {
	  panic(err)
	}

	updateRegistracion(registracion)
}

func obtenerRegistracionPorID(registracionID int) Registracion {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `select * from xreregistracion where iDRegistracion ilike $1;`
	
	row := db.QueryRow(sqlStatement, registracionID)
	if err != nil {
	  panic(err)
	}

	var registracion Registracion

	registracion.iDRegistracion = registracionID

	row.Scan(&(registracion.Nombre))
	row.Scan(&(registracion.Apellido))
	row.Scan(&(registracion.Email))
	row.Scan(&(registracion.Telefono))
	row.Scan(&(registracion.Carrera))
	row.Scan(&(registracion.Clave))
	row.Scan(&(registracion.NombreProfesor))
	row.Scan(&(registracion.ApellidoProfesor))
	row.Scan(&(registracion.EmailProfesor))
	row.Scan(&(registracion.Materia))
	row.Scan(&(registracion.Catedra))
	row.Scan(&(registracion.Facultad))
	row.Scan(&(registracion.Universidad))
	row.Scan(&(registracion.estado))
}