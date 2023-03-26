package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/shuhaib-kv/csv-go.git/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/", controllers.Upload)

	r.POST("/", controllers.Display)

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/static", "./static")
	err := r.Run(":8087")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
