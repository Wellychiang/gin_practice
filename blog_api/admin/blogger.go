package admin

import (
	"api/cache"
	"api/middleware"
	"api/service"
	"api/utils"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	var blogger service.Blogger
	err := c.BindJSON(&blogger)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "數據格式錯誤"}
		res.Json(c)
		return
	}
	if blogger.Username == "" || blogger.Password == "" {
		res := &utils.Response{Code: 1001, Msg: "please input username or password"}
		res.Json(c)
		return
	}
	result := blogger.SearchUser()

	if result.Id <= 0 {
		res := &utils.Response{Code: 1001, Msg: "username not found"}
		res.Json(c)
		return
	}

	if result.Password != utils.Md5(blogger.Password) {
		res := &utils.Response{Code: 1002, Msg: "password incorrect"}
		res.Json(c)
		return
	}

	//存入緩存
	key := strconv.Itoa(time.Now().Nanosecond())

	token := &middleware.TokenData{
		Username: utils.Md5(result.Username),
		Password: utils.Md5(result.Password),
		Key:      key,
	}
	tokenKey := token.SetToken()
	utils.Cache.Set(tokenKey, token, cache.DefaultExpiration)

	data := make(map[string]interface{})
	data["userid"] = result.Id
	res := &utils.Response{Code: 0, Msg: "", Data: data, Token: tokenKey}
	res.Json(c)
}

func Logout(c *gin.Context) {
	token := c.GetHeader("token")
	utils.Cache.Delete(token)

	res := &utils.Response{Code: 0, Msg: "Logout success"}
	res.Json(c)
}

func Register(c *gin.Context) {
	var blogger service.Blogger
	err := c.BindJSON(&blogger)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "數據格式錯誤"}
		res.Json(c)
		return
	}

	if blogger.Username == "" || blogger.Password == "" {
		res := &utils.Response{Code: 1001, Msg: "please input username or password"}
		res.Json(c)
		return
	}
	if len(blogger.Username) > 13 || len(blogger.Password) > 13 {
		res := &utils.Response{Code: 1001, Msg: "username or password's limit string is 13"}
		res.Json(c)
		return
	}

	result := blogger.SearchUser()
	if result.Username == blogger.Username {
		res := &utils.Response{Code: 1001, Msg: "Username already used"}
		res.Json(c)
		return
	}

	blogger.Password = utils.Md5(blogger.Password)
	fmt.Println(blogger)
	fmt.Println(blogger.Password)
	blogger.Register()

	res := &utils.Response{Code: 1000, Msg: "Register success"}
	res.Json(c)

}

func FindBlogger(c *gin.Context) {
	var blogger service.Blogger
	// result := blogger.Find()

	bloggerName := c.Query("username")
	// err := c.BindJSON(&blogger)

	blogger.Username = bloggerName
	result := blogger.SearchUser()
	if result == nil {
		res := &utils.Response{Code: 1000, Msg: "找不到blogger"}
		res.Json(c)
		return
	}
	res := &utils.Response{Code: 0, Msg: "", Data: result}
	res.Json(c)
}

func BloggerUpdateInfo(c *gin.Context) {
	var blogger service.Blogger

	err := c.BindJSON(&blogger)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "數據格式錯誤"}
		res.Json(c)
		return
	}
	nickName := blogger.Nickname
	sign := blogger.Sign
	profile := blogger.Profile

	userInfo := blogger.SearchUser()
	if userInfo.Id <= 0 {
		res := &utils.Response{Code: 1000, Msg: "Error username"}
		res.Json(c)
		return
	}

	if nickName == "" && sign == "" && profile == "" {
		res := &utils.Response{Code: 1000, Msg: "please input something"}
		res.Json(c)
		return
	}

	blogger.Nickname = nickName
	blogger.Sign = sign
	blogger.Profile = profile

	result := blogger.UpdateInfo()

	if result.Error != nil {
		res := &utils.Response{Code: 1000, Msg: "提交出錯"}
		res.Json(c)
		return
	}
	res := &utils.Response{Code: 0, Msg: ""}
	res.Json(c)

}

func BloggerUpdatePassword(c *gin.Context) {
	var blogger service.Blogger
	type inputInfo struct {
		Old_password string `json:"old_password"`
		New_password string `json:"new_password"`
		Username     string `json:"username"`
	}
	input_info := inputInfo{Old_password: "", New_password: "", Username: ""}
	err := c.BindJSON(&input_info)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "數據格式錯誤"}
		res.Json(c)
		return
	}
	old_password := input_info.Old_password
	new_password := input_info.New_password
	blogger.Username = input_info.Username

	userInfo := blogger.SearchUser()
	if userInfo.Id <= 0 {
		res := &utils.Response{Code: 1000, Msg: "Error username"}
		res.Json(c)
		return
	}
	blogger.Password = input_info.New_password

	if old_password == "" {
		res := &utils.Response{Code: 1001, Msg: "please input your old password"}
		res.Json(c)
		return
	}
	if old_password == new_password {
		res := &utils.Response{Code: 1001, Msg: "old and new password can not be the same"}
		res.Json(c)
		return
	}
	old_password = utils.Md5(input_info.Old_password)

	if blogger.Password == "" {
		res := &utils.Response{Code: 1001, Msg: "please input your new password"}
		res.Json(c)
		return
	} else if old_password != userInfo.Password {
		res := &utils.Response{Code: 1001, Msg: "Error old password"}
		res.Json(c)
		return
	}
	blogger.Password = utils.Md5(blogger.Password)
	// if blogger.Password != "" {
	// 	blogger.Password = utils.Md5(blogger.Password)
	// }
	fmt.Println(blogger.Password)
	fmt.Println(blogger.Id)
	fmt.Println(blogger.Username)

	var result *gorm.DB
	if blogger.Id <= 0 {
		result = blogger.UpdatePassword()
	} else {
		res := &utils.Response{Code: 1003, Msg: "Id incorrect"}
		res.Json(c)
		return
	}

	if result.Error != nil {
		res := &utils.Response{Code: 1000, Msg: "提交出錯"}
		res.Json(c)
		return
	}
	res := &utils.Response{Code: 0, Msg: "change password success"}
	res.Json(c)

}
