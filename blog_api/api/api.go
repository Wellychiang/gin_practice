package api

import (
	"api/service"
	"api/utils"
	"fmt"
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

func BlogListWtihType(c *gin.Context) {

	blog := new(service.Blog)
	// 字串轉 Int
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	typeId, err := strconv.Atoi(c.Query("type_id"))
	fmt.Println(page, size)
	if page == 0 || size == 0 {
		res := &utils.Response{Code: 1000, Msg: "error input"}
		res.Json(c)
		return
	}

	strInt64 := strconv.FormatInt(blog.Count(), 10)
	blogCount, _ := strconv.Atoi(strInt64)
	pageVo := &utils.Page{Page: page, Size: size, Total: blogCount}

	if err == nil {
		fmt.Println("nilll~")
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

func FindBlog(c *gin.Context) {
	var blog service.Blog

	blogId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: err.Error()}
		res.Json(c)
		return
	}
	blog.Id = blogId

	fmt.Printf("blog: %+v", blog)
	blog.UpdateClick()
	fmt.Printf("%+v", blog)
	blogContent := blog.FindBlogContent()
	blogComment := blog.FindCommentByBlog()
	next := blog.FindNextOne()
	previous := blog.FindPreviousOne()

	blogInfo := make(map[string]interface{})
	blogInfo["blog_content"] = blogContent
	blogInfo["blog_comment"] = blogComment
	blogInfo["next"] = next
	blogInfo["previous"] = previous

	res := &utils.Response{Code: 0, Msg: "", Data: blogInfo}
	res.Json(c)
}

func UpdateCommentId(c *gin.Context) {
	var comment service.Comment

	// comment_id, id_err := strconv.Atoi(c.Param("id"))
	comment_id, id_err := strconv.ParseUint(c.Param("id"), 10, 64)
	comment_status, status_err := strconv.Atoi(c.Query("status"))

	// err := c.BindJSON(&comment)
	if id_err != nil || status_err != nil {
		res := &utils.Response{Code: 1000, Msg: "數據格式錯誤"}
		res.Json(c)
		return
	}
	comment.Id = comment_id
	comment.Status = comment_status
	result := comment.UpdateStatus()
	if result.Error != nil {
		fmt.Println(result)
		res := &utils.Response{Code: 1000, Msg: "數據格式錯誤"}
		res.Json(c)
		return
	}

	res := &utils.Response{Code: 0, Msg: ""}
	res.Json(c)
}

func PostComment(c *gin.Context) {
	var comment service.Comment
	err := c.BindJSON(&comment)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "數據格式錯誤"}
		res.Json(c)
		return
	}
	//IP
	comment.Ip = c.ClientIP()
	comment.AddTime = utils.GetDate(utils.DateFormat)
	comment.Status = 1
	comment.Insert()

	blog := &service.Blog{Id: comment.BlogId}
	fmt.Println(comment)
	fmt.Println(blog)
	blog.UpdateReplay()
	res := &utils.Response{Code: 0, Msg: "success"}
	res.Json(c)
}
