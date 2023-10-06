package routes

import (
	"crud-api/controllers"

	"github.com/gin-gonic/gin"
)

type PostRouteController struct {
	postController controllers.PostController
}

func NewRoutePostController(postController controllers.PostController) PostRouteController {
	return PostRouteController{postController}
}


func (pc *PostRouteController) PostRoute(rg *gin.RouterGroup) {
	router := rg.Group("posts")
	router.POST("/", pc.postController.CreatePost)
}