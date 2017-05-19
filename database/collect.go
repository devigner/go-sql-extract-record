package database

import (
    "fmt"

    "github.com/davecgh/go-spew/spew"
    "github.com/devigner/record-extract/model"
)



func Init() {
    SetupDatabase()

    db := GetInfoDb()

    columns := []*model.TableColumns{}
    constraints := []*model.TableConstraints{}

    err := db.Select(&columns, "SELECT * FROM `COLUMNS` where `TABLE_SCHEMA`='test_db' order by `TABLE_NAME`")
    if err != nil {
        fmt.Printf("Error: %v", err)
    }

    err = db.Select(&constraints, "SELECT * FROM `TABLE_CONSTRAINTS` where `TABLE_SCHEMA`='test_db' order by `TABLE_NAME`")
    if err != nil {
        fmt.Printf("Error: %v", err)
    }

    var tables = &model.Tables{
        Current: 0,
        Tables:  []*model.Table{},
    }

    for _, v := range columns {
        if len(tables.Tables) == 0 || tables.Tables[tables.Current].Name != v.TableName {
            tables.Tables = append(tables.Tables, &model.Table{})
            tables.Current = len(tables.Tables) - 1
        }
        tables.Tables[tables.Current].Columns = append(tables.Tables[tables.Current].Columns, v)
        tables.Tables[tables.Current].Name = v.TableName
    }
    spew.Dump(tables)
}
