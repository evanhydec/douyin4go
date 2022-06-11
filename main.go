package main

import (
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	initRouter(r)

	if err := r.Run(); err != nil {
		utils.LogrusObj.Info(err)
		panic(err)
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
