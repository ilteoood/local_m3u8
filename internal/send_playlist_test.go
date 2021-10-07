package internal

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestGetPlaylist(test *testing.T) {
	echo := echo.New()
	request := httptest.NewRequest(http.MethodPost, "/playlist/generate", nil)
	recorder := httptest.NewRecorder()
	echoContext := echo.NewContext(request, recorder)

	playlist := GeneratePlaylist(echoContext)

	if playlist != nil {
		test.Error("Error during playlist generation")
	}
}
