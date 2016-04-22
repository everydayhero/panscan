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
	Source  string
	Ignores []string
	Regex   string
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
	ignores := multistring{}

	flag.Var(&ignores, "i", "Ignores a database, table, or column. Format DATABASE[.TABLE[.COLUMN]]")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Errorf("No database connections provided.")
		os.Exit(2)
	}

	return Config{
		Source:  args[0],
		Ignores: ignores,
		Regex:   DefaultRegex,
	}
}
