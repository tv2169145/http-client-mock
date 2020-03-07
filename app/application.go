package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApplication() {
	MappingUrls()
	router.Run(":8080")
}
