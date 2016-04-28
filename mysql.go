package main

import (
	"fmt"
	"reflect"
)

type MySQL struct {
	Config *Config
}

func (m MySQL) Scan(r chan Result) error {
	c := m.Config
	db, err := openDatabase(c.Source)
	if err != nil {
		return err
	}
	defer db.Close()

	columns := GetColumns(db, c)
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
