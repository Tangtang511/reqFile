package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	var router = gin.Default()
	//handlers.HttpTest(router)
	router.Run(":8080")
}
