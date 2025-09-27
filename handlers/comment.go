package handlers

import (
	"net/http"
	"strconv"
	"web3-go-blog/middleware"
	"web3-go-blog/models"

	"github.com/gin-gonic/gin"
)

type CreateCommentRequest struct {
	Content string `json:"content"`
}

func CreateComment(c *gin.Context) {
	var req CreateCommentRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "invalid request"})
		return
	}
	uid := c.MustGet(middleware.CtxUserID).(uint)
	comment := models.Comment{Content: req.Content, UserID: uid}
	idstr, _ := strconv.Atoi(c.Param("id"))
	comment.PostID = uint(idstr)
	if err := models.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "message": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "message": "success", "data": comment})
}

func ListComments(c *gin.Context) {
	var comments []models.Comment
	models.DB.Where("post_id = ?", c.Param("id")).Find(&comments)
	c.JSON(http.StatusOK, gin.H{"code": 1, "data": comments})
}
