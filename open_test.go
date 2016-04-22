package main

import (
	"os"
	"testing"
)

func Test__openDatabase(t *testing.T) {
	c := os.Getenv("SOURCE")
	db, err := openDatabase(c)
	if err != nil {
		t.Errorf("Received error for test case %s: %v", c, err)
	} else if db == nil {
		t.Fatalf("Expected *gorm.DB but received nil for test case %s", c)
	}
}
