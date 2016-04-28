package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Column struct {
	db       *gorm.DB
	Config   *Config
	Database string
	Table    string
	Name     string
	Type     string
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
