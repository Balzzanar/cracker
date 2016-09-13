package main

import (
    _ "github.com/mattn/go-sqlite3"
    "database/sql"
    "fmt"
)

/* -- Constants -- */
const TABLE_WPA = `create table if not exists wpa (id integer not null primary key, name text, bssid varchar(30));`
const TABLE_WORDLISTS = `create table if not exists wordlists (id integer not null primary key, name text, size varchar(10), avg_run int);`
const TABLE_RUNS = `create table if not exists runs (id_wpa int, id_wordlist, result text, time int, started int);`
const DB_FILE_NAME = "./list.db"


type Wordlist struct {
	id int
	name string
	size string
	avg_run int
}

type Wpa struct {
	id int
	name string
	bssid string
}

type DBHandler struct {
	db *sql.DB
}


/**
 * Opens a connection to the databasefile, creates one if it does not exits
 * 
 * @name Init
 */
func (this *DBHandler) Init() {
	var derr error
	this.db, derr = sql.Open("sqlite3", DB_FILE_NAME)
	if derr != nil {
		fmt.Println(derr)
	}
	this.createNewTable(TABLE_WPA)
	this.createNewTable(TABLE_WORDLISTS)
	this.createNewTable(TABLE_RUNS)
}


/**
 * Closes the connection to the databasefile
 * 
 * @name Close
 */
func (this *DBHandler) Close() {
	this.db.Close()
}


/**
 * Stores a wordlist to the databasefile
 * 
 * @name StoreWordlist
 */
func (this *DBHandler) StoreWordlist(wordlist *Wordlist) {
	if ! this.isWordlistUniqe(wordlist.name) {
		log.Info(fmt.Sprintf("Wordlist (%s) already exists, ignoring.", wordlist.name))
		return
	}

	tx, err := this.db.Begin()
	if err != nil {
		fmt.Println(err)
	}
	stmt, err := tx.Prepare("insert into wordlists(id, name, size, avg_run) values(null, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(wordlist.name, wordlist.size, wordlist.avg_run)
	if err != nil {
		fmt.Println(err)
	}
	tx.Commit()
}


/**
 * Gets a list with all the knows wpas.
 * 
 * @name GetAllWpa
 * @return []Wpa
 */
func (this *DBHandler) GetAllWpa() []Wpa {
	listwpa := []Wpa{}
	rows, err := this.db.Query("select * from wpa")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var wpa Wpa
		err = rows.Scan(&wpa.id, &wpa.name, &wpa.bssid)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("ID: %d\n", wpa.id)
		listwpa = append(listwpa, wpa)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	return listwpa
}


/**
 * Stores a new Wpa to the database file
 * Will enforce uniqeness on name.
 * 
 * @name StoreWpa
 */
func (this *DBHandler) StoreWpa(wpa *Wpa) {
	if ! this.isWpaUniqe(wpa.name) {
		log.Info(fmt.Sprintf("Wpa (%s) already exists, ignoring.", wpa.name))
		return
	}

	tx, err := this.db.Begin()
	if err != nil {
		fmt.Println(err)
	}
	stmt, err := tx.Prepare("insert into wpa(id, name, bssid) values(null, ?, ?)")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(wpa.name, wpa.bssid)
	if err != nil {
		fmt.Println(err)
	}
	tx.Commit()
}


/**
 * Runs a table script on the database file.
 * 
 * @name createNewTable
 */
func (this *DBHandler) createNewTable(tablescript string) {
	_, err := this.db.Exec(tablescript)
	if err != nil {
		log.Error(fmt.Sprintf("%q: %s\n", err, tablescript))
		return
	}
}



/**
 * Checks if the given wpa name is uniqe
 * 
 * @name isWpaUniqe
 * @return bool
 */
func (this *DBHandler) isWpaUniqe(name string) bool {
	stmt, err := this.db.Prepare("select count(*) from wpa where name = ?")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	var count int
	err = stmt.QueryRow(name).Scan(&count)
	if err != nil {
		fmt.Println(err)
	}
	if count > 0 {
		return false
	}
	return true
}


/**
 * Checks if the given wordlist name is uniqe
 * 
 * @name isWordlistUniqe
 * @return bool
 */
func (this *DBHandler) isWordlistUniqe(name string) bool {
	stmt, err := this.db.Prepare("select count(*) from wordlists where name = ?")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	var count int
	err = stmt.QueryRow(name).Scan(&count)
	if err != nil {
		fmt.Println(err)
	}
	if count > 0 {
		return false
	}
	return true
}