package imSQL

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/juju/errors"
)

type (
	Conn struct {
		Addr      string
		Port      uint64
		User      string
		Password  string
		Database  string
		Charset   string
		Collation string
		DBI       string
	}
)

// NewConn new mysql connection handler.
func NewConn(addr string, port uint64, user string, password string) (*Conn, error) {
	imsql := new(Conn)
	imsql.Addr = addr
	imsql.Port = port
	imsql.User = user
	imsql.Password = password
	imsql.Database = "information_schema"
	imsql.Charset = "utf8"
	imsql.Collation = "utf8_general_ci"

	return imsql, nil
}

// set characterset ,default utf8
func (imsql *Conn) SetCharset(charset string) {
	imsql.Charset = charset
}

// set collation.default utf8_general_ci.
func (imsql *Conn) SetCollation(collation string) {
	imsql.Collation = collation
}

// make a mysql dbi string.
func (imsql *Conn) MakeDBI() {
	imsql.DBI = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s", imsql.User, imsql.Password, imsql.Addr, imsql.Port, imsql.Database, imsql.Charset, imsql.Collation)
}

// open a new mysql connection.
func (imsql *Conn) OpenConn() (*sql.DB, error) {

	db, err := sql.Open("mysql", imsql.DBI)
	if err != nil {
		return nil, errors.Trace(err)
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.Trace(err)
	}
	return db, nil
}

// close a mysql connection.
func (imsql *Conn) CloseConn(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}
