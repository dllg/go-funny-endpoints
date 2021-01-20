package main

import (
	"github.com/dllg/go-funny-endpoints/router"
)

func main() {
	r := router.Setup()
	r.Run() // listen and serve on 0.0.0.0:$PORT (8080 by default)
}
