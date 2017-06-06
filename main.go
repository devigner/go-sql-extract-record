package main

import (
    "fmt"

    "github.com/devigner/mysql-extract-record/database"
    "github.com/devigner/mysql-extract-record/flags"
)

func main() {

    config := flags.Init()

    fmt.Printf(
        "/*\nExtracting:\n  Database: `%v`\n  Table: `%v`\n  Field: `%v`\n  Value: `%v`\n*/ \n\n",
        config.DbDatabase,
        config.Table,
        config.Field,
        config.Value,
    )

    database.Init(
        config.DbHost,
        config.DbUser,
        config.DbPass,
        config.DbPort,
        config.DbDatabase,
    )

    path := []string{}
    path = append(path, config.Table)
    database.SelectFromDB(
        config.Table,
        config.Field,
        config.Value,
        0,
        path,
    )
    database.PrintResult(
        config.DbDatabase,
        config.Output,
    )
}
