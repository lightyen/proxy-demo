package main

import (
	"encoding/json"
	"flag"
	"net"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func monitor() gin.HandlerFunc {
	type Response struct {
		Data int64 `json:"data"`
	}
	hub := melody.New()
	hub.HandleConnect(func(s *melody.Session) {
		data, _ := json.Marshal(&Response{Data: time.Now().Unix()})
		_ = s.Write(data)
	})
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for now := range ticker.C {
			data, _ := json.Marshal(&Response{Data: now.Unix()})
			_ = hub.Broadcast(data)
		}
	}()
	return func(c *gin.Context) {
		_ = hub.HandleRequest(c.Writer, c.Request)
	}
}

var Port = flag.String("port", "8888", "port")

func main() {
	e := gin.Default()
	e.GET("/apis/hello", func(c *gin.Context) {
		c.String(200, "helloworld")
	})
	e.GET("apis/ws", monitor())
	_ = e.Run(net.JoinHostPort("127.0.0.1", *Port))
}
