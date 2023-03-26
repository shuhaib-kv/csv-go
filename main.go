package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router.
	r := gin.Default()

	// Set up a route to handle GET requests to the root URL.
	r.GET("/", func(c *gin.Context) {
		// Render an HTML form that allows the user to upload a CSV file.
		c.HTML(http.StatusOK, "upload.html", gin.H{})
	})

	// Set up a route to handle POST requests to the root URL.
	r.POST("/", func(c *gin.Context) {
		// Get the uploaded CSV file from the form data.
		file, _, err := c.Request.FormFile("csv")
		if err != nil {
			c.String(http.StatusBadRequest, "Error getting file: %v", err)
			return
		}
		defer file.Close()

		// Parse the CSV data into a slice of maps.
		reader := csv.NewReader(file)
		records := make([]map[string]string, 0)
		headers := make([]string, 0)
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				c.String(http.StatusBadRequest, "Error reading CSV: %v", err)
				return
			}
			if len(headers) == 0 {
				headers = record
			} else {
				row := make(map[string]string)
				for i, value := range record {
					if i < len(headers) {
						row[headers[i]] = value
					}
				}
				records = append(records, row)
			}
		}

		// Render an HTML table with the CSV data.
		c.HTML(http.StatusOK, "tables.html", gin.H{
			"headers": headers,
			"rows":    records,
		})
	})

	// Serve the HTML templates and static files.
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/static", "./static")

	// Start the server on port 8080.
	err := r.Run(":8087")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
