package config

import (
	"testing"

	"rxdrag.com/entify-schema-registry/consts"
)

func TestSetString(t *testing.T) {
	Init()
	SetString(consts.DB_DRIVER, "test_value")
	if GetString(consts.DB_DRIVER) != "test_value" {
		t.Error("Error SetString,expected 'test_value', but is:" + GetString(consts.DB_DRIVER))
	}
}
