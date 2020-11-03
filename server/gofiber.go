package main

import (
	"fmt"
	"net/url"

	"github.com/gofiber/fiber/v2"
	proxy "github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/valyala/fasthttp"
	fasthttpproxy "github.com/valyala/fasthttp/fasthttpproxy"
)

func local_fasthttp(target string) fiber.Handler {
	proxyUrl, _ := url.Parse(target)
	client := &fasthttp.Client{
		Dial: fasthttpproxy.FasthttpHTTPDialer(proxyUrl.Host),
	}
	return func(c *fiber.Ctx) error {
		fmt.Println(c.Path())
		req := fasthttp.AcquireRequest()
		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseRequest(req)
		defer fasthttp.ReleaseResponse(resp)
		req.SetRequestURI(c.Path())
		err := client.Do(req, resp)
		if err != nil {
			fmt.Println(err)
		}
		return err
	}
}

func local_fiber(target string) fiber.Handler {
	proxyUrl, _ := url.Parse(target)
	return proxy.Balancer(proxy.Config{Servers: []string{proxyUrl.Host}})
}

func runFiber() error {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("test")
	})
	app.Group("/apis", local_fasthttp("http://127.0.0.1:8888"))
	app.Static("/", "web")
	return app.Listen(":" + *port)
}
