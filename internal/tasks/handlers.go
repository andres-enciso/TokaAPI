package tasks

import (
	"net/http"
	"strconv"

	"TokaAPI/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type createTaskDTO struct {
	Titulo     string `json:"titulo" binding:"required"`
	Completada *bool  `json:"completada"`
}
type updateTaskDTO struct {
	Titulo     *string `json:"titulo"`
	Completada *bool   `json:"completada"`
}

func RegisterRoutes(g *gin.RouterGroup, db *gorm.DB) {
	g.POST("/", func(c *gin.Context) {
		var in createTaskDTO
		if err := c.ShouldBindJSON(&in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		t := models.Task{Titulo: in.Titulo}
		if in.Completada != nil {
			t.Completada = *in.Completada
		}
		if err := db.Create(&t).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
			return
		}
		c.JSON(http.StatusCreated, t)
	})

	g.GET("/", func(c *gin.Context) {
		var items []models.Task
		if err := db.Order("id desc").Find(&items).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
			return
		}
		c.JSON(http.StatusOK, items)
	})

	g.GET("/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var t models.Task
		if err := db.First(&t, id).Error; err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, t)
	})

	g.PUT("/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var t models.Task
		if err := db.First(&t, id).Error; err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		var in updateTaskDTO
		if err := c.ShouldBindJSON(&in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if in.Titulo != nil {
			t.Titulo = *in.Titulo
		}
		if in.Completada != nil {
			t.Completada = *in.Completada
		}
		if err := db.Save(&t).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
			return
		}
		c.JSON(http.StatusOK, t)
	})

	g.DELETE("/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := db.Delete(&models.Task{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
			return
		}
		c.Status(http.StatusNoContent)
	})
}
