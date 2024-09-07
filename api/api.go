package api

import (
	"github.com/gin-gonic/gin"
	"github.com/temur-shamshidinov/task_app/api/handlers"
	"github.com/temur-shamshidinov/task_app/storage"
	"github.com/temur-shamshidinov/task_app/middleware" 
)

func Api(storage storage.StorageI) {
	
	router := gin.Default()

	
	h := handlers.NewHandler(storage)

	
	router.GET("/ping", h.Ping)

	
	router.POST("/register", h.Register)
	router.POST("/login", h.Login)

	
	taskRoutes := router.Group("/tasks", middleware.AuthMiddleware())
	{
		taskRoutes.POST("/create-task", h.CreateTask)
		taskRoutes.GET("/get-task", h.GetTasks)
		taskRoutes.PUT("/update-task/:id", h.UpdateTask)
		taskRoutes.DELETE("/delete-task/:id", h.DeleteTask)
	}

	
	router.Run()
}
