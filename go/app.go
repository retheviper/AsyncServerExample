package main

import (
	"github.com/gin-gonic/gin"
	"goroutineServer/handler"
)

func main() {
	router := gin.Default()
	router.Group("/api/v1").
		GET("/call-kotlin-server", handler.CallKotlinServer).
		GET("/call-kotlin-server-async", handler.CallKotlinServerAsync).
		GET("/call-kotlin-server-async-dual", handler.CallKotlinServerAsyncDual).
		GET("/call-kotlin-server-async-dual-err", handler.CallKotlinServerAsyncDualErr).
		POST("/some-process", handler.SomeProcess)
	err := router.Run(":8800")
	if err != nil {
		return
	}
}
