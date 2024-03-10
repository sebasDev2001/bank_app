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
	host     = util.GetEnvVar("DB_HOST")
	port     = util.GetEnvVar("DB_PORT")
	user     = util.GetEnvVar("DB_USER")
	password = util.GetEnvVar("DB_PASSWORD")
	db_name  = util.GetEnvVar("DB_NAME")
)

var lock = &sync.Mutex{}

var dbInstance *Database

type Database struct {
	db *sql.DB
}

func (d *Database) Init() (*sql.DB, error) {
	if err := d.createAccountTable(); err != nil {
		return nil, err
	}
	if err := d.createTrasactionTable(); err != nil {
		return nil, err
	}

	return d.db, nil
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

func (D *Database) createAccountTable() error {
	_, err := D.db.Exec(`
    CREATE TABLE IF NOT EXISTS account (
      id uuid default get_random_uuid(),
      first_name varchar(128) not null,
      last_name varchar(128) not null,
      email varchar(64) unique not null,
      balance numeric(2),
      created_at timestamp,
      PRIMARY KEY (id)
    );
  `)
	return err
}

func (D *Database) createTrasactionTable() error {
	_, err := D.db.Exec(`
    CREATE TABLE IF NOT EXISTS transaction (
      id uuid default get_random_uuid(),
      origin_acc uuid not null,
      det_acc uuid not null,
      amount numeric(2) check (amount >= 0),
      created_at datetime default current_timestamp,
      updated_at datetime,

      foreign key (origin_acc) references account(id),
      foreign key (dest_acc) references account(id)
    );
  `)
	return err
}
