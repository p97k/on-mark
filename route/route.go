package route

import (
	"github.com/gin-gonic/gin"
	"github.com/p97k/on-mark/controller"
)

func UserRoutes(routes *gin.Engine) {
	routes.POST("/users/signup", controller.SignUp())
	routes.POST("/users/login", controller.Login())
	routes.POST("/admin/add-product", controller.ProductViewerAdmin())

	routes.GET("/users/product-view", controller.SearchProduct())
	routes.GET("/users/search", controller.SearchProductByQuery())
}
