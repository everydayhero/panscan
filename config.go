package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

const (
	DefaultRegex = "([0-9][ -]*){13,16}"
)

type Config struct {
	Driver           string
	Conn             string
	IgnoredDatabases []string
	IgnoredTables    []string
	Regex            string
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

	if flag.NArg() != 2 {
		panic("Invalid arguments")
	}

	args := flag.Args()
	return Config{
		Regex:            DefaultRegex,
		IgnoredDatabases: ignoredDatabases,
		IgnoredTables:    ignoredTables,
		Driver:           args[0],
		Conn:             args[1],
	}
}
