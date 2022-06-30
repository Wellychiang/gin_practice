package router

import (
	"api/admin"
	"api/api"
	"api/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	// 當機時可以恢復
	router.Use(gin.Recovery())

	router.Static("/static", "static")

	// 跨域中間件
	router.Use(middleware.Cors())

	// 日誌中間件
	router.Use(middleware.Logger())

	register(router)

	return router
}

func register(router *gin.Engine) {
	// api
	// 查找blogger
	v1 := router.Group("api/v1")
	v1.GET("/blogger", api.FindBlogger)
	v1.GET("/blog/type", api.FindType)
	v1.GET("/blog/list", api.BlogListWtihType)
	v1.GET("/blog", api.FindBlog)
	v1.PUT("/blog/:id", api.UpdateCommentId)
	v1.POST("/blog/comment", api.PostComment)

	// back
	v1.POST("/login", admin.Login)
	v1.POST("/logout", admin.Logout)
	v1.POST("/register", admin.Register)

	jwt := v1.Group("/admin", middleware.Jwt())
	jwt.GET("/blogger", admin.FindBlogger)
	jwt.PUT("/info", admin.BloggerUpdateInfo) //目前更改 put post 之類的到這
	jwt.PUT("/info/password", admin.BloggerUpdatePassword)
	jwt.POST("/blog", admin.PostBlog)
	jwt.POST("/type/list", admin.TypeList)
	jwt.POST("/type/save", admin.CreateType)
	jwt.POST("/type/delete", admin.DeleteType)
	jwt.PUT("/type/update", admin.UpdateType)

}
