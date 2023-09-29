package platform

import (
	"database/sql"
	"fmt"
	"school/config"

	_ "github.com/lib/pq"
)

func DBConn() *sql.DB{
	cfg := config.Load()
	psqlInfo :=fmt.Sprintf("host=%s port=%d user=%s "+
     "password=%s dbname=%s sslmode=disable",
    cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DB)

db, err :=sql.Open("postgres", psqlInfo)
  if err !=nil {
	panic(err)
  }

  fmt.Println("Succesfully connected!")
  return db
} 