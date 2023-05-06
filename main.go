package main

import (
	"bytedanceYouthCampSimpleDouyinBackend/config"
	"bytedanceYouthCampSimpleDouyinBackend/router"
	"fmt"
)

func main() {
	r := router.InitDouyinRouter()
	err := r.Run(fmt.Sprintf(":%d", config.Info.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		return
	}
}
