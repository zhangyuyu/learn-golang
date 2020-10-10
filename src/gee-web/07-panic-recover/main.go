package main

import (
	"net/http"

	"gee"
)

func main() {
	r := gee.Default()
	r.Use(gee.Logger())
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":8090")
}
