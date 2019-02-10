package config

import "testing"

func TestRead(t *testing.T) {
	config := Config{}
	err := config.Read("test_config.toml")

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if config.Database != "testdb1" {
		t.Errorf("Incorrect database. Expected: %s, actual: %s", "testdb1", config.Database)
	}

	if config.Server != "testserver1" {
		t.Errorf("Incorrect server. Expected: %s, actual: %s", "testserver1", config.Server)
	}
}

func TestReadError(t *testing.T) {
	config := Config{}
	err := config.Read("no_such_config.toml")

	if err.Error() != "open no_such_config.toml: The system cannot find the file specified." {
		t.Errorf("Incorrect error, expected: %s, actual: %s", "open no_such_config.toml: The system cannot find the file specified.", err)
	}

	if config.Database != "" {
		t.Errorf("Incorrect database. Expected: %s, actual: %s", "", config.Database)
	}

	if config.Server != "" {
		t.Errorf("Incorrect server. Expected: %s, actual: %s", "", config.Server)
	}
}
