package main

//poly2 branch
import (
	"database/sql"
	"errors"
	"fmt"
)

type Table interface { //the abstraction we want
	Name() string
	Query()
}
type SQL struct { //SQL Table
	db   *sql.DB
	name string
}
type NoSQL struct { //implements BD & Table interfaces - HOW?
	guid int
}

type foo interface {
	NoSQL | SQL
}
type foo2 struct {
}

func (t *SQL) Query() { //implements Table interface
	fmt.Println("SQL Query")
}
func (t *NoSQL) Query() { //also implements Table interface
	fmt.Println("BigData Query")
}
func (t *SQL) Name() string { //also implements Table interface
	return t.name
}
func (t *NoSQL) Name() int { //also implements Table interface
	return t.guid
}

func InitNoSQL(dbtype, dbfilename string) (*NoSQL, error) {
	return &NoSQL{0}, errors.New("no-error")
}

func InitSQL(dbtype, dbfilename string) (*SQL, error) {
	return &SQL{nil, "funkySQLDB"}, errors.New("no-error")
}

/* func InitDB(dbtype, dbfilename string) (*AnyDB, error) {
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
} */

func main() {
	//	SQLTb, _ := InitSQLDB("postgres", "foo")
	myBD, _ := InitNoSQL("nosql", "foo")
	myBD.Query()
	sqldb, _ := InitSQL("sqlite3", "bar")
	sqldb.Query()
	//	fmt.Printf("Hello, 世界 %v-%v-%v", UT, errDB, conn)
}
