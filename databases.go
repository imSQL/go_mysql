package imSQL

import (
	"database/sql"
	"fmt"

	"github.com/juju/errors"
)

type (
	DB struct {
		Name                 string `json:"name" db:"name"`
		DefaultCharacterSet  string `json:"default_character_set_name" db:"default_character_set_name"`
		DefaultCollationName string `json:"default_collation_name" db:"default_collation_name"`
	}
)

const (
	// StmtCreateOneDatabase /*Create one database*/
	StmtCreateOneDatabase = `
	CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARACTER SET utf8 DEFAULT COLLATE utf8_general_ci
	`

	// StmtDropOneDatabase /*Drop a database.*/
	StmtDropOneDatabase = `
	DROP DATABASE %s	
	`

	// StmtQueryUsersDB /*Query user's database*/
	StmtQueryUsersDB = `
	SELECT Db FROM mysql.db WHERE User = '%s'
	`
	// StmtQueryAllDB /*Query all databases.*/
	StmtQueryAllDB = `
	SELECT schema_name,default_character_set_name,default_collation_name FROM information_schema.schemata WHERE schema_name NOT IN ('information_schema','performance_schema','sys','mysql')
	`
)

/*
NewDB return a new db handler.
*/
func NewDB(name string) (*DB, error) {
	newdb := new(DB)

	newdb.Name = name
	newdb.DefaultCharacterSet = "utf8"
	newdb.DefaultCollationName = "utf8_general_ci"

	return newdb, nil
}

/*
Create one database.
*/
func (dbi *DB) CreateOneDB(db *sql.DB) error {

	Query := fmt.Sprintf(StmtCreateOneDatabase, dbi.Name)
	_, err := db.Exec(Query)
	if err != nil {
		return errors.Trace(err)
	}

	return nil
}

/*
Drop one databases.
*/
func (dbi *DB) DropOneDB(db *sql.DB) error {

	Query := fmt.Sprintf(StmtDropOneDatabase, dbi.Name)
	_, err := db.Exec(Query)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}
