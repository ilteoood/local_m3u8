package main

import (
	"ilteoood/local_m3u8/internal"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const START_PORT = ":3000"

func main() {
	echoServer := echo.New()

	echoServer.Use(middleware.Logger())
	echoServer.Use(middleware.Recover())

	echoServer.GET("/playlist/generate", internal.GeneratePlaylist)
	echoServer.GET("/playlist", internal.SendPlaylist)
	echoServer.Start(START_PORT)
}
