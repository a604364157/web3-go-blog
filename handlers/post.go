package handlers

import (
	"net/http"
	"time"
	"web3-go-blog/middleware"
	"web3-go-blog/models"

	"github.com/gin-gonic/gin"
)

type CreatePostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func CreatePost(c *gin.Context) {
	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uid := c.MustGet(middleware.CtxUserID).(uint)
	post := models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  uid,
	}
	if err := models.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"code": 1, "data": post})
}

func ListPosts(c *gin.Context) {
	var posts []models.Post
	if err := models.DB.Preload("Comments").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "data": posts})
}

func GetPost(c *gin.Context) {
	var post models.Post
	if err := models.DB.First(&post, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "data": post})
}

type UpdatePostRequest struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
}

func UpdatePost(c *gin.Context) {
	var post models.Post
	if err := models.DB.First(&post, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "message": "post not found"})
		return
	}
	uid := c.MustGet(middleware.CtxUserID).(uint)
	if post.UserID != uid {
		c.JSON(http.StatusForbidden, gin.H{"code": -1, "message": "permission denied"})
		return
	}
	var req UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updates := map[string]interface{}{"updated_at": time.Now()}
	if req.Title != nil {
		updates["title"] = *req.Title
	}
	if req.Content != nil {
		updates["content"] = *req.Content
	}
	if err := models.DB.Model(&post).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "message": "update failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "data": post})
}

func DeletePost(c *gin.Context) {
	var post models.Post
	if err := models.DB.First(&post, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "message": "post not found"})
		return
	}
	uid := c.MustGet(middleware.CtxUserID).(uint)
	if post.UserID != uid {
		c.JSON(http.StatusForbidden, gin.H{"code": -1, "message": "permission denied"})
		return
	}
	if err := models.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "message": "delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "message": "delete success"})
}
