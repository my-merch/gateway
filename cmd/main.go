// api-gateway/main.go
package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()

	// Connect to the gRPC services
	productServiceConn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Product Service: %v", err)
	}
	defer productServiceConn.Close()

	orderServiceConn, err := grpc.Dial("localhost:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Order Service: %v", err)
	}
	defer orderServiceConn.Close()

	userServiceConn, err := grpc.Dial("localhost:8083", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to User Service: %v", err)
	}
	defer userServiceConn.Close()

	// Create gRPC client instances
	//productServiceClient := gateway.NewProductServiceClient(productServiceConn)
	//orderServiceClient := gateway.NewOrderServiceClient(orderServiceConn)
	//userServiceClient := gateway.NewUserServiceClient(userServiceConn)

	// Define routes to forward requests to the microservices
	router.GET("/getProductInfo", func(c *gin.Context) {
		// Forward the request to the Product Service
		// You would typically make a gRPC call here
		//resp, err := productServiceClient.GetProductInfo(context.Background(), &gateway.ProductInfoRequest{})
		//if err != nil {
		//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		//	return
		//}
		//c.JSON(http.StatusOK, resp)
		c.JSON(http.StatusOK, "hola")
	})

	router.POST("/placeOrder", func(c *gin.Context) {
		// Forward the request
		// Forward the request to the Order Service
		// You would typically make a gRPC call here
		//var placeOrderReq gateway.PlaceOrderRequest
		//if err := c.BindJSON(&placeOrderReq); err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		//	return
		//}
		//
		//resp, err := orderServiceClient.PlaceOrder(context.Background(), &placeOrderReq)
		//if err != nil {
		//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		//	return
		//}
		c.JSON(http.StatusOK, "placeOrder")
	})

	router.GET("/getUserProfile", func(c *gin.Context) {
		// Forward the request to the User Service
		// You would typically make a gRPC call here
		//resp, err := userServiceClient.GetUserProfile(context.Background(), &gateway.UserProfileRequest{})
		//if err != nil {
		//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		//	return
		//}
		c.JSON(http.StatusOK, "getUserProfile")
	})

	// Run the API Gateway on port 8080
	router.Run(":8079")
}
