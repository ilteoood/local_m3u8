package internal

import (
	"ilteoood/local_m3u8/internal/env"
	"ilteoood/local_m3u8/internal/playlist"
	"io/fs"
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

func walker(playlist *playlist.Playlist) fs.WalkDirFunc {
	pathsToExclude := env.RetrievePathsToExclude()

	return func(path string, d fs.DirEntry, err error) error {
		canBeInserted := !d.IsDir() && !isExcludedPath(pathsToExclude, path)
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

func savePlaylist (playlistPath string, playlist *playlist.Playlist) {
	file, _ := os.Create(retrievePlaylistPath())
	file.WriteString(playlist.Content)
}

func GeneratePlaylist(echoContext echo.Context) error {
	pathToScan := env.RetrievePathToScan()
	playlist := playlist.Playlist{}
	playlist.AddPlaylistHeader()
	filepath.WalkDir(pathToScan, walker(&playlist))
	playlistPath := retrievePlaylistPath()
	savePlaylist(playlistPath, &playlist)
	return echoContext.Attachment(playlistPath, env.RetrieveFileName())
}