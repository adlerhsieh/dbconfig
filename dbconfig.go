package dbconfig

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Development struct {
		username string
		password string
	}
}

func ReadFile(filepath string, env string) map[string]string {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	data := string(file)

	m := make(map[interface{}]map[string]string)
	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatal(err)
	}

	return m[env]
}
