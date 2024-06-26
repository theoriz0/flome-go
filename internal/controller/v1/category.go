package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/theoriz0/flome-go/internal/service"
)

func AddCategoryRoutes(rg *gin.RouterGroup) {
	posts := rg.Group("/categories")
	posts.GET("", getCategories)
}

func getCategories(c *gin.Context) {
	categories, err := service.GetCategoryList()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Categories": categories,
	})

}
