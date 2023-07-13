package crawler

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/labstack/echo/v4"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	htmlTemplate = `
    <!DOCTYPE html>
    <html>
        <head>
            <title>Output</title>
        </head>
        <body>
               {{range .}}
        		<p>{{.}}</p>
				  <br>
              {{end}}
        </body>
    </html>
`
)

type crawlerRequest struct {
	URLs []string `json:"urls"`
}

// Website is function to in crawler package that do scrap to get all the element text from a page
func Website(c echo.Context) error {
	var crawlerRequest crawlerRequest

	if err := c.Bind(&crawlerRequest); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	collector := colly.NewCollector()

	for _, url := range crawlerRequest.URLs {
		var elementText []string
		log.Printf("start scrapping url %s", url)
		collector.OnRequest(func(r *colly.Request) {
			// Check if the request is for a SPA, SSR, or PWA website
			if r.Headers.Get("X-Frame-Options") == "DENY" {
				// Don't send the request
				r.Abort()
			}
		})

		collector.OnHTML("*", func(e *colly.HTMLElement) {
			elementText = append(elementText, e.Text)
		})

		collector.Visit(url)
		log.Printf("scrap in process")
		collector.Wait()
		log.Printf("scrap %s is done", url)

		log.Printf("write the result into file")
		t := template.Must(template.New("").Parse(htmlTemplate))

		// Create a new HTML file
		splitUrls := strings.Split(url, "/")

		fileName := fmt.Sprintf("output_%d_%s.html", time.Now().Unix(), strings.Replace(strings.Join(splitUrls, ""), ":", "", -1))
		f, err := os.Create(filepath.Join("outputs", fileName))
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusOK, fmt.Sprintf("error: %v", err))
		}

		// Write the HTML to the file
		err = t.Execute(f, elementText)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusOK, fmt.Sprintf("error: %v", err))
		}
	}

	return c.JSON(http.StatusOK, "success!")
}
