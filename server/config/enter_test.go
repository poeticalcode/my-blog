package config

import (
	"testing"
)

func TestReadConfig(t *testing.T) {
	err := readConfig()
	if err != nil {
		t.Error(err)
	}
}
