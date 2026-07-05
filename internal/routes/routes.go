package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/toanlamt/go-crud-assignment/internal/handlers"
)

func RegisterRoutes(router *gin.Engine, handler *handlers.ProductHandler) {
	router.POST("/products", handler.Create)

	router.GET("/products", handler.GetAll)

	router.GET("/products/:id", handler.GetByID)

	router.PUT("/products/:id", handler.Update)

	router.DELETE("/products/:id", handler.Delete)

	router.GET("/health", func(c *gin.Context) {
    c.JSON(200, gin.H{
        "status": "ok",
    })
})
}