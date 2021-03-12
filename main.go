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
		config.Database,
		config.Table,
		config.Field,
		config.Value,
	)

	database.Init(
		config.Host,
		config.User,
		config.Pass,
		config.Port,
		config.Database,
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
		config.Database,
		config.Output,
	)
}
