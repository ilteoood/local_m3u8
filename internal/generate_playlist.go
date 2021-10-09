package internal

import (
	"fmt"
	"ilteoood/local_m3u8/internal/env"
	"ilteoood/local_m3u8/internal/playlist"
	"io/fs"
	"net/http"
	"os"
	"path"
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
	for _, supportedExtension := range(supportedExtensions) {
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

func retrievePlaylistPath () string {
	pathToScan := env.RetrievePathToScan()
	fileName := env.RetrieveFileName()
	return path.Join(pathToScan, fileName)
}

func savePlaylist (playlistPath string, playlist *playlist.Playlist) error {
	file, error := os.Create(retrievePlaylistPath())
	if error == nil {
		defer file.Close()
		file.WriteString(playlist.Content)
	}
	return error
}

func enrichHeaders(echoContext echo.Context) {
	header := echoContext.Response().Header()
	header.Set("Content-Type", "application/x-mpegURL")
	header.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", env.RetrieveFileName()))
}

func GeneratePlaylist(echoContext echo.Context) error {
	enrichHeaders(echoContext)
	pathToScan := env.RetrievePathToScan()
	playlist := playlist.Playlist{}
	playlist.AddPlaylistHeader()
	filepath.WalkDir(pathToScan, walker(&playlist))
	playlistPath := retrievePlaylistPath()
	saveError := savePlaylist(playlistPath, &playlist)
	if saveError != nil {
		return echoContext.Attachment(playlistPath, env.RetrieveFileName())
	}
	return echoContext.JSON(http.StatusNoContent, saveError)
}