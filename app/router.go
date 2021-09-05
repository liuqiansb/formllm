package main

import (
	"formllm/dao"
	"formllm/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	s *service.Service
)

func InitRouter(d dao.DBInterface) (server *http.Server) {
	s = &service.Service{d}
	router := gin.Default()
	// Simple group: v1
	v1 := router.Group("/user")
	{
		v1.POST("/create", createUser)
	}
	server = &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	return
}
