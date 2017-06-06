package database

import (
    "fmt"
    "io/ioutil"
    "strings"

    "github.com/devigner/mysql-extract-record/model"
)

func PrintResult(databaseName string, output bool) {
    result := fmt.Sprintf("%v\n", strings.Join(insertQueries, "\n"))
    if output {
        fmt.Printf(result)
    } else {
        d1 := []byte(result)
        err := ioutil.WriteFile(fmt.Sprintf("%v.sql",databaseName), d1, 0644)
        if err != nil {
            fmt.Printf("Error writing file: %v",err)
        }
    }
}

func SliceIndex(needle string, haystack []string) int {
    for i := 0; i < len(haystack); i++ {
        if haystack[i] == needle {
            return i
        }
    }
    return -1
}

func queryReferencedRelation(k *model.KeyColumnUsage, value string, depth int, path []string) {
    if SliceIndex(k.TableName, path) > -1 {
        fmt.Printf("%v→ skipped: %v\n", strings.Repeat(" ", depth), k.TableName)
        return;
    }
    SelectFromDB(k.TableName, k.ColumnName, value, depth, path)
}

func queryRelationReverse(k *model.KeyColumnUsage, value map[string]string, depth int, path []string) {
    SelectFromDB(k.ReferencedTableName.String, k.ReferencedColumnName.String, value[k.ColumnName], depth, path)
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

func createInsert(table string, resultSet map[string]string) {
    colums := []string{}
    data := []string{}
    for key, value := range resultSet {
        colums = append(colums, key)
        data = append(data, strings.Replace(value, "'", `\'`, -1))
    }

    output := fmt.Sprintf("INSERT INTO `%v` (`%v`) VALUES ('%v');", table, strings.Join(colums, "`,`"), strings.Join(data, "','"))
    registerInsertQuery(output)
}

var skipped = 0

func registerSelectQuery(table string, field string, value string, depth int) string {
    key := fmt.Sprintf("%v.%v = %v", table, field, value)

    for _, d := range selectQueries {
        if d == key {
            skipped ++
            return ""
        }
    }

    fmt.Printf("%v→ %v\n", strings.Repeat(" ", depth), key)

    query := fmt.Sprintf("SELECT * FROM `%v` where `%v`='%v';", table, field, value)

    selectQueries = append(selectQueries, key)

    return query
}

func registerInsertQuery(query string) bool {
    for _, d := range insertQueries {
        if d == query {
            return false
        }
    }
    insertQueries = append(insertQueries, query)
    return true
}

func SelectFromDB(table string, field string, value string, depth int, path []string) {

    if "" == value {
        return
    }

    t, _ := getTable(table)
    if false == hasColumn(t, field) {
        return
    }

    query := registerSelectQuery(table, field, value, depth)
    if query == "" {
        return
    }

    rows, err := DBTarget.Queryx(query)
    if err != nil {
        fmt.Printf("Queryx error: %v\n", err)
        return
    }
    listReferencedTables := getKeyColumnReferencedTable(table)
    listTables := getKeyColumnTable(table)

    defer rows.Close()

    path = append(path, table)

    for rows.Next() {
        results := make(map[string]interface{})
        resultSet := make(map[string]string)
        err = rows.MapScan(results)
        if err != nil {
            fmt.Printf("MapScan error: %v\n", err)
            return
        }
        for n, f := range results {
            ab := ""
            if f != nil {
                ab = B2S(f.([]uint8))
            }
            resultSet[n] = ab
        }

        if len(listTables) > 0 {
            for _, keyColumn := range listTables {
                queryRelationReverse(keyColumn, resultSet, depth+1, path)
            }
        }
        createInsert(table, resultSet)
    }

    if len(listReferencedTables) > 0 {
        for _, keyColumn := range listReferencedTables {
            queryReferencedRelation(keyColumn, value, depth+1, path)
        }
    }

    //spew.Dump(t)
    //spew.Dump(listReferencedTables)
    //spew.Dump(p)
    //spew.Dump(query)

    //spew.Dump(vals)

}
