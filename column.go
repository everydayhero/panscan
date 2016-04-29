package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type Column struct {
	db         *gorm.DB
	Config     *Config
	Database   string
	Table      string
	Name       string
	PrimaryKey string
}

func (c Column) Scan(results chan Result) error {
	table := joinString(".", c.Database, c.Table)
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
			loc := joinString(".", c.Name, c.PrimaryKey)
			results <- Result{Source: table, Loc: loc, Value: value}
		}
	}

	return nil
}

func (c Column) String() string {
	return fmt.Sprintf("{Database: \"%s\", Table: \"%s\", Name: \"%s\", PrimaryKey: \"%s\"}", c.Database, c.Table, c.Name, c.PrimaryKey)
}

func joinString(del string, a ...string) string {
	b := a[:0]
	for _, x := range a {
		if x != "" {
			b = append(b, x)
		}
	}

	return strings.Join(b, del)
}
