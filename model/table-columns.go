package model

import "database/sql"

type Tables struct {
    Current int
    Tables  []*Table
}

type Table struct {
    Name    string
    Columns []*TableColumns
}

type TableColumns struct {
    TableCatalog           string `db:"TABLE_CATALOG"`
    TableSchema            string `db:"TABLE_SCHEMA"`
    TableName              string `db:"TABLE_NAME"`
    ColumnName             string `db:"COLUMN_NAME"`
    OrdinalPosition        uint `db:"ORDINAL_POSITION"`
    ColumnDefault          sql.NullString `db:"COLUMN_DEFAULT"`
    IsNullable             string `db:"IS_NULLABLE"`
    DataType               string `db:"DATA_TYPE"`
    CharacterMaximumLength sql.NullInt64 `db:"CHARACTER_MAXIMUM_LENGTH"`
    CharacterOctetLength   sql.NullInt64 `db:"CHARACTER_OCTET_LENGTH"`
    NumericPrecision       sql.NullInt64 `db:"NUMERIC_PRECISION"`
    NumericScale           sql.NullInt64 `db:"NUMERIC_SCALE"`
    DatetimePrecision      sql.NullInt64 `db:"DATETIME_PRECISION"`
    CharacterSetName       sql.NullString `db:"CHARACTER_SET_NAME"`
    CollationName          sql.NullString `db:"COLLATION_NAME"`
    ColumnType             string `db:"COLUMN_TYPE"`
    ColumnKey              string `db:"COLUMN_KEY"`
    Extra                  string `db:"EXTRA"`
    Privileges             string `db:"PRIVILEGES"`
    ColumnComment          string `db:"COLUMN_COMMENT"`
    GenerationExpression   string `db:"GENERATION_EXPRESSION"`
}

type TableConstraints struct {
    ConstraintCatalog string `db:"CONSTRAINT_CATALOG"`
    ConstraintSchema  string `db:"CONSTRAINT_SCHEMA"`
    ConstraintName    string `db:"CONSTRAINT_NAME"`
    TableSchema       string `db:"TABLE_SCHEMA"`
    TableName         string `db:"TABLE_NAME"`
    ConstraintType    string `db:"CONSTRAINT_TYPE"`
}

type ReferentialConstraints struct {
    ConstraintCatalog       string `db:"CONSTRAINT_CATALOG"`
    ConstraintSchema        string `db:"CONSTRAINT_SCHEMA"`
    ConstraintName          string `db:"CONSTRAINT_NAME"`
    UniqueConstraintCatalog string `db:"UNIQUE_CONSTRAINT_CATALOG"`
    UniqueConstraintSchema  string `db:"UNIQUE_CONSTRAINT_SCHEMA"`
    UniqueConstraintName    string `db:"UNIQUE_CONSTRAINT_NAME"`
    MatchOption             string `db:"MATCH_OPTION"`
    UpdateRule              string `db:"UPDATE_RULE"`
    DeleteRule              string `db:"DELETE_RULE"`
    TableName               string `db:"TABLE_NAME"`
    ReferencedTableName     string `db:"REFERENCED_TABLE_NAME"`
}
