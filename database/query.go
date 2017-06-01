package database

import (
    "fmt"
    "strings"

    "github.com/devigner/mysql-extract-record/model"
)

func PrintResult(){
    fmt.Printf("Inserts:\n%v\n",strings.Join(insertQueries,"\n"))
    fmt.Printf("Selects:\n%v\n",strings.Join(selectQueries,"\n"))
}

func queryReferencedRelation(k *model.KeyColumnUsage, value string) {
    Query(k.TableName, k.ColumnName, value)
}

func queryRelationReverse(k *model.KeyColumnUsage, value map[string]string) {
    Query(k.ReferencedTableName.String, k.ReferencedColumnName.String, value[k.ColumnName])
}

var selectQueries = []string{}
var insertQueries = []string{}

type KeyValue struct {
    key   string
    value []byte
}

func B2S(bs []uint8) string {
    ba := []byte{}
    for _, b := range bs {
        ba = append(ba, byte(b))
    }
    return string(ba)
}

func createInsert(table string, resultSet map[string]string){
    colums := []string{}
    data := []string{}
    for key, value := range resultSet{
        colums = append(colums,key)
        data = append(data,value)
    }

    output := fmt.Sprintf("INSERT INTO `%v` (`%v`) VALUES ('%v');", table, strings.Join(colums, "`,`"), strings.Join(data, "','"))
    registerInsertQuery(output)
}

func registerSelectQuery(query string) bool {
    for _, d := range selectQueries {
        if d == query {
            //fmt.Printf("Query already executed, skipped")
            return false
        }
    }
    selectQueries = append(selectQueries,query)
    return true
}

func registerInsertQuery(query string) bool {
    for _, d := range insertQueries {
        if d == query {
            //fmt.Printf("Query already executed, skipped")
            return false
        }
    }
    insertQueries = append(insertQueries,query)
    return true
}

func Query(table string, field string, value string) {

    if "" == value {
        return
    }

    t, _ := getTable(table)
    if !hasColumn(t,field){
        return
    }

    listReferencedTables := getKeyColumnReferencedTable(table)
    listTables := getKeyColumnTable(table)
    //p := getPrimaryKey(table)

    query := fmt.Sprintf("SELECT * FROM `%v` where `%v`='%v'", table, field, value)
    if !registerSelectQuery(query){
        return
    }


    rows, _ := DBTarget.Queryx(query)

    data := []string{}

    defer rows.Close()

    for rows.Next() {
        results := make(map[string]interface{})
        resultSet := make(map[string]string)
        _ = rows.MapScan(results)
        for n, f := range results {
            ab := B2S(f.([]uint8))
            data = append(data, ab)
            resultSet[n] = ab
        }

        if len(listTables) > 0 {
            for _, keyColumn := range listTables {
                queryRelationReverse(keyColumn, resultSet)
            }
        }
        createInsert(table, resultSet)
    }




    if len(listReferencedTables) > 0 {
        for _, keyColumn := range listReferencedTables {
            queryReferencedRelation(keyColumn, value)
        }
    }



    //spew.Dump(t)
    //spew.Dump(listReferencedTables)
    //spew.Dump(p)
    //spew.Dump(query)

    //spew.Dump(vals)

}
