package main

import (
	"github.com/labstack/echo/v4"
	"ilteoood/local_m3u8/internal"
)

const START_PORT = ":3000"

func main() {
	echoServer := echo.New()
	echoServer.GET("/playlist/generate", internal.GeneratePlaylist)
	echoServer.GET("/playlist", internal.SendPlaylist)
	echoServer.Start(START_PORT)
}