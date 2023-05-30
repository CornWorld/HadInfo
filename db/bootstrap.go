package db

import (
	"database/sql"
	"fmt"
	"github.com/gookit/ini/v2"
	_ "github.com/lib/pq"
	"log"
	"strconv"
)

var db *sql.DB

func Bootstrap() {
	log.Println("Connecting PGSQL...")
	dbConfig := ini.StringMap("db")
	port, _ := strconv.Atoi(dbConfig["port"])
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbConfig["host"], port, dbConfig["user"], dbConfig["password"], dbConfig["name"], dbConfig["sslMode"])
	if d, err := sql.Open("postgres", psqlInfo); err != nil {
		log.Fatal("Connect PGSQL Failed: ", err)
	} else {
		db = d
	}
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	if err := db.Ping(); err != nil {
		log.Fatal("Ping PGSQL Failed: ", err)
	}
	log.Println("PGSQL Successful Connected!")
}

func Exit() {
	_ = db.Close()
}
