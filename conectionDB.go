package main

import (
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
  "errors"
)

const (
  host     = "192.168.30.111"
  port     = 5432
  user     = "postgres"
  password = "Post66MM/"
  dbname   = "faf_multitenant_go"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func CrearTablas()  {
	CrearTablaXRERegistracion()
	CrearTablaXRELink()
}

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

func CrearTablaXRELink(){

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `CREATE TABLE IF NOT EXISTS XRELink 
	(Url VARCHAR PRIMARY KEY, input VARCHAR, ValidationCode VARCHAR, IDRegistracion INT REFERENCES xreregistracion);`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}


func insertarNuevaRegistracion(registracion *Registracion) error{

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
        return err
    }

	sqlStatement := `INSERT INTO XRERegistracion (Nombre, Apellido, Email, Telefono, Carrera, Clave, NombreProfesor, ApellidoProfesor, EmailProfesor, Materia, Catedra, Facultad, Universidad, estado)
	 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING idregistracion`

	err = tx.QueryRow(sqlStatement, registracion.Nombre, registracion.Apellido, registracion.Email, registracion.Telefono, registracion.Carrera,
		registracion.Clave, registracion.NombreProfesor, registracion.ApellidoProfesor, registracion.EmailProfesor, registracion.Materia, registracion.Catedra,
		registracion.Facultad, registracion.Universidad, estadoPendienteAprobacionID).Scan(&registracion.IDRegistracion)
	if err != nil {
	  tx.Rollback()
	  return err
	}
	tx.Commit()
	fmt.Println("New record ID is:", registracion.IDRegistracion)
	return nil
}

func updateRegistracion(registracion *Registracion) error {

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	defer db.Close()


	tx, err := db.Begin()
	if err != nil {
        return err
    }

	sqlStatement := `
	UPDATE XRERegistracion
	SET Nombre = $2, Apellido = $3, Email = $4, Telefono = $5, Carrera = $6, Clave = $7, NombreProfesor = $8, ApellidoProfesor = $9, EmailProfesor = $10, Materia = $11, Catedra = $12, Facultad = $13, Universidad = $14, estado = $15
	WHERE idregistracion = $1;`
	_, err = tx.Exec(sqlStatement, registracion.IDRegistracion ,registracion.Nombre, registracion.Apellido, registracion.Email, registracion.Telefono, registracion.Carrera,
		registracion.Clave, registracion.NombreProfesor, registracion.ApellidoProfesor, registracion.EmailProfesor, registracion.Materia, registracion.Catedra,
		registracion.Facultad, registracion.Universidad, registracion.estado)
	if err != nil {
	  tx.Rollback()
	  return err
	}
	tx.Commit()
	fmt.Println("Se updateo el registro con ID:", registracion.IDRegistracion)

	return nil
}

func reingresarRegistracion(registracion *Registracion) error {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `select idregistracion from xreregistracion where email ilike $1;`
	
	row := db.QueryRow(sqlStatement, registracion.Email)
	err = row.Scan(&registracion.IDRegistracion)

	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("SQL: No se encontraron registraciones con ese mail")
		}
	  	return err
	}

	err = updateRegistracion(registracion)

	if err != nil {
	  	return err
	}
	return nil
}

func emailDeRegistroLibre(mail string) bool{
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


func obtenerRegistracionPorID(registracionID int) (Registracion, error){
	var registracion Registracion

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `select * from xreregistracion where iDRegistracion = $1;`

	err = db.QueryRow(sqlStatement, registracionID).Scan(&registracion.IDRegistracion, &registracion.Nombre, &registracion.Apellido, &registracion.Email, &registracion.Telefono, &registracion.Carrera, &registracion.Clave, 
		&registracion.NombreProfesor, &registracion.ApellidoProfesor, &registracion.EmailProfesor, &registracion.Materia, &registracion.Catedra, &registracion.Facultad, &registracion.Universidad, &registracion.estado)

	if err != nil {
		if err == sql.ErrNoRows {
			return registracion , errors.New("SQL: No se encontraron registraciones con ese ID")
		}
	  	return registracion, err
	}

	return registracion, nil
}

func obtenerRegistracionPorEmail(email string) (Registracion, error){
	var registracion Registracion

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `select * from xreregistracion where email = $1;`

	err = db.QueryRow(sqlStatement, email).Scan(&registracion.IDRegistracion, &registracion.Nombre, &registracion.Apellido, &registracion.Email, &registracion.Telefono, &registracion.Carrera, &registracion.Clave,
		&registracion.NombreProfesor, &registracion.ApellidoProfesor, &registracion.EmailProfesor, &registracion.Materia, &registracion.Catedra, &registracion.Facultad, &registracion.Universidad, &registracion.estado)

	if err != nil {
		if err == sql.ErrNoRows {
			return registracion , errors.New("SQL: No se encontraron registraciones con ese email")
		}
		return registracion, err
	}

	return registracion, nil
}

func obtenerEstadoIDPorEmail(email string) (estadoID, error){

	var estado estadoID

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	sqlStatement := `select estado from xreregistracion where email = $1;`

	err = db.QueryRow(sqlStatement, email).Scan(&estado)

	if err != nil {
		if err == sql.ErrNoRows {
			return estadoInicioRegistracionID, nil
		}
	  	return 0, err
	}

	return estado, nil
}

func insertarNuevoLink(link *Link) error {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	sqlStatement := `INSERT INTO XRELink (Url, Input, ValidationCode, idregistracion)
	 VALUES ($1, $2, $3, $4)`

	_ , err = tx.Exec(sqlStatement, link.Url, link.Input, link.ValidationCode, link.RegistracionID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	fmt.Println("Se inserto el link: ", link.Url)
	return nil
}

func eliminarLink(url Url) error {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	sqlStatement := `DELETE FROM XRELink WHERE url = $1`

	_ , err = tx.Exec(sqlStatement, url)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	fmt.Println("Se elimino el link: ", url)
	return nil
}

func obtenerLinkPorUrl(url Url) (Link, error){
	var link Link

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return link, err
	}
	defer db.Close()

	sqlStatement := `select * from xrelink where url = $1;`

	err = db.QueryRow(sqlStatement, url).Scan(&link.Url,&link.Input, &link.ValidationCode, &link.RegistracionID)

	if err != nil {
		if err == sql.ErrNoRows {
			return link , errors.New("El link no es valido")
		}
		return link, err
	}

	return link, nil
}

func eliminarLinksPorID(IDRegistracion int) error {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	sqlStatement := `DELETE FROM XRELink WHERE IDRegistracion = $1`

	_ , err = tx.Exec(sqlStatement, IDRegistracion)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
