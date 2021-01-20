package router

import (
	"github.com/dllg/go-funny-endpoints/funny"
	"github.com/dllg/go-funny-endpoints/httpclient"
	"github.com/gin-gonic/gin"
)

type msgfunc func(hc httpclient.HTTPClient) string

func ginMessage(c *gin.Context, f msgfunc) {
	c.JSON(200, gin.H{
		"message": f(&hc),
	})
}

var (
	hc httpclient.Impl
)

// Setup will setup all endpoints handling different http requests
func Setup() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/advice", func(c *gin.Context) {
			ginMessage(c, funny.GetAdviceFromAPI)
		})
		v1.GET("/chucknorris", func(c *gin.Context) {
			ginMessage(c, funny.GetChuckNorrisJokeFromAPI)
		})
		v1.GET("/dadjoke", func(c *gin.Context) {
			ginMessage(c, funny.GetDadJokeFromAPI)
		})
		v1.GET("/random", func(c *gin.Context) {
			ginMessage(c, funny.GetRandomMessage)
		})
	}
	return r
}
