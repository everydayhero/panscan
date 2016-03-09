package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"reflect"
)

type Database struct {
	Config Config
}

func (d Database) Scan(r chan Result) error {
	c := d.Config
	db, err := gorm.Open(c.Driver, c.Conn)
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

func NewDatabase(c Config) Database {
	return Database{c}
}
