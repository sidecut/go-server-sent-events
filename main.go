package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.Use(cors.Default())
	server.Use(static.Serve("/", static.LocalFile("./static", true)))
	server.Use(static.Serve("/js", static.LocalFile("./static", true)))
	server.GET("/progress", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Flush()

		time.Sleep(time.Second)

		const max = 20
		for i := 0; i < max; i++ {
			c.Writer.WriteString(fmt.Sprintf("id: %d\n", i))
			c.Writer.WriteString("event: onProgress\n")
			data, _ := json.Marshal(gin.H{
				"progressPercentage": i + 1,
			})
			c.Writer.WriteString(fmt.Sprintf("data: %s\n", data))
			c.Writer.WriteString("\n")
			c.Writer.Flush()

			time.Sleep(time.Second / 5)
		}

		c.Writer.WriteString(fmt.Sprintf("id: %d\n", max))
		c.Writer.WriteString("event: onComplete\n")
		c.Writer.WriteString("data: {}\n")
		c.Writer.WriteString("\n")
		c.Writer.Flush()
	})

	if err := server.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}
