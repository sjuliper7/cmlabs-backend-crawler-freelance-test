package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sjuliper7/cmlabs-backend-crawler-freelance-test/crawler"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! welcome to test cmlabs")
	})

	e.POST("/scraping", crawler.Website)

	e.Logger.Fatal(e.Start(":5111"))
}
