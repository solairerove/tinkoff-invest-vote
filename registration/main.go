package main

import (
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	load "github.com/solairerove/tinkoff-invest-vote/file"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {

	for _, user := range load.ReadFromFile() {
		email := user.Email
		password := user.Password

		log.Println("trying to registrate new email: " + email)

		// create a new collector
		c := colly.NewCollector(
			colly.Debugger(&debug.LogDebugger{}),
			colly.AllowURLRevisit())

		c.WithTransport(&http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		})

		// Rotate two socks5 proxies
		// rp, err := proxy.RoundRobinProxySwitcher("socks5://147.135.116.172:1871")
		// if err != nil {
		// log.Fatal(err)
		// }
		// c.SetProxyFunc(rp)

		c.OnRequest(func(r *colly.Request) {
			r.Headers.Set("User-Agent", RandomString())
		})

		// "assign":                strings.ReplaceAll(uuid.NewString(), "-", ""),
		err := c.Post("https://invest-terminal.useresponse.com/registration",
			map[string]string{
				"email":                 email,
				"full_name":             password,
				"password":              password,
				"password_confirmation": password,
				"submit":                "Регистрация"})

		if err != nil {
			log.Fatal(err)
		}

		err = c.Request("GET", "https://invest-terminal.useresponse.com/logout", nil, nil, nil)

		if err != nil {
			log.Fatal(err)
		}

		c.OnError(func(_ *colly.Response, err error) {
			log.Println("Something went wrong:", err)
		})

		c.OnResponseHeaders(func(r *colly.Response) {
			log.Println("Visited", r.Request.URL)
		})

		c.OnResponse(func(r *colly.Response) {
			log.Println("Visited", r.Request.URL)
		})
	}
}
