package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/gorp.v2"
)

type DBManager struct {
	DBMap *gorp.DbMap
}

func handler(w http.ResponseWriter, r *http.Request) {
	text := "hogehoge."
	fmt.Fprint(w, "Hi there, I'm ", text)
}

func CreateManager(user string, password string) *DBManager {
	var source = user + ":" + password + "@tcp(localhost:3306)/discord_clone_database?parseTime=true"
	var db, error = sql.Open("mysql", source)
	if error != nil {
		log.Fatal("error connecting to database:", error)
	}
	var gorpDialect = gorp.MySQLDialect{}
	var dbMap = &gorp.DbMap{Db: db, Dialect: gorpDialect}
	var manager = &DBManager{DBMap: dbMap}
	manager.setup()

	return manager
}
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
