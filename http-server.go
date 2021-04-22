package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v2"
	"log"
	"net/http"
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
		fmt.Printf("setup failed")
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
	fmt.Printf("creteManager")
	CreateManager("mypass", "password")
	log.Fatal(http.ListenAndServe(":8080", nil))
	CreateManager("mypass", "password")
}

func (dbManger *DBManager) setup() {
	fmt.Printf("setup succceeded")
}
