package database

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func DBConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "phong"
    dbName := "go_test"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error)
    }

    return db
}
// func main() {
