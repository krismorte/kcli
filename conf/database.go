package conf

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var dbPath = "./db/conf.db"

func initDatabase() {
	if _, err := os.Stat("/db"); os.IsNotExist(err) {
		os.Mkdir("db", 0777)
	}

	database, _ := sql.Open("sqlite3", dbPath)
	defer database.Close()
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS repo_conf (id INTEGER PRIMARY KEY, repo_type TEXT, key_value TEXT, value TEXT)")
	statement.Exec()
	//SelectAll()

}

func RunDatabaseCommand(command string) {
	initDatabase()
	database, err := sql.Open("sqlite3", dbPath)
	if database == nil {
		fmt.Println("db nulo")
	}
	defer database.Close()
	if err != nil {
		err.Error()
	}
	statement, err := database.Prepare(command)
	if statement == nil {
		fmt.Println("db nulo")
	}
	defer statement.Close()

	if err != nil {
		err.Error()
	}
	_, err = statement.Exec()
	if err != nil {
		err.Error()
	}
}

func RunDatabaseQuery(query string) *sql.Rows {
	initDatabase()
	database, _ := sql.Open("sqlite3", dbPath)
	defer database.Close()
	rows, err := database.Query(query)
	if err != nil {
		err.Error()
	}
	return rows
}

func SelectAll() {
	fmt.Println("consulta")
	database, _ := sql.Open("sqlite3", dbPath)
	defer database.Close()
	rows, _ := database.Query("SELECT repo_type , key_value , value  FROM repo_conf")
	var id int
	var typeG string
	var key string
	var value string
	for rows.Next() {
		rows.Scan(&id, &typeG, &key, &value)
		fmt.Println(strconv.Itoa(id) + ": " + typeG + " " + key + " " + value)
	}
}
