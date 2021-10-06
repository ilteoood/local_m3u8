package playlist

import (
	"os"
	"testing"
)

func TestAddPlaylistHeader (test *testing.T) {
	playlist := Playlist{}
	header := playlist.AddPlaylistHeader().Content
	if header != "#EXTM3U\n#PLAYLIST:Rclone.m3u8\n" {
		test.Errorf("Generated wrong header: %s", header)
	}
}

func TestAddPlaylistHeaderCustomPlaylistName (test *testing.T) {
	playlist := Playlist{}
	os.Setenv("PLAYLIST_NAME", "TestPlaylist")
	header := playlist.AddPlaylistHeader().Content
	if header != "#EXTM3U\n#PLAYLIST:TestPlaylist.m3u8\n" {
		test.Errorf("Generated wrong header: %s", header)
	}
}

func TestAddInformation (test *testing.T) {
	playlist := Playlist{}
	information := playlist.AddInformation("/path/to/base/file.mp4").Content
	if information != "#EXTINF:-1 tvg-name=\"file.mp4\", file.mp4\n" {
		test.Errorf("Generated wrong information: %s", information)
	}
}

func TestAddInformationWithSpaces (test *testing.T) {
	playlist := Playlist{}
	information := playlist.AddInformation("/path/to/base/file/with lots of spaces.mp4").Content
	if information != "#EXTINF:-1 tvg-name=\"with lots of spaces.mp4\", with lots of spaces.mp4\n" {
		test.Errorf("Generated wrong information: %s", information)
	}
}

func TestAddGroup (test *testing.T) {
	playlist := Playlist{}
	group := playlist.AddGroup("/path/to/base/file/with lots of spaces.mp4").Content
	if group != "#EXTGRP:/path/to/base/file\n" {
		test.Errorf("Generated wrong group: %s", group)
	}
}

func TestAddGroupStripBase (test *testing.T) {
	playlist := Playlist{}
	os.Setenv("PATH_TO_SCAN", "/path/to")
	group := playlist.AddGroup("/path/to/base/file/with lots of spaces.mp4").Content
	if group != "#EXTGRP:/base/file\n" {
		test.Errorf("Generated wrong group: %s", group)
	}
	os.Clearenv()
}

func TestAddFile (test *testing.T) {
	playlist := Playlist{}
	file := playlist.AddFile("/path/to/base/file/with lots of spaces.mp4").Content
	if file != "localhost:8080/path/to/base/file/with lots of spaces.mp4\n" {
		test.Errorf("Generated wrong file: %s", file)
	}
}

func TestAddGroupBaseUrl (test *testing.T) {
	playlist := Playlist{}
	os.Setenv("BASE_URL", "test.com")
	file := playlist.AddFile("/path/to/base/file/with lots of spaces.mp4").Content
	if file != "test.com/path/to/base/file/with lots of spaces.mp4\n" {
		test.Errorf("Generated wrong file: %s", file)
	}
	os.Clearenv()
}

func TestAddGroupBaseUrlStripped (test *testing.T) {
	playlist := Playlist{}
	os.Setenv("BASE_URL", "test.com")
	os.Setenv("PATH_TO_SCAN", "/path/to")
	file := playlist.AddFile("/path/to/base/file/with lots of spaces.mp4").Content
	if file != "test.com/base/file/with lots of spaces.mp4\n" {
		test.Errorf("Generated wrong file: %s", file)
	}
	os.Clearenv()
}

func TestAddEntry (test *testing.T) {
	playlist := Playlist{}
	os.Setenv("BASE_URL", "test.com")
	os.Setenv("PATH_TO_SCAN", "/path/to")
	file := playlist.AddNewEntry("/path/to/base/file/with lots of spaces.mp4").Content
	expectedContent := "#EXTINF:-1 tvg-name=\"with lots of spaces.mp4\", with lots of spaces.mp4\n" +
		"#EXTGRP:/path/to/base/file\n" +
		"test.com/base/file/with lots of spaces.mp4\n"
	if file != expectedContent {
		test.Errorf("Generated wrong file: %s", file)
	}
	os.Clearenv()
}