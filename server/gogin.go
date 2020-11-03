package main

import (
	"fmt"
	"net/http/httputil"
	"net/url"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func local_gin(target string) gin.HandlerFunc {
	url, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(url)
	return func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path)
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func runGin() error {
	e := gin.Default()
	e.NoRoute(
		static.Serve("/", static.LocalFile("web", false)),
	)
	a := e.Group("/apis", local_gin("http://127.0.0.1:8888"))
	a.Any("/*any")
	return e.Run(":" + *port)
}
