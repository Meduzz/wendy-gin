package main

import (
	"github.com/Meduzz/wendy"
	wendygin "github.com/Meduzz/wendy-gin"
	"github.com/Meduzz/wendy/example/service"
	"github.com/gin-gonic/gin"
)

func main() {
	srv := gin.Default()

	// serve a local module
	logic := wendy.FromModules("example", service.ServiceModule())

	// add static paths and what else is needed for the app

	// register wendy api path
	srv.POST("/api", wendygin.WithWendy(logic))

	// alt. wire each function to it's own path
	// srv.POST("/api/add", wendygin.WendyMethod(service.ServiceModule().Method("add")))
	// srv.POST("/api/remove", wendygin.WendyMethod(service.ServiceModule().Method("remove")))
	// srv.POST("/api/list", wendygin.WendyMethod(service.ServiceModule().Method("list")))
	// srv.POST("/api/find", wendygin.WendyMethod(service.ServiceModule().Method("find")))

	// or bind pretty paths to the function in a module
	srv.POST("/api/add", wendygin.WithWendyOnlyBody("example", "service", "add", logic))
	srv.POST("/api/remove", wendygin.WithWendyOnlyBody("example", "service", "remove", logic))
	srv.POST("/api/list", wendygin.WithWendyOnlyBody("example", "service", "list", logic))
	srv.POST("/api/find", wendygin.WithWendyOnlyBody("example", "service", "find", logic))

	srv.Run(":8080")
}
