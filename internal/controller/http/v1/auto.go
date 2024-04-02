package v1

import "github.com/gin-gonic/gin"

func newAutoRoutes(handler *gin.RouterGroup) {
	h := handler.Group("/cars")
	{
		h.GET("/info/:id", Info)
	}
}

// Info ShowAccount
// @Summary      Show an account
// @Description  get cars by ID
// @Tags         auto
// @Accept       json
// @Produce      json
// @Param name path string true "The name to say hello to" example(Henry)
// @Success      200
// @Router       /auto/info/{id} [get]
func Info(c *gin.Context) {
	c.JSON(200, gin.H{"message": c.Param("id")})
}
