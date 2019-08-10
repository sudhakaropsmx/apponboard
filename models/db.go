package models

import (
   "fmt"
   "database/sql"
    _ "github.com/lib/pq"
    "log"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "password"
  dbname   = "apponboard"
)
var db *sql.DB

func InitDB(){

   psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
     host, port, user, password, dbname)
   var err error
   db, err = sql.Open("postgres", psqlInfo)
  
   if err != nil {
   	   fmt.Printf("Encountered Postgress Db connect",err)
       log.Panic(err)
    }

    if err = db.Ping(); err != nil {
    	fmt.Printf("Encountered Postgress Db connect",err)
        log.Panic(err)
        
    }
}