package admin

import (
	"api/service"
	"api/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func PostBlog(c *gin.Context) {
	var blog service.Blog
	err := c.BindJSON(&blog)
	fmt.Printf("blog: %+v", blog)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "數據格式錯誤"}
		res.Json(c)
	}

	blog.Insert()
	res := &utils.Response{Code: 0, Msg: ""}
	res.Json(c)
}
