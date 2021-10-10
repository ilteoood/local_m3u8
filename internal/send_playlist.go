package internal

import (
	"ilteoood/local_m3u8/internal/env"

	"github.com/labstack/echo/v4"
)

func SendPlaylist(echoContext echo.Context) error {
	return echoContext.Attachment(env.RetrievePlaylistPath(), env.RetrieveFileName())
}
