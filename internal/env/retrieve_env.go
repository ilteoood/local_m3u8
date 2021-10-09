package env

import (
	"os"
	"path"
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

func RetrievePlaylistPath() string {
	return path.Join(RetrievePathToScan(), RetrieveFileName())
}

func retrieveSplittedEnv(envName string, fallback []string) []string {
	value, exists := os.LookupEnv(envName)
	if exists {
		return strings.Split(value, ",")
	}
	return fallback
}

func RetrievePathsToExclude() []string {
	return retrieveSplittedEnv("PATHS_TO_EXCLUDE", []string{})
}

func RetrieveSupportedExtensions() []string {
	return retrieveSplittedEnv("SUPPORTED_EXTENSIONS", []string{".avi", ".mkv", ".mp4"})
}
