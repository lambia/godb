/* ToDo
	method to populate adb struct from a config file
	sql.Open working even without stuff (es. porta)
	"$variabile" invece di ""+variabile+""
	.porta è un int16 con casting invece di una stringa

	insert, update ecc.. return id not print
	-> sono methodi del db
	-> valori e campi = unico struct

	query può accettare query specifica o costruirla da parametri
	-> restituisce struct coi dati
	-> argomenti da usare come filtri (per query e per echo successivo)

	delete accetta struct {ID: 4}, restituisce bool

	sostituire nomi funzioni troppo generici
	mettere tutte queste funzioni in un package unico

	DoPanic non panic, ma log/echo
*/

package godb

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

type adb struct {
	meth string
	user string
	pass string
	prot string
	host string
	port string
	name string
	char string
}

func Connect() *sql.DB {
	thisdb := config()
	/* Both host AND port must be setted (or not). 
	if thisdb.host != "" && thisdb.port != "nil" {
		address := "("+thisdb.host+":"+thisdb.port+")"
	} else {
		address := ""
	}
	*/
	db, err := sql.Open(thisdb.meth, thisdb.user+":"+thisdb.pass+"@"+thisdb.prot+"("+thisdb.host+":"+thisdb.port+")/"+thisdb.name+"?"+thisdb.char)
	DoPanic(err)
	return db
}

//Ora scrive 3 su 3 valori dati, deve essere modulare
func Insert(db *sql.DB, tabella string, campi [3]string, valori [3]string) int64 {
	stmt, err := db.Prepare("INSERT "+tabella+" SET "+campi[0]+"=?,"+campi[1]+"=?,"+campi[2]+"=?")
	DoPanic(err)

	res, err := stmt.Exec(valori[0], valori[1], valori[2])
	DoPanic(err)

	id, err := res.LastInsertId()
	DoPanic(err)

	return id
}

//Ora aggiorna il primo dei 3 valori dati, deve essere modulare
func Update(db *sql.DB, tabella string, campi [3]string, valori [3]string, id int64) int64 {
	stmt, err := db.Prepare("update "+tabella+" set "+campi[0]+"=? where ID=?")
	DoPanic(err)

	res, err := stmt.Exec(valori[0], id)
	DoPanic(err)

	affect, err := res.RowsAffected()
	DoPanic(err)

	return affect
}

func QuerySelect(db *sql.DB, tabella string) {
	rows, err := db.Query("SELECT * FROM "+tabella+"")
	DoPanic(err)

	for rows.Next() {
		var uid int
		var nome string
		var username string
		var password string
		err = rows.Scan(&uid, &nome, &username, &password)
		DoPanic(err)
		fmt.Println(uid)
		fmt.Println(nome)
		fmt.Println(username)
		fmt.Println(password)
	}
}

func Delete(db *sql.DB, tabella string, id int64) {
	stmt, err := db.Prepare("delete from "+tabella+" where ID=?")
	DoPanic(err)

	res, err := stmt.Exec(id)
	DoPanic(err)

	affect, err := res.RowsAffected()
	DoPanic(err)

	fmt.Println(affect)
}

func Close(db *sql.DB) {
	db.Close()
}

func DoPanic(err error) {
	if err != nil {
		panic(err)
	}
}

//Prendere da file di config
func config() adb {
	thisdb := adb{
		meth: "mysql",
		user: "root",
		pass: "",
		prot: "tcp",
		host: "localhost",
		port: "3306",
		name: "go_test",
		char: "charset=utf8"}	
	return thisdb
}