package api

import (
	"api/service"
	"api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindBlogger(c *gin.Context) {
	var blogger service.Blogger
	result := blogger.Find()
	result.Password = ""
	res := &utils.Response{Code: 0, Msg: "", Data: result}
	res.Json(c)
}

// 查找 blogger 分類數量
func FindType(c *gin.Context) {
	var blogType service.BlogType
	result := blogType.FindTypeCount()
	res := &utils.Response{Code: 0, Msg: "", Data: result}
	res.Json(c)
}

func BlogList(c *gin.Context) {
	json := make(map[string]interface{})
	err := c.ShouldBind(&json) // 綁定前端傳來的值, 匹配 json 的 map 格式
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "數據格式出錯"}
		res.Json(c)
		return
	}
	blog := new(service.Blog)
	// 字串轉 Int
	page, _ := strconv.Atoi(utils.StrVal(json["page"]))
	size, _ := strconv.Atoi(utils.StrVal(json["size"]))

	strInt64 := strconv.FormatInt(blog.Count(), 10)
	blogCount, _ := strconv.Atoi(strInt64)
	pageVo := &utils.Page{Page: page, Size: size, Total: blogCount}

	typeId, err := strconv.Atoi(utils.StrVal(json["type_id"]))
	if err == nil {
		blog.TypeId = typeId
	}
	// 查詢blog list
	result, err := blog.FindList(pageVo)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: err.Error()}
		res.Json(c)
		return
	}

	res := &utils.Response{Code: 0, Msg: "", Data: result, Count: pageVo.Total}
	res.Json(c)
}
