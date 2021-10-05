package playlist

import (
	"fmt"
	"ilteoood/local_m3u8/internal/env"
	"path/filepath"
	"strings"
)

type Playlist struct {
	content string
}

func stripPathToScan(fileCompletePath string) string {
	pathToScan := env.RetrievePathToScan()
	return strings.Replace(fileCompletePath, pathToScan, "", 1)
}

func (playlist Playlist) AddPlaylistHeader() Playlist {
	playlistName := fmt.Sprintf("#PLAYLIST:%s", env.RetrieveFileName())
	return playlist.AddRow("#EXTM3U").AddRow(playlistName)
}

func (playlist Playlist) AddInformation(fileCompletePath string) Playlist {
	fileName := filepath.Base(fileCompletePath)
	informationWithPrefix := fmt.Sprintf("#EXTINF:-1 tvg-name=\"%s\", %s", fileName, fileName)
	return playlist.AddRow(informationWithPrefix)
}

func (playlist Playlist) AddGroup(fileCompletePath string) Playlist {
	relativePath := filepath.Dir(stripPathToScan(fileCompletePath))
	groupWithPrefix := fmt.Sprintf("#EXTGRP:%s", relativePath)
	return playlist.AddRow(groupWithPrefix)
}

func (playlist Playlist) AddFile(fileCompletePath string) Playlist {
	baseUrl := env.RetrieveBaseUrl()
	fileRelativePath := stripPathToScan(fileCompletePath)
	return playlist.AddRow(baseUrl + fileRelativePath)
}

func (playlist Playlist) AddRow(rowContent string) Playlist {
	playlist.content = playlist.content + rowContent + "\n"
	return playlist
}
