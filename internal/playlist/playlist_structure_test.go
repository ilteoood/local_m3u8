package playlist

import (
	"os"
	"testing"
)

func TestAddInformation (test *testing.T) {
	playlist := Playlist{}
	information := playlist.AddInformation("/path/to/base/file.mp4").content
	if information != "#EXTINF:-1 tvg-name=\"file.mp4\", file.mp4\n" {
		test.Errorf("Generated wrong information: %s", information)
	}
}

func TestAddInformationWithSpaces (test *testing.T) {
	playlist := Playlist{}
	information := playlist.AddInformation("/path/to/base/file/with lots of spaces.mp4").content
	if information != "#EXTINF:-1 tvg-name=\"with lots of spaces.mp4\", with lots of spaces.mp4\n" {
		test.Errorf("Generated wrong information: %s", information)
	}
}

func TestAddGroup (test *testing.T) {
	playlist := Playlist{}
	group := playlist.AddGroup("/path/to/base/file/with lots of spaces.mp4").content
	if group != "#EXTGRP:/path/to/base/file\n" {
		test.Errorf("Generated wrong group: %s", group)
	}
}

func TestAddGroupStripBase (test *testing.T) {
	playlist := Playlist{}
	os.Setenv("PATH_TO_SCAN", "/path/to")
	group := playlist.AddGroup("/path/to/base/file/with lots of spaces.mp4").content
	if group != "#EXTGRP:/base/file\n" {
		test.Errorf("Generated wrong group: %s", group)
	}
	os.Clearenv()
}

func TestAddFile (test *testing.T) {
	playlist := Playlist{}
	file := playlist.AddFile("/path/to/base/file/with lots of spaces.mp4").content
	if file != "localhost:8080/path/to/base/file/with lots of spaces.mp4\n" {
		test.Errorf("Generated wrong file: %s", file)
	}
}

func TestAddGroupBaseUrl (test *testing.T) {
	playlist := Playlist{}
	os.Setenv("BASE_URL", "test.com")
	file := playlist.AddFile("/path/to/base/file/with lots of spaces.mp4").content
	if file != "test.com/path/to/base/file/with lots of spaces.mp4\n" {
		test.Errorf("Generated wrong file: %s", file)
	}
	os.Clearenv()
}

func TestAddGroupBaseUrlStripped (test *testing.T) {
	playlist := Playlist{}
	os.Setenv("BASE_URL", "test.com")
	os.Setenv("PATH_TO_SCAN", "/path/to")
	file := playlist.AddFile("/path/to/base/file/with lots of spaces.mp4").content
	if file != "test.com/base/file/with lots of spaces.mp4\n" {
		test.Errorf("Generated wrong file: %s", file)
	}
	os.Clearenv()
}