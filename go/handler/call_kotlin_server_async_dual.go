package handler

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"goroutineServer/client"
	"log"
	"net/http"
	"sort"
	"sync"
)

func CallKotlinServerAsyncDual(c *gin.Context) {
	log.Print("[CallKotlinServerAsyncDual] start")

	results := &callResults{}
	tries := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	concurrency := 2
	group := &sync.WaitGroup{}
	guard := make(chan struct{}, concurrency)

	for _, i := range tries {

		group.Add(1)
		guard <- struct{}{}

		go func(i int) {
			log.Print("[CallKotlinServerAsyncDual] before request with id: ", i)

			var result string

			r, err := client.Post(i)
			if err != nil || r == nil {
				result = "failed"
			} else {
				result = r.Result
			}

			log.Print("[CallKotlinServerAsyncDual] after request with id: ", i)

			results.Response = append(results.Response, &client.Response{
				ID:     i,
				Result: result,
			})

			group.Done()
			<-guard
		}(i)
	}

	group.Wait()

	sort.Slice(results.Response, func(i, j int) bool {
		return results.Response[i].ID < results.Response[j].ID
	})

	log.Print("[CallKotlinServerAsyncDual] done")

	c.JSON(http.StatusOK, results)
}

func CallKotlinServerAsyncDualErr(c *gin.Context) {
	log.Print("[CallKotlinServerAsyncDualErr] start")

	results := &callResults{}
	tries := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	concurrency := 2
	group := new(errgroup.Group)
	guard := make(chan struct{}, concurrency)

	for _, i := range tries {

		guard <- struct{}{}
		func(i int) {
			group.Go(func() error {
				defer func() { <-guard }()
				log.Print("[CallKotlinServerAsyncDualErr] before request with id: ", i)

				var result string

				r, err := client.Post(i)
				if err != nil {
					return err
				}

				if r == nil {
					result = "failed"
				} else {
					result = r.Result
				}

				log.Print("[CallKotlinServerAsyncDualErr] after request with id: ", i)

				results.Response = append(results.Response, &client.Response{
					ID:     i,
					Result: result,
				})
				return nil
			})
		}(i)
	}

	if err := group.Wait(); err != nil {
		log.Print("[CallKotlinServerAsyncDualErr] error: ", err)
		return
	}

	sort.Slice(results.Response, func(i, j int) bool {
		return results.Response[i].ID < results.Response[j].ID
	})

	log.Print("[CallKotlinServerAsyncDualErr] done")

	c.JSON(http.StatusOK, results)
}
