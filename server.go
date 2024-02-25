package main

import (
	"log"

	db "github.com/sebasdev2001/bank_app/db"
)

func main() {
  database, err := db.GetDbInstance().Init()
  if err != nil {
    log.Fatal(err)
  }
  store := NewStore(database)
	server := NewAPIServer(":4545", store)
	server.Run()
}
