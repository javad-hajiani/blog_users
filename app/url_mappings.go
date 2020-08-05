package app

import (
	"github.com/javad-hajiani/blog_users/controllers/ping"
	"github.com/javad-hajiani/blog_users/controllers/users"
)

func urlMappings() {
	router.GET("/ping/", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users/", users.CreateUser)
}
