package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type processRequest struct {
	ID int `json:"id"`
}

type processReponse struct {
	ID     int    `json:"id"`
	Result string `json:"result"`
}

// SomeProcess do some process for 5 seconds
func SomeProcess(c *gin.Context) {
	var request processRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var result string
	if request.ID%2 != 0 {
		result = "success"
	} else {
		result = "something went wrong"
	}

	time.Sleep(5 * time.Second)

	c.JSON(http.StatusOK, processReponse{
		ID:     request.ID,
		Result: result,
	})
}
