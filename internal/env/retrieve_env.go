package env

import (
	"os"
	"strings"
)

func getEnv(envName string, fallback string) string {
	value, exists := os.LookupEnv(envName)
	if !exists {
		value = fallback
	}
	return value
}

func RetrieveFileName() string {
	playlistName := getEnv("PLAYLIST_NAME", "Rclone")
	return playlistName + ".m3u8"
}

func RetrievePathToScan() string {
	return getEnv("PATH_TO_SCAN", "./")
}

func RetrieveBaseUrl() string {
	return getEnv("BASE_URL", "localhost:8080")
}

func RetrievePathsToExclude() []string {
	value, exists := os.LookupEnv("PATHS_TO_EXCLUDE")
	pathsToExclude := []string{}
	if exists {
		pathsToExclude = strings.Split(value, ",")
	}
	return pathsToExclude
}