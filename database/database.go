package database

import (
	"fmt"
	"os"

	// driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// ConnectionInformationSchema the connection with the information schema database
var ConnectionInformationSchema *sqlx.DB

// ConnectionTargetSchema the connection with the actual database
var ConnectionTargetSchema *sqlx.DB

// SetupDatabase make the actual connection
func SetupDatabase(host string, user string, pass string, port int64, database string) {

	var err error

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, pass, host, port, "information_schema")

	ConnectionInformationSchema, err = sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		fmt.Printf("Error :%v", err)
		os.Exit(1)
	}

	dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, pass, host, port, database)

	ConnectionTargetSchema, err = sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		fmt.Printf("Error :%v", err)
		os.Exit(1)
	}
}
