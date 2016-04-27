package main

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

const (
	DefaultRegex = "(^|[[:blank:]])[1-9]([ -]?[0-9]){12,15}([[:blank:]]|$)"
)

var (
	NoDatabaseSource = errors.New("No database connections provided")
)

type Config struct {
	Source  *url.URL
	Ignores []string
	Regex   string
}

func (c *Config) Check(content string) bool {
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

func GetConfig(arguments []string) (*Config, error) {
	ignores := multistring{}
	flagSet := flag.NewFlagSet("panscan", flag.ContinueOnError)

	flagSet.Var(&ignores, "i", "Ignores a database, table, or column. Format DATABASE[.TABLE[.COLUMN]]")
	if err := flagSet.Parse(arguments); err != nil {
		return nil, err
	}

	rawurl := flagSet.Arg(0)
	if rawurl == "" {
		return nil, NoDatabaseSource
	}

	source, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}

	return &Config{
		Regex:   DefaultRegex,
		Source:  source,
		Ignores: ignores,
	}, nil
}
