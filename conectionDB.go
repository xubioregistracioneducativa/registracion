package main

import (
  "database/sql"
	"fmt"
  _ "github.com/lib/pq"
  "errors"
	"github.com/xubioregistracioneducativa/registracion/configuracion"
	"log"
	"sync"
)

 type DataBase interface {
	crearTablas()
	insertarNuevaRegistracion(*Registracion) error
	updateRegistracion(*Registracion) error
	reingresarRegistracion(*Registracion) error
	emailDeRegistroLibre(string) bool
	obtenerRegistracionPorEmail(string) (Registracion, error)
	obtenerEstadoIDPorEmail(string) (estadoID, error)
	insertarNuevoLink(*Link) error
	eliminarLinksPorID(int) error
	obtenerLink(int, string) (Link, error)
}

type Postgres struct {

}
var DBHelper *Postgres
var onceDB sync.Once

func GetDBHelper() *Postgres {
	onceDB.Do( func() {
		DBHelper = &Postgres{}
	})
	return DBHelper
}

func (postgres Postgres) psqlInfo() string{
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		configuracion.DBHost(), configuracion.DBPort(), configuracion.DBUser(), configuracion.DBPassword(), configuracion.DBName())
}

func (postgres Postgres) crearTablas()  {
	postgres.crearTablaXRERegistracion()
	postgres.crearTablaXRELink()
	postgres.crearTablaMails()
}

func (postgres Postgres) crearTablaMails(){

	db, err := sql.Open("postgres", postgres.psqlInfo())
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	sqlStatement := `CREATE TABLE IF NOT EXISTS mailNoEnviado 
	(idmail SERIAL PRIMARY KEY, email VARCHAR, asunto VARCHAR, cuerpo VARCHAR);`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		log.Panic(err)
	}
}

func (postgres Postgres) crearTablaXRERegistracion(){
	
	db, err := sql.Open("postgres", postgres.psqlInfo())
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	sqlStatement := `CREATE TABLE IF NOT EXISTS XRERegistracion 
	(idregistracion SERIAL PRIMARY KEY, Nombre VARCHAR, Apellido VARCHAR, Email VARCHAR UNIQUE NOT NULL, Telefono VARCHAR,
		Carrera VARCHAR, Clave VARCHAR, NombreProfesor VARCHAR, ApellidoProfesor VARCHAR, EmailProfesor VARCHAR NOT NULL,
		Materia VARCHAR, Catedra VARCHAR, Facultad VARCHAR, Universidad VARCHAR, estado INT);`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		log.Panic(err)
	}
}

func (postgres Postgres) crearTablaXRELink(){

	db, err := sql.Open("postgres", postgres.psqlInfo())
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	sqlStatement := `CREATE TABLE IF NOT EXISTS XRELink 
	(IDRegistracion INT REFERENCES xreregistracion, accion VARCHAR, ValidationCode VARCHAR, primary key(accion, IDRegistracion));`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		log.Panic(err)
	}
}

func (postgres Postgres) insertarNuevaRegistracion(registracion *Registracion) error{

	db, err := sql.Open("postgres", postgres.psqlInfo())
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
        return err
    }

	sqlStatement := `INSERT INTO XRERegistracion (Nombre, Apellido, Email, Telefono, Carrera, Clave, NombreProfesor, ApellidoProfesor, EmailProfesor, Materia, Catedra, Facultad, Universidad, estado)
	 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING idregistracion`

	err = tx.QueryRow(sqlStatement, registracion.Nombre, registracion.Apellido, registracion.Email, registracion.Telefono, registracion.Carrera,
		registracion.Clave, registracion.NombreProfesor, registracion.ApellidoProfesor, registracion.EmailProfesor, registracion.Materia, registracion.Catedra,
		registracion.Facultad, registracion.Universidad, estadoPendienteAprobacionID).Scan(&registracion.IDRegistracion)
	if err != nil {
		log.Println(err)
		errRollback := tx.Rollback()
		if errRollback != nil {
			log.Println(errRollback)
			return errRollback
		}
	  	return err
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("New record ID is:", registracion.IDRegistracion)
	return nil
}

func (postgres Postgres) updateRegistracion(registracion *Registracion) error {

	db, err := sql.Open("postgres", postgres.psqlInfo())
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()


	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		errRollback := tx.Rollback()
		if errRollback != nil {
			return errRollback
		}
	  	return err
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("Se updateo el registro con ID:", registracion.IDRegistracion)

	return nil
}

func (postgres Postgres) reingresarRegistracion(registracion *Registracion) error {
	db, err := sql.Open("postgres", postgres.psqlInfo())
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer db.Close()

	sqlStatement := `select idregistracion from xreregistracion where email ilike $1;`
	
	row := db.QueryRow(sqlStatement, registracion.Email)
	err = row.Scan(&registracion.IDRegistracion)

	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("ERROR_DATABASE_REGISTRACIONEMAIL")
		}
		log.Println(err)
	  	return err
	}

	err = postgres.updateRegistracion(registracion)

	if err != nil {
		log.Println(err)
	  	return err
	}
	return nil
}

func (postgres Postgres) emailDeRegistroLibre(mail string) bool{
	db, err := sql.Open("postgres", postgres.psqlInfo())
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

    var cantidadDeCuentas int

	sqlStatement := `select count(idregistracion) as cantidadDeCuentas from xreregistracion where email ilike $1;`

	row := db.QueryRow(sqlStatement, mail)
	err = row.Scan(&cantidadDeCuentas)
	if err != nil {
	  log.Panic(err)
	}

	fmt.Println(cantidadDeCuentas)

	if cantidadDeCuentas > 0 {
		return false
	}
	return true

}

func (postgres Postgres) obtenerRegistracionPorEmail(email string) (Registracion, error){
	var registracion Registracion

	db, err := sql.Open("postgres", postgres.psqlInfo())
	if err != nil {
		log.Println(err)
		return registracion, err
	}
	defer db.Close()

	sqlStatement := `select * from xreregistracion where email = $1;`

	err = db.QueryRow(sqlStatement, email).Scan(&registracion.IDRegistracion, &registracion.Nombre, &registracion.Apellido, &registracion.Email, &registracion.Telefono, &registracion.Carrera, &registracion.Clave,
		&registracion.NombreProfesor, &registracion.ApellidoProfesor, &registracion.EmailProfesor, &registracion.Materia, &registracion.Catedra, &registracion.Facultad, &registracion.Universidad, &registracion.estado)

	if err != nil {
		if err == sql.ErrNoRows {
			return registracion , errors.New("ERROR_DATABASE_REGISTRACIONEMAIL")
		}
		log.Println(err)
		return registracion, err
	}

	return registracion, nil
}

func (postgres Postgres) obtenerEstadoIDPorEmail(email string) (estadoID, error){

	var estado estadoID

	db, err := sql.Open("postgres", postgres.psqlInfo())
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

func (postgres Postgres) insertarNuevoLink(link *Link) error {
	db, err := sql.Open("postgres", postgres.psqlInfo())
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	sqlStatement := `INSERT INTO XRELink (idregistracion, Accion, ValidationCode)
	 VALUES ($1, $2, $3)`

	_ , err = tx.Exec(sqlStatement, link.RegistracionID, link.Accion, link.ValidationCode)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return errRollback
		}
		log.Println(err)
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("Se inserto el link.")
	return nil
}

func (postgres Postgres) eliminarLinksPorID(IDRegistracion int) error {
	db, err := sql.Open("postgres", postgres.psqlInfo())
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	sqlStatement := `DELETE FROM XRELink WHERE IDRegistracion = $1`

	_ , err = tx.Exec(sqlStatement, IDRegistracion)
	if err != nil {
		log.Println(err)
		errRollback := tx.Rollback()
		if errRollback != nil {
			log.Println(errRollback)
			return errRollback
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (postgres Postgres) obtenerLink(idregistracion int, accion string) (Link, error) {
	var link Link

	db, err := sql.Open("postgres", postgres.psqlInfo())
	if err != nil {
		return link, err
	}
	defer db.Close()

	sqlStatement := `select * from xrelink where idregistracion = $1 and accion = $2;`

	err = db.QueryRow(sqlStatement, idregistracion, accion).Scan(&link.RegistracionID, &link.Accion, &link.ValidationCode)

	if err != nil {
		if err == sql.ErrNoRows {
			return link , errors.New("ERROR_DATABASE_LINK")
		}
		log.Println(err)
		return link, err
	}

	return link, nil
}

func (postgres Postgres) guardarMail(email string, asunto string, html string) error {
	db, err := sql.Open("postgres", postgres.psqlInfo())
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	sqlStatement := `INSERT INTO mailNoEnviado (email, asunto, cuerpo)
	 VALUES ($1, $2, $3)`

	_ , err = tx.Exec(sqlStatement, email, asunto, html)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return errRollback
		}
		log.Println(err)
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (postgres Postgres) obtenerMailsNoEnviados() ([]Mail, error){
	var mails []Mail;

	db, err := sql.Open("postgres", postgres.psqlInfo())

	if err != nil {
		log.Println(err)
		return mails, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM mailnoenviado")

	if err != nil {
		return mails, err
	}

	for rows.Next(){
		mail := Mail{}
		err = rows.Scan(&mail.IDMail, &mail.Email, &mail.Asunto, &mail.Cuerpo)
		if err != nil {
			return mails, err
		}
		mails = append(mails, mail)
	}

	return mails, err
}

func (postgres Postgres) eliminarMails (sliceIDMails []int) {
	db, err := sql.Open("postgres", postgres.psqlInfo())
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	sqlStatement := `DELETE FROM mailnoenviado WHERE idmail = $1`

	for  i := 0; i < len(sliceIDMails); i++ {
		_ , err = db.Exec(sqlStatement, sliceIDMails[i])
		if err != nil {
			log.Printf("No se pudo eliminar el mail con ID: %d ", sliceIDMails[i])
		}
	}
}