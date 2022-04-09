package admin

import (
	"api/service"
	"api/utils"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func TypeList(c *gin.Context) {
	var page utils.Page
	err := c.BindJSON(&page)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "數據格式錯誤"}
		res.Json(c)
		return
	}
	blogType := &service.BlogType{}

	countString := strconv.FormatInt(blogType.Count(), 10)
	page.Total, _ = strconv.Atoi(countString)

	ListData := blogType.FindList(&page)
	res := &utils.Response{Code: 0, Msg: "", Count: page.Total, Data: ListData}
	res.Json(c)
}

func CreateType(c *gin.Context) {
	var blogType service.BlogType
	err := c.BindJSON(&blogType)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "數據格式錯誤"}
		res.Json(c)
		return
	}
	result := blogType.Insert()
	if result.Error != nil {
		res := &utils.Response{Code: 1000, Msg: "提交出錯"}
		res.Json(c)
		return
	}
	res := &utils.Response{Code: 0, Msg: "Create type success"}
	res.Json(c)
}

func DeleteType(c *gin.Context) {
	var blogType service.BlogType
	err := c.BindJSON(&blogType)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "數據格式錯誤"}
		res.Json(c)
		return
	}
	blog := new(service.Blog)
	blog.TypeId = blogType.Id
	// blogCount := blog.FindById()
	countString := strconv.FormatInt(blog.FindById(), 10)
	blogCount, _ := strconv.Atoi(countString)
	fmt.Println(blogCount)

	if blogCount > 0 {
		res := &utils.Response{Code: 1000, Msg: "此分類下有文章, 請先刪除相關文章"}
		res.Json(c)
		return
	}

	result := blogType.Delete()
	if result.Error != nil {
		res := &utils.Response{Code: 1000, Msg: "提交出錯"}
		res.Json(c)
		return
	}
	res := &utils.Response{Code: 0, Msg: ""}
	res.Json(c)
}

func UpdateType(c *gin.Context) {
	var blogType service.BlogType
	err := c.BindJSON(&blogType)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "數據格式錯誤"}
		res.Json(c)
		return
	}
	result := blogType.Update()
	if result.Error != nil {
		res := &utils.Response{Code: 1000, Msg: "提交出錯"}
		res.Json(c)
		return
	}
	res := &utils.Response{Code: 0, Msg: "Update type success"}
	res.Json(c)
}
