package handler

import (
	"github.com/gin-gonic/gin"
	"goroutineServer/client"
	"log"
	"net/http"
)

// callResults is a response of callKotlinServer
type callResults struct {
	Response []*client.Response `json:"response"`
}

func CallKotlinServer(c *gin.Context) {
	log.Print("[CallKotlinServer] start")

	results := &callResults{}
	tries := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, i := range tries {
		log.Print("[CallKotlinServer] before request with id: ", i)

		var result string

		r, err := client.Post(i)
		if err != nil || r == nil {
			result = "failed"
		} else {
			result = r.Result
		}

		log.Print("[CallKotlinServer] after request with id: ", i)

		results.Response = append(results.Response, &client.Response{
			ID:     i,
			Result: result,
		})
	}

	log.Print("[CallKotlinServer] done")

	c.JSON(http.StatusOK, results)
}
