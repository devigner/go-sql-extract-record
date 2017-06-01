package database

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)


var DB *sqlx.DB
var DBTarget *sqlx.DB

func SetupDatabase(host string, user string, pass string, port int64, database string) {

    var err error

    dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, pass, host, port, "information_schema")

    DB, err = sqlx.Connect("mysql", dataSourceName)

    if err != nil {
        fmt.Errorf("Error :%v", err)
    }

    dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, pass, host, port, database)

    DBTarget, err = sqlx.Connect("mysql", dataSourceName)

    if err != nil {
        fmt.Errorf("Error :%v", err)
    }
}
