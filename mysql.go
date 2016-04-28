package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
	"strings"
)

type MySQL struct {
	Config *Config
}

func (m MySQL) Columns(db *gorm.DB) []Column {
	columns := make([]Column, 0)
	scope := db.Select("table_schema, table_name, column_name, data_type").Table("information_schema.columns")

	for _, ignore := range m.Config.Ignores {
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
		col := Column{db: db, Config: m.Config}
		rows.Scan(&col.Database, &col.Table, &col.Name, &col.Type)
		columns = append(columns, col)
	}

	return columns
}

func (m MySQL) Scan(r chan Result) error {
	db, err := openDatabase(m.Config.Source)
	if err != nil {
		return err
	}
	defer db.Close()

	columns := m.Columns(db)
	fmt.Printf("Scanning %d database columns\n", len(columns))

	for _, col := range columns {
		if col.Kind() != reflect.String {
			continue
		}

		if err := col.Scan(r); err != nil {
			return err
		}
	}

	return nil
}

func ScanMySQL(c *Config, r chan Result) error {
	m := MySQL{c}
	return m.Scan(r)
}
