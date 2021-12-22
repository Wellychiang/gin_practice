package middleware

import (
	"api/utils"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	logger := utils.Log()

	return func(c *gin.Context) {
		// 開始時間
		startTime := time.Now()

		// 處理請求
		c.Next()

		// 結束時間
		endTime := time.Now()

		// 執行時間
		latencyTime := endTime.Sub(startTime) / time.Millisecond

		// 請求方式
		reqMethod := c.Request.Method

		// 請求路由
		reqUri := c.Request.RequestURI

		header := c.Request.Header
		proto := c.Request.Proto

		// 狀態馬
		statusCode := c.Writer.Status()

		// 請求IP
		clientIP := c.ClientIP()

		err := c.Err()

		body, _ := ioutil.ReadAll(c.Request.Body)

		logger.WithFields(logrus.Fields{
			"start_time":   startTime,
			"latency_time": latencyTime,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
			"header":       header,
			"proto":        proto,
			"status_code":  statusCode,
			"client_ip":    clientIP,
			"err":          err,
			"body":         body,
		}).Info()
	}
}
