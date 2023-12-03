package internal

import (
	"ilteoood/local_m3u8/internal/env"
	"ilteoood/local_m3u8/internal/playlist"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

func isExcludedPath(pathsToExclude []string, path string) bool {
	for _, pathToExclude := range pathsToExclude {
		if strings.Contains(path, pathToExclude) {
			return true
		}
	}
	return false
}

func isSupportedExtension(supportedExtensions []string, path string) bool {
	for _, supportedExtension := range supportedExtensions {
		if strings.HasSuffix(path, supportedExtension) {
			return true
		}
	}
	return false
}

func walker(playlist *playlist.Playlist) fs.WalkDirFunc {
	pathsToExclude := env.RetrievePathsToExclude()
	supportedExtensions := env.RetrieveSupportedExtensions()

	return func(path string, d fs.DirEntry, err error) error {
		canBeInserted := !d.IsDir() && !isExcludedPath(pathsToExclude, path) && isSupportedExtension(supportedExtensions, path)
		if canBeInserted {
			playlist.AddNewEntry(path)
		}
		return nil
	}
}

func savePlaylist(playlist *playlist.Playlist) error {
	file, error := os.Create(env.RetrievePlaylistPath())
	if error == nil {
		defer file.Close()
		file.WriteString(playlist.Content)
	}
	return error
}

func GeneratePlaylist(echoContext echo.Context) error {
	pathToScan := env.RetrievePathToScan()
	playlist := playlist.Playlist{}
	playlist.AddPlaylistHeader()
	filepath.WalkDir(pathToScan, walker(&playlist))
	saveError := savePlaylist(&playlist)
	if saveError == nil {
		return echoContext.Attachment(env.RetrievePlaylistPath(), env.RetrieveFileName())
	}

	return echoContext.NoContent(http.StatusNoContent)
}
