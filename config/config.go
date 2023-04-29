package config

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// LoadConfig reading config file and set os envs
func LoadConfig() {
	readConfigFile(configPath())
}

// readConfigFile reading config file
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

// GetWorkdirFromOsEnv return string that contain path to project dir
func GetWorkdirFromOsEnv() string {
	return os.Getenv("WORKDIR")
}

// configPath return string that contain path to config file
func configPath() string {
	workDir := GetWorkdirFromOsEnv()
	dir := filepath.Join(workDir, "config", ".env")
	return dir
}

// parseConfigLine separate file line by = and return key and value
func parseConfigLine(line string) (string, string) {
	configLine := strings.Split(line, "=")
	varr := configLine[0]
	value := strings.ReplaceAll(configLine[1], `"`, "")
	return varr, value
}
