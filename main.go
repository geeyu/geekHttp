package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.Get("/", func(c *gee.Context) {
		c.String(http.StatusOK, "URL.Path = %q\n", c.Path)
	})

	r.Get("/hello", func(c *gee.Context) {
		for key, header := range c.Req.Header {
			c.String(http.StatusOK, "Header[%q] = %q\n", key, header)
		}
	})

	r.Post("/login", func(c *gee.Context) {
		c.Json(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
