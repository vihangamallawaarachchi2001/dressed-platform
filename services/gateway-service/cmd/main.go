package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Auth
	authURL, _ := url.Parse("http://auth-service:8001")
	r.Any("/auth/*path", proxy(authURL))

	// Design
	designURL, _ := url.Parse("http://design-service:8002")
	r.Any("/designs/*path", proxy(designURL))
	r.GET("/uploads/*filepath", gin.WrapH(http.FileServer(http.Dir("/app/uploads"))))

	// Quotes / Orders
	orderURL, _ := url.Parse("http://order-service:8003")
	r.Any("/quotes/*path", proxy(orderURL))

	// Suppliers
	supplierURL, _ := url.Parse("http://supplier-service:8004")
	r.Any("/suppliers/*path", proxy(supplierURL))
	r.GET("/designs", proxy(supplierURL)) // public design list

	// Payments
	paymentURL, _ := url.Parse("http://payment-service:8005")
	r.Any("/payments/*path", proxy(paymentURL))

	log.Println("🚀 API Gateway running on :8000")
	r.Run(":8000")
}

func proxy(target *url.URL) gin.HandlerFunc {
	return func(c *gin.Context) {
		proxy := httputil.NewSingleHostReverseProxy(target)
		proxy.Director = func(req *http.Request) {
			req.Host = target.Host
			req.URL.Scheme = target.Scheme
			req.URL.Host = target.Host
			req.URL.Path = c.Param("path")
		}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}