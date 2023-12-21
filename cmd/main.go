package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pbUser "github.com/my-merch/gateway/stubs/user-service/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func proxy(c *gin.Context) {
	remote, err := url.Parse("http://localhost:8083")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host

		// Extract the proxy path from the request and set it as the new path
		req.URL.Path = c.Param("proxyPath")
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}

func main() {
	router := gin.Default()

	userServiceConn, err := grpc.Dial("localhost:8083", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to User Service: %v", err)
	}
	defer userServiceConn.Close()

	ctx := context.Background()
	mux := runtime.NewServeMux()

	// Register gRPC endpoint to reverse proxy
	err = pbUser.RegisterUserServiceHandler(ctx, mux, userServiceConn)
	if err != nil {
		log.Fatalf("Failed to register UserService handler: %v", err)
	}

	// Use the reverse proxy as a handler for Gin
	router.Any("/*proxyPath", proxy)

	// Run the API Gateway on port 8079
	router.Run(":8079")
}
