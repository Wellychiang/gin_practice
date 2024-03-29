package utils

import (
	"api/cache"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"os"
	"strconv"
	"time"
)

// var Cache *cache.Cache = cache.New(time.Duration(conf.Conf.Cache.Expire)*time.Second, time.Duration(conf.Conf.Cache.Clearup)*time.Second)
// 上面的要想要使用 config, 就要先字串轉換為 int
var Cache *cache.Cache = cache.New(time.Duration(7200)*time.Second, time.Duration(600)*time.Second)

const DateFormat = "2020-01-02 15:03:00"

func GetDate(format string) string {
	return time.Now().Format(format)
}

func CreateDir(filePath string) error {
	if !IsExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

func IsExist(path string) bool {
	_, err := os.Stat(path) //os.Stat 獲取文件訊息
	if err != nil {
		if os.IsExist(err) {
			return false
		}
	}
	return true
}

func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func StrVal(value interface{}) string {
	var key string
	if value == nil {
		return key
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}
