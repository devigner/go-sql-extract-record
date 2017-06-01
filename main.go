package main

import (
    "flag"
    "github.com/devigner/mysql-extract-record/database"
    "fmt"
)

func main() {
    host := flag.String("localhost", "localhost", "a string")
    user := flag.String("user", "root", "a string")
    pass := flag.String("pass", "root", "a string")
    port := flag.Int64("port", 3306, "a string")
    databaseName := flag.String("database", "test-database", "a string")

    table := flag.String("table", "user", "a string")
    field := flag.String("field", "id", "a string")
    value := flag.String("value", "1", "a string")
    flag.Parse()

    fmt.Printf("Extracting: %v %v %v %v", *databaseName, *table, *field, *value)

    database.Init(*host, *user, *pass, *port, *databaseName)
    database.Query(*table, *field, *value)
    database.PrintResult()
}
