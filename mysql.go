package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

var (
	DataTypes = []string{"varchar", "char", "text"}
)

type MySQL struct {
	Config *Config
}

func (m MySQL) Columns(db *gorm.DB) []Column {
	columns := make([]Column, 0)
	scope := db.Select("c1.table_schema, c1.table_name, c1.column_name, c2.column_name").Table("information_schema.columns c1").Joins("LEFT JOIN information_schema.columns c2 ON c1.table_schema = c2.table_schema AND c1.table_name = c2.table_name AND c2.column_key = 'PRI'")
	scope = scope.Where("c1.data_type IN (?)", DataTypes)

	for _, ignore := range m.Config.Ignores {
		parts := strings.Split(ignore, ".")
		columns := strings.Join([]string{"c1.table_schema", "c1.table_name", "c1.column_name"}[0:len(parts)], ", \".\", ")

		scope = scope.Where("CONCAT("+columns+") != ?", ignore)
	}

	rows, err := scope.Rows()
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		col := Column{db: db, Config: m.Config}
		rows.Scan(&col.Database, &col.Table, &col.Name, &col.PrimaryKey)
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
