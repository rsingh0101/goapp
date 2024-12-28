package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Kafka struct {
		Brokers  string `yaml:"brokers"`
		ClientID string `yaml:"clientID"`
		Topic    string `yaml:"topic"`
		Message  string `yaml:"message"`
	} `yaml:"kafka"`
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		DBName   string `yaml:"dbName"`
	} `yaml:"database"`
	KeyDB struct {
		Path string `yaml:"path"`
	} `yaml:"keydb"`
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
