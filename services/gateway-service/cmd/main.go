package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func proxyWithPrefix(targetURL string, prefix string) gin.HandlerFunc {
	url, _ := url.Parse(targetURL)
	return func(c *gin.Context) {
		proxy := httputil.NewSingleHostReverseProxy(url)
		proxy.Director = func(req *http.Request) {
			req.Header.Set("X-Forwarded-Host", req.Host)
			req.Header.Set("X-Forwarded-Proto", "http")
			req.URL.Scheme = url.Scheme
			req.URL.Host = url.Host
			req.URL.Path = strings.TrimPrefix(c.Request.URL.Path, prefix)
		}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func proxyRaw(targetURL string) gin.HandlerFunc {
	url, _ := url.Parse(targetURL)
	return func(c *gin.Context) {
		proxy := httputil.NewSingleHostReverseProxy(url)
		proxy.Director = func(req *http.Request) {
			req.Header.Set("X-Forwarded-Host", req.Host)
			req.URL.Scheme = url.Scheme
			req.URL.Host = url.Host
			req.URL.Path = c.Request.URL.Path
		}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}


func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Any("/auth", proxyWithPrefix("http://auth-service:8001", "/auth"))
	r.Any("/auth/*any", proxyWithPrefix("http://auth-service:8001", "/auth"))

	r.Any("/designs", proxyRaw("http://design-service:8002"))
	r.Any("/designs/*any", proxyRaw("http://design-service:8002"))

	r.GET("/public/designs", proxyWithPrefix("http://supplier-service:8004", "/public/designs"))

	r.Any("/quotes", proxyWithPrefix("http://order-service:8003", "/quotes"))
	r.Any("/quotes/*any", proxyWithPrefix("http://order-service:8003", "/quotes"))

	r.Any("/suppliers", proxyWithPrefix("http://supplier-service:8004", "/suppliers"))
	r.Any("/suppliers/*any", proxyWithPrefix("http://supplier-service:8004", "/suppliers"))

	r.Any("/payments", proxyWithPrefix("http://payment-service:8005", "/payments"))
	r.Any("/payments/*any", proxyWithPrefix("http://payment-service:8005", "/payments"))

	r.GET("/uploads/*filepath", gin.WrapH(http.FileServer(http.Dir("/app/uploads"))))

	log.Println("✅ API Gateway started on :8000")
	log.Println("   → Auth:        /auth")
	log.Println("   → Designs:     /designs")
	log.Println("   → Public:      /public/designs")
	log.Println("   → Quotes:      /quotes")
	log.Println("   → Suppliers:   /suppliers")
	log.Println("   → Payments:    /payments")
	log.Println("   → Uploads:     /uploads/...")

	r.Run(":8000")
}