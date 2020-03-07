package app

import "github.com/tv2169145/golang-http-client-mock/controllers"

func MappingUrls() {
	router.POST("/get-user", controllers.GetUser)
}
