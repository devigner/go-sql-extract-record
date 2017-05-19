package database

import (
    "fmt"
    "os"
    "reflect"
    "strconv"

    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)

var targetDB *sqlx.DB
var infoDB *sqlx.DB

func SetupDatabase() {
    mysqlHost := os.Getenv("MYSQL_HOSTNAME")
    if mysqlHost == "" {
        mysqlHost = "localhost"
    }
    mysqlUser := os.Getenv("MYSQL_USERNAME")
    if mysqlUser == "" {
        mysqlUser = "root"
    }
    mysqlPass := os.Getenv("MYSQL_PASSWORD")
    if mysqlPass == "" {
        mysqlPass = "root"
    }
    mysqlDb := os.Getenv("MYSQL_DATABASE")
    if mysqlDb == "" {
        mysqlDb = "test_db"
    }
    port := os.Getenv("MYSQL_PORT")
    if port == "" {
        port = "3306"
    }
    mysqlPort, _ := strconv.ParseInt(port, 10, 64)

    //var err error

    /*targetDB, err = createDbConnection(mysqlHost, mysqlUser, mysqlPass, mysqlPort, mysqlDb)
    if err != nil {
        fmt.Printf("DB connection failed: %v\n", err)
        return
    }*/
    infoDB, _ = createDbConnection(mysqlHost, mysqlUser, mysqlPass, mysqlPort, "information_schema")
}

// Creates a database connection and check connectivity in one.
func createDbConnection(host string, user string, pass string, port int64, db string) (*sqlx.DB, error) {
    //dataSourceName := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v sslmode=disable", user, pass, host, port, db)
    dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, pass, host, port, db)
    fmt.Printf("DB connection: %v\n", dataSourceName)
    return sqlx.Connect("mysql", dataSourceName)
}

func QueryDb(db *sqlx.DB, query string, args ...interface{}) {
    fmt.Printf("Re Query: %v\n", query)
    rows, err := db.Queryx(query, args...)
    if err != nil {
        fmt.Printf("DB error 1: %v\n", err)
    }

    if rows == nil {
        fmt.Printf("No results for query: %v\n", query)
        return
    }
    for rows.Next() {
        results := make(map[string]interface{})
        err = rows.MapScan(results)
        if err != nil {
            fmt.Errorf("DB error 2: %v\n", err)
            os.Exit(0)
        }

        for k, v := range results {
            fmt.Printf("%v: %v\n", k, reflect.TypeOf(v))
        }
    }
}

func QueryTargetDb(query string, args ...interface{}) {
    QueryDb(targetDB, query, args...)
}

func QueryInfoDb(query string, args ...interface{}) {
    QueryDb(infoDB, query, args...)
}

func GetInfoDb() *sqlx.DB {
    return infoDB
}
