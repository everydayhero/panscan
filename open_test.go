package main

import (
	"net/url"
	"os"
	"testing"
)

func Test__openDatabase(t *testing.T) {
	url, err := url.Parse(os.Getenv("SOURCE"))
	if err != nil {
		t.Errorf("Received error parsing env SOURCE: %v", err)
	}

	db, err := openDatabase(url)
	if err != nil {
		t.Errorf("Received error for test case %s: %v", url, err)
	} else if db == nil {
		t.Fatalf("Expected *gorm.DB but received nil for test case %s", url)
	}
}
