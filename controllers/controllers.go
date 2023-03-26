package controllers

import (
	"encoding/csv"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {

	c.HTML(http.StatusOK, "upload.html", gin.H{})

}
func Display(c *gin.Context) {
	file, _, err := c.Request.FormFile("csv")
	if err != nil {
		c.String(http.StatusBadRequest, "Error getting file: %v", err)
		return
	}
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

	c.HTML(http.StatusOK, "tables.html", gin.H{
		"headers": headers,
		"rows":    records,
	})
	defer file.Close()

}
