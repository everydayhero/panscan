package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	DefaultRegex = "(^|[[:blank:]])[1-9]([ -]?[0-9]){12,15}([[:blank:]]|$)"
)

type Config struct {
	Driver           string
	Conn             string
	IgnoredDatabases []string
	IgnoredTables    []string
	Regex            string
}

func (c Config) Check(content string) bool {
	re := regexp.MustCompile(c.Regex)

	for _, s := range re.FindAllString(content, -1) {
		if Check(s) {
			return true
		}
	}

	return false
}

type multistring []string

func (m *multistring) Set(value string) error {
	if len(*m) > 0 {
		return errors.New("value already set")
	}

	for _, s := range strings.Split(value, ",") {
		*m = append(*m, s)
	}

	return nil
}

func (m *multistring) String() string {
	return fmt.Sprint(*m)
}

func GetConfig() Config {
	ignoredDatabases := multistring{}
	ignoredTables := multistring{}

	flag.Var(&ignoredDatabases, "d", "A list of databases to ignore")
	flag.Var(&ignoredTables, "t", "A list of tables to ignore")
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("The driver and connection string must be provided. Received: %s", args)
		os.Exit(2)
	}

	return Config{
		Regex:            DefaultRegex,
		IgnoredDatabases: ignoredDatabases,
		IgnoredTables:    ignoredTables,
		Driver:           args[0],
		Conn:             args[1],
	}
}
