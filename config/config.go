package config

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Config struct {
	PostgresqlUsername string
	PostgresqlDb       string
	PostgresqlHost     string
	PostgresqlPort     string
}

var GlobalConfig Config

func LoadEnv() {
	file, err := os.Open(".env")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	config := Config{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		fmt.Println(key, " - ", value)

		switch key {
		case "POSTGRESQL_USERNAME":
			config.PostgresqlUsername = value
		case "POSTGRESQL_PORT":
			config.PostgresqlPort = value
		case "POSTGRESQL_HOST":
			config.PostgresqlHost = value
		case "POSTGRESQL_DB":
			config.PostgresqlDb = value
		}
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	GlobalConfig = config
}
