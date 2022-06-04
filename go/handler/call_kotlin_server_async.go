package handler

import (
	"github.com/gin-gonic/gin"
	"goroutineServer/client"
	"log"
	"net/http"
	"sort"
	"sync"
)

func CallKotlinServerAsync(c *gin.Context) {
	log.Print("[CallKotlinServerAsync] start")

	results := &callResults{}
	tries := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	group := &sync.WaitGroup{}

	for _, i := range tries {
		group.Add(1)

		go func(i int) {
			log.Print("[CallKotlinServerAsync] before request with id: ", i)

			var result string

			r, err := client.Post(i)
			if err != nil || r == nil {
				result = "failed"
			} else {
				result = r.Result
			}

			log.Print("[CallKotlinServerAsync] after request with id: ", i)

			results.Response = append(results.Response, &client.Response{
				ID:     i,
				Result: result,
			})

			group.Done()
		}(i)
	}

	group.Wait()

	sort.Slice(results.Response, func(i, j int) bool {
		return results.Response[i].ID < results.Response[j].ID
	})

	log.Print("[CallKotlinServerAsync] done")

	c.JSON(http.StatusOK, results)
}
