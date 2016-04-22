package main

import (
	"github.com/jinzhu/gorm"
	"os"
)

func open(c Config) *gorm.DB {
	db, err := openDatabase(c.Source)
	if err != nil {
		panic(err)
	}

	return db
}

func config() Config {
	return Config{
		Source: os.Getenv("SOURCE"),
		Regex:  DefaultRegex,
		Ignores: []string{
			"information_schema",
			"mysql",
			"sys",
			"performance_schema",
			"test_db.creditcards.id",
			"test_db.creditcards.cctype",
			"test_db.posts.id",
			"test_db.posts.title",
			"test_db.posts.published_date",
			"test_db.posts.published_time",
		},
	}
}
