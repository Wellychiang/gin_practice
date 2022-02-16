package admin

import (
	"api/service"
	"api/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var blogger service.Blogger
	err := c.Bind(&blogger)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "數據格式出錯"}
		res.Json(c)
		return
	}

}
