package router

import (
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
	v1 := router.Group("v1")
	v1.GET("/blogger", api.FindBlogger)
	v1.GET("/blog/type", api.FindType)
	v1.GET("/blog/list", api.BlogList)

}
