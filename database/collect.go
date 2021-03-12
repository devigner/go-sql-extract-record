package database

import (
	"fmt"
	"os"

	"github.com/devigner/mysql-extract-record/model"
)

var tables = &model.Tables{
	Current: 0,
	Tables:  []*model.Table{},
}

var keyColumnUsage = []*model.KeyColumnUsage{}

// Init build the connection with mysql
func Init(host string, user string, pass string, port int64, database string) {

	SetupDatabase(host, user, pass, port, database)

	columns := []*model.TableColumns{}

	query := fmt.Sprintf("SELECT * FROM `COLUMNS` where `TABLE_SCHEMA`='%v' order by `TABLE_NAME`", database)
	fmt.Printf("/* COLUMNS: */ %v;\n", query)
	err := ConnectionInformationSchema.Select(&columns, query)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	query = fmt.Sprintf("SELECT * FROM `KEY_COLUMN_USAGE` where `TABLE_SCHEMA`='%v' order by `TABLE_NAME`", database)
	fmt.Printf("/* KEY_COLUMN_USAGE: */ %v;\n", query)
	err = ConnectionInformationSchema.Select(&keyColumnUsage, query)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	for _, v := range columns {
		if len(tables.Tables) == 0 || tables.Tables[tables.Current].Name != v.TableName {
			tables.Tables = append(tables.Tables, &model.Table{})
			tables.Current = len(tables.Tables) - 1
		}
		tables.Tables[tables.Current].Columns = append(tables.Tables[tables.Current].Columns, v)
		tables.Tables[tables.Current].Name = v.TableName
	}
	//spew.Dump(keyColumnUsage)
}

func getTable(table string) (*model.Table, error) {
	for _, v := range tables.Tables {
		if v.Name == table {
			return v, nil
		}
	}
	return nil, fmt.Errorf("Could not find table: %v", table)
}

func hasColumn(table *model.Table, field string) bool {
	has := false
	for _, column := range table.Columns {
		if field == column.ColumnName {
			has = true
		}
	}
	return has
}

func getKeyColumnReferencedTable(table string) []*model.KeyColumnUsage {
	usages := []*model.KeyColumnUsage{}
	for _, keyColumn := range keyColumnUsage {
		if keyColumn.ReferencedTableName.Valid {
			if table == keyColumn.ReferencedTableName.String {
				usages = append(usages, keyColumn)
			}
		}
	}
	//spew.Dump(fmt.Sprintf("getKeyColumnReferencedTable: %v -> %v",table,len(usages)))
	return usages
}

func getKeyColumnTable(table string) []*model.KeyColumnUsage {
	usages := []*model.KeyColumnUsage{}
	for _, keyColumn := range keyColumnUsage {
		if keyColumn.ReferencedTableName.Valid {
			if table == keyColumn.TableName {
				usages = append(usages, keyColumn)
			}
		}
	}
	//spew.Dump(fmt.Sprintf("getKeyColumnTable: %v -> %v",table,len(usages)))
	return usages
}

// getPrimaryKey
func getPrimaryKey(table string) *model.KeyColumnUsage {
	for _, keyColumn := range keyColumnUsage {
		if table == keyColumn.TableName && "PRIMARY" == keyColumn.ConstraintName {
			return keyColumn
		}
	}
	return nil
}
