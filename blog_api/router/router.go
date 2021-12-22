package router

import (
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

	// 日治中間件
	router.Use(middleware.Logger())

	register(router)

	return router
}

func register(router *gin.Engine) {

}
