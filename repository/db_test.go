package repository

import (
	"testing"

	"rxdrag.com/entify-schema-registry/config"
)

func TestDbString(t *testing.T) {
	dbConfig := config.DbConfig{
		Driver:   "mysql",
		Host:     "127.0.0.1",
		Port:     "3306",
		Database: "Test",
		User:     "root",
		Password: "111",
	}

	cfgStr := DbString(dbConfig)
	if cfgStr != "root:111@tcp(127.0.0.1:3306)/Test?parseTime=true" {
		t.Error("DbString error")
	}
}
