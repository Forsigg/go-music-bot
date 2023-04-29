package config

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func LoadConfig() {
	readConfigFile(configPath())
}

func readConfigFile(configPath string) {
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("Error when open config .env file: %s", err)
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		key, value := parseConfigLine(fileScanner.Text())
		err = os.Setenv(key, value)
		if err != nil {
			log.Fatalf("Error while setting in os env key=%s value=%s: %s", key, value, err)
		}
	}
	if err = fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading config .env file: %s", err)
	}
}

func GetWorkdirFromOsEnv() string {
	return os.Getenv("WORKDIR")
}

func configPath() string {
	workDir := GetWorkdirFromOsEnv()
	dir := filepath.Join(workDir, "config", ".env")
	return dir
}

func parseConfigLine(line string) (string, string) {
	configLine := strings.Split(line, "=")
	varr := configLine[0]
	value := strings.ReplaceAll(configLine[1], `"`, "")
	return varr, value
}
