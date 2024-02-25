package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
	"github.com/sebasdev2001/bank_app/util"
)

var (
  host = util.GetEnvVar("DB_HOST")
  port = util.GetEnvVar("DB_PORT")
  user = util.GetEnvVar("DB_USER")
  password = util.GetEnvVar("DB_PASSWORD")
  db_name = util.GetEnvVar("DB_NAME")
)

var lock = &sync.Mutex{}

var dbInstance *Database

type Database struct {
	db *sql.DB
}

func (d *Database) Init() (*sql.DB, error) {
  // set tables for the first time


  return  d.db, nil
}


func GetDbInstance() *Database {
	if dbInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		db, err := newDatabase()
		if err != nil {
      log.Fatal(err)
		}
		dbInstance = db
	}
	return dbInstance
}

func newDatabase() (*Database, error) {
  fmt.Printf("port: %v\n", port)
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
                          host, port, user, password, db_name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to databse: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Error pingin database: %v", err)
	}

  log.Println("Connection established to database!")

	return &Database{
		db: db,
	}, nil
}
