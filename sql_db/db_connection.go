package sql_db

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Connections []DbConnection `yaml:"connections"`
}

type DbConnection struct {
	Server   string `yaml:"server"`
	Database string `yaml:"database"`
	Env      string `yaml:"env"`
	Uid      string `yaml:"uid"`
	Pwd      string `yaml:"pwd"`
}

func (c *DbConnection) GetConnectionString() string {
	query := url.Values{}
	query.Add("app name", "sql-reports-go")
	query.Add("database", c.Database)
	query.Add("encrypt", "disable")

	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(c.Uid, c.Pwd),
		Host:     fmt.Sprintf("%s:%d", c.Server, 1433),
		RawQuery: query.Encode(),
	}

	return u.String()
}

func GetDbConnections(configPath *string) []DbConnection {
	file, err := os.Open(*configPath)

	if err != nil {
		log.Fatalf("Cannot open file: %v", err)
	}

	defer file.Close()

	decoder := yaml.NewDecoder(file)

	var config Config
	err = decoder.Decode(&config)

	if err != nil {
		log.Fatalf("Cannot decode file: %v", err)
	}

	return config.Connections
}
