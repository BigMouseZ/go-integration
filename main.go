package main

import (
	"github.com/gin-gonic/gin"
	"log"

	_"go-integration/cache"
)
func main() {
	log.Println("启动程序！")
	r := gin.Default()
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}



