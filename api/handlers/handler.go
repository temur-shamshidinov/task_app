package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/temur-shamshidinov/task_app/storage"
)

type handler struct {
	storage storage.StorageI
}

func NewHandler(storage storage.StorageI) handler {
	return handler{storage: storage}
}

func (h handler) Ping(ctx *gin.Context) {
	ctx.JSON(200, map[string]string{"massage": "pong"})
}
