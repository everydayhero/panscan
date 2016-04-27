package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
	"strings"
)

type Column struct {
	db       *gorm.DB
	Config   *Config
	Database string
	Table    string
	Name     string
	Type     string
}

func (c Column) Kind() reflect.Kind {
	switch c.Type {
	case "varchar", "char", "text":
		return reflect.String
	default:
		return reflect.Invalid
	}
}

func (c Column) Scan(results chan Result) error {
	table := c.Database + "." + c.Table
	rows, err := c.db.Table(table).Select(c.Name).Rows()
	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		var value string
		if err = rows.Scan(&value); err != nil {
			return err
		}
		if c.Config.Check(value) {
			results <- Result{Source: table, Loc: c.Name, Value: value}
		}
	}

	return nil
}

func (c Column) String() string {
	return fmt.Sprintf("{Database: \"%s\", Table: \"%s\", Name: \"%s\", Type: \"%s\"}", c.Database, c.Table, c.Name, c.Type)
}

func GetColumns(db *gorm.DB, c *Config) []Column {
	columns := make([]Column, 0)
	scope := db.Select("table_schema, table_name, column_name, data_type").Table("information_schema.columns")

	for _, ignore := range c.Ignores {
		parts := strings.Split(ignore, ".")
		columns := strings.Join([]string{"table_schema", "table_name", "column_name"}[0:len(parts)], ", \".\", ")

		scope = scope.Where("CONCAT("+columns+") != ?", ignore)
	}

	rows, err := scope.Rows()
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		col := Column{db: db, Config: c}
		rows.Scan(&col.Database, &col.Table, &col.Name, &col.Type)
		columns = append(columns, col)
	}

	return columns
}
