package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"net/url"
	"strings"
)

func openDatabase(url *url.URL) (*gorm.DB, error) {
	var conn string

	driver := url.Scheme
	switch driver {
	case "mysql":
		conn = formatMysql(url)
	default:
		conn = formatDefault(url)
	}

	return gorm.Open(driver, conn)
}

func formatDefault(url *url.URL) string {
	return url.String()
}

func formatMysql(url *url.URL) string {
	url.Scheme = ""
	if url.Host != "" {
		if !strings.Contains(url.Host, ":") {
			url.Host = url.Host + ":3306"
		}
		url.Host = "tcp(" + url.Host + ")"
	}
	if url.Path == "" {
		url.Path = "/"
	}
	return url.String()[2:]
}
