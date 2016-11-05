package dbconfig_test

import (
	"testing"

	"github.com/adlerhsieh/dbconfig"
)

func TestReadFile(t *testing.T) {
	config := dbconfig.ReadFile("./example/database.yml", "development")
	username := "foo"
	password := "bar"
	if config["username"] != username {
		t.Error("Expecting config username as " + username + ". Got " + config["username"])
	}
	if config["password"] != password {
		t.Error("Expecting config username as " + password + ". Got " + config["password"])
	}
}

func TestEnvironment(t *testing.T) {
	config := dbconfig.ReadFile("./example/database.yml", "production")
	host := "the_host"
	if config["host"] != host {
		t.Error("Expecting config host as " + host + ". Got " + config["host"])
	}
}
