package env

import (
	"os"
	"testing"
)

func TestRetrieveFallbackFileName(test *testing.T) {
	fileName := RetrieveFileName()
	if fileName != "Rclone.m3u8" {
		test.Errorf("Fallback name not valid: %s", fileName)
	}
}

func TestRetrieveFallbackClearedFileName(test *testing.T) {
	os.Clearenv()
	fileName := RetrieveFileName()
	if fileName != "Rclone.m3u8" {
		test.Errorf("Fallback name not valid: %s", fileName)
	}
}

func TestRetrieveRightFileName(test *testing.T) {
	os.Setenv("PLAYLIST_NAME", "MyName")
	fileName := RetrieveFileName()
	if fileName != "MyName.m3u8" {
		test.Errorf("Right file name not valid: %s", fileName)
	}
}

func TestRetrieveFallbackPathToScan(test *testing.T) {
	pathToScan := RetrievePathToScan()
	if pathToScan != "./" {
		test.Errorf("Invalid path to scan fallback: %s", pathToScan)
	}
}

func TestRetrieveClearedPathToScan(test *testing.T) {
	os.Clearenv()
	pathToScan := RetrievePathToScan()
	if pathToScan != "./" {
		test.Errorf("Invalid path to scan fallback: %s", pathToScan)
	}
}

func TestRetrieveRightPathToScan(test *testing.T) {
	os.Setenv("PATH_TO_SCAN", "MyName")
	pathToScan := RetrievePathToScan()
	if pathToScan != "MyName" {
		test.Errorf("Invalid path to scan: %s", pathToScan)
	}
}

func TestRetrieveFallbackBaseUrl(test *testing.T) {
	baseUrl := RetrieveBaseUrl()
	if baseUrl != "localhost:8080" {
		test.Errorf("Invalid base url fallback: %s", baseUrl)
	}
}

func TestRetrieveClearedFallbackBaseUrl(test *testing.T) {
	os.Clearenv()
	baseUrl := RetrieveBaseUrl()
	if baseUrl != "localhost:8080" {
		test.Errorf("Invalid base url fallback: %s", baseUrl)
	}
}

func TestRetrieveRightFallbackBaseUrl(test *testing.T) {
	os.Setenv("BASE_URL", "test.com")
	baseUrl := RetrieveBaseUrl()
	if baseUrl != "test.com" {
		test.Errorf("Invalid base url: %s", baseUrl)
	}
}

func TestRetrieveFallbackPathsToExclude(test *testing.T) {
	pathsToExclude := RetrievePathsToExclude()
	if len(pathsToExclude) != 0 {
		test.Errorf("Invalid paths to exclude fallback: %+q", pathsToExclude)
	}
}

func TestRetrieveClearedPathsToExclude(test *testing.T) {
	os.Clearenv()
	pathsToExclude := RetrievePathsToExclude()
	if len(pathsToExclude) != 0 {
		test.Errorf("Invalid paths to exclude fallback: %+q", pathsToExclude)
	}
}

func TestRetrieveRightPathsToExclde(test *testing.T) {
	os.Setenv("PATHS_TO_EXCLUDE", "/hello,/foo")
	pathsToExclude := RetrievePathsToExclude()
	if len(pathsToExclude) != 2 || pathsToExclude[0] != "/hello" || pathsToExclude[1] != "/foo" {
		test.Errorf("Invalid paths to exclude: %+q", pathsToExclude)
	}
}

func TestRetrieveFallbackSupportedExtensions(test *testing.T) {
	supportedExtensions := RetrieveSupportedExtensions()
	if len(supportedExtensions) == 0 {
		test.Errorf("Invalid supported extensions fallback: %+q", supportedExtensions)
	}
}

func TestRetrieveClearedSupportedExtensions(test *testing.T) {
	os.Clearenv()
	supportedExtensions := RetrieveSupportedExtensions()
	if len(supportedExtensions) == 0 {
		test.Errorf("Invalid supported extensions fallback: %+q", supportedExtensions)
	}
}

func TestRetrieveRightSupportedExtensions(test *testing.T) {
	os.Setenv("SUPPORTED_EXTENSIONS", ".mp4,.mkv")
	supportedExtensions := RetrieveSupportedExtensions()
	if len(supportedExtensions) != 2 || supportedExtensions[0] != ".mp4" || supportedExtensions[1] != ".mkv" {
		test.Errorf("Invalid supported extensions: %+q", supportedExtensions)
	}
}

