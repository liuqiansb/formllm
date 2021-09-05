package main

import (
	"formllm/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func createUser(c *gin.Context) {
	params := model.User{}
	c.ShouldBindJSON(&params)
	id, err := s.CreateUser(c, &params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &model.Result{
			Code: 500,
			Data: nil,
			Msg:  err.Error(),
		})
	}

	c.JSON(http.StatusOK, &model.Result{
		Code: 200,
		Data: id,
		Msg:  "ok",
	})
}
