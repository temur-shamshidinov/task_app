package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temur-shamshidinov/task_app/models"
)

func (h *handler) CreateTask(c *gin.Context) {

	var task models.Task

	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Foydalanuvchi autentifikatsiya qilinmagan"})
		return
	}

	
	task.UserID = int(userID.(float64)) 

	if err := h.storage.GetTaskRepo().CreateTask(context.Background(), task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Task yaratildi"})
}

func (h *handler) GetTasks(c *gin.Context) {


	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Foydalanuvchi autentifikatsiya qilinmagan"})
		return
	}

	tasks, err := h.storage.GetTaskRepo().GetTasks(context.Background(), int(userID.(float64)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *handler) UpdateTask(c *gin.Context) {
	
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID noto'g'ri formatda"})
		return
	}

	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.storage.GetTaskRepo().UpdateTask(context.Background(), id, task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Task muvaffaqiyatli yangilandi"})
}

func (h *handler) DeleteTask(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID noto'g'ri formatda"})
		return
	}

	if err := h.storage.GetTaskRepo().DeleteTask(context.Background(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Task muvaffaqiyatli o'chirildi"})
}
