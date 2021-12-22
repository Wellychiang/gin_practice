package middleware

import (
	"api/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TokenData struct {
	Username string
	Password string
	Key      string
}

func (token *TokenData) SetToken() string {
	return utils.Md5(token.Username + token.Password + token.Key)
}

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token == "" {
			res := &utils.Response{
				Code: 1100,
				Msg:  "請求沒攜帶 token, 無權限訪問",
			}
			res.Json(c)
			c.Abort()
			return
		}
		logrus.Debug("get token:", token)

		// 從快取裡抓取 token
		data, found := utils.Cache.Get(token)
		if found == false {
			res := &utils.Response{
				Code: 1100,
				Msg:  "token 找不到或過期",
			}
			res.Json(c)
			c.Abort()
			return
		}

		tokenData := data.(*TokenData)

		b := token != tokenData.SetToken()

		if b {
			res := &utils.Response{
				Code: 1100,
				Msg:  "token 認證出錯",
			}
			res.Json(c)
			c.Abort()
			return
		}

		c.Set("token", data)

		c.Next()
	}
}
