package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// Создаем URL для первого сервера
	url1, err := url.Parse("http://localhost:9001")
	if err != nil {
		log.Fatal(err)
	}

	// Создаем URL для второго сервера
	url2, err := url.Parse("http://localhost:9000")
	if err != nil {
		log.Fatal(err)
	}

	// Создаем балансировщик с двумя серверами
	balancer := httputil.NewSingleHostReverseProxy(&url.URL{})
	balancer.Director = func(req *http.Request) {
		// Используем round-robin балансировку
		if req.URL.Path == "/create" {
			fmt.Println("server 9001")
			req.URL.Scheme = url1.Scheme
			req.URL.Host = url1.Host
		} else {
			fmt.Println("server 9000")

			req.URL.Scheme = url2.Scheme
			req.URL.Host = url2.Host
		}
	}

	// Запускаем сервер на порту 8080
	log.Fatal(http.ListenAndServe(":8080", balancer))
}
