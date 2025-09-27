package router

import (
	"web3-go-blog/handlers"
	"web3-go-blog/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.POST("/register", handlers.Register) // 注册
		api.POST("/login", handlers.Login)       // 登录

		api.GET("/posts", handlers.ListPosts)                // 列出所有文章
		api.GET("/post/:id", handlers.GetPost)               // 获取单篇文章
		api.GET("/post/:id/comments", handlers.ListComments) // 列出文章的评论

		auth := api.Group("")
		auth.Use(middleware.AuthMiddleware()) //其他路径需要登录才能访问
		{
			auth.POST("/post", handlers.CreatePost)                // 创建文章
			auth.PUT("/post/:id", handlers.UpdatePost)             // 更新文章
			auth.DELETE("/post/:id", handlers.DeletePost)          // 删除文章
			auth.POST("/post/:id/comment", handlers.CreateComment) // 创建评论
		}
	}
	return r
}
