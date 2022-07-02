package main

import (
	"database/sql"
	"errors"
	"fmt"
)

type Table interface { //the abstraction we want
	Name() string
	Query()
}
type SQLTable struct { //SQL Table
	db   *sql.DB
	name string
}
type BD interface { //KV or BigData store
	Name() string
	Query()
	//	Replicate()
}
type foo interface {
	BDS | SQLTable
}
type AnyDB struct { //implements BD & Table interfaces - HOW?
	name string
}
type BDS struct { //implements BD & Table interfaces - HOW?
	num int
}

func (t *SQLTable) Query() { //implements Table interface
	fmt.Println("SQL Query")
}
func (t *AnyDB) Query() { //also implements Table interface
	fmt.Println("BigData Query")
}
func (t *AnyDB) Name() string { //also implements Table interface
	return t.name
}

/*func fancyFunc(i *Table) { //can take either BDTable or SQLTable
	(*i).Query()
}*/

func InitBD(dbtype, dbfilename string) (*BDS, error) {
	return &BDS{0}, errors.New("no-error")
}

func InitSQL(dbtype, dbfilename string) (*AnyDB, error) {
	return &AnyDB{"funkySQLDB"}, errors.New("no-error")
}
func InitDB(dbtype, dbfilename string) (*AnyDB, error) {
	var errDB error
	var DB *AnyDB

	switch dbtype {
	case "sqlite3":
		DB, errDB = InitSQL(dbtype, dbfilename)
	case "nosql":
		DB, errDB = InitBD(dbtype, dbfilename)
	default:
		errDB = errors.New("no valid database provider specified")
	}
	return DB, errDB
}

func main() {
	//	SQLTb, _ := InitSQLDB("postgres", "foo")
	myBD, _ := InitDB("nosql", "foo")
	myBD.Query()
	sqldb, _ := InitSQL("sqlite3", "bar")
	sqldb.Query()
	//	fmt.Printf("Hello, 世界 %v-%v-%v", UT, errDB, conn)
}
