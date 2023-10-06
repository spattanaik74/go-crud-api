package controllers

import (
	"net/http"
	"strings"

	"time"

	"crud-api/models"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)


type PostController struct {
	DB *gorm.DB
}


func NewPostController (DB *gorm.DB) PostController {
	return PostController{DB}
}


func (pc *PostController) CreatePost(ctx *gin.Context) {
	// currentUser := ctx.MustGet("currentUser").(models.User)

	var payload *models.CreatePostRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error)
	}

	now := time.Now()

	newPost := models.Post{
		Title: payload.Title,
		Content: payload.Content,
		Image: payload.Image,
		// User: currentUser.ID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := pc.DB.Create(&newPost)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Post with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"Status": "sucess", "data": newPost})
}