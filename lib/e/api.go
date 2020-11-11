package e

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wss/lib/utils"
)

// 提供空 obj类型
func GetEmptyStruct() interface{} {
	return struct {
	}{}
}

type ApiPageLists struct {
	Page    int                    `json:"page"`
	Limit   int                    `json:"limit"`
	Total   int                    `json:"total_size"`
	Lists   interface{}            `json:"lists"`
	Map     []utils.WhereData      `json:"map"`
	OptParm map[string]interface{} `json:"opt_parm"`
}

type ApiId struct {
	Id int64 `json:"id"`
}

//检测默认值
func CheckApiPageListDefault(cb *ApiPageLists) {
	if cb.Page <= 0 {
		cb.Page = 1
	}
	if cb.Limit <= 0 {
		cb.Limit = 20
	}
}

type ApiJson struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//返回成功
func ApiOk(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, ApiJson{
		Code: SUCCESS,
		Msg:  msg,
		Data: data,
	})
	c.Abort()
	return
}

// 返回错误
func ApiErr(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, ApiJson{
		Code: ERROR,
		Msg:  msg,
		Data: nil,
	})
	c.Abort()
	return
}

// 返回其他数据类型
func ApiOpt(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, ApiJson{
		Code: code,
		Msg:  msg,
		Data: data,
	})
	c.Abort()
	return
}

// 任意类型转 in特64
func ToInt64(in interface{}) int64 {
	switch in.(type) {
	case string:
		if cb, err := strconv.ParseInt(in.(string), 10, 64); err != nil {
			return 0
		} else {
			return cb
		}
		break
	case int:
		return int64(in.(int))
		break
	case int64:
		return in.(int64)
		break
	case float64:
		return int64(in.(float64))
	case float32:
		return int64(in.(float32))
	}
	return 0
}

// 任意类型转float
func ToFloat64(in interface{}) float64 {
	switch in.(type) {
	case string:
		if cb, err := strconv.ParseFloat(in.(string), 64); err != nil {
			return 0
		} else {
			return cb
		}
		break
	case int:
		return float64(in.(int))
		break
	case int64:
		return float64(in.(int64))
		break
	case float64:
		return in.(float64)
	case float32:
		return float64(in.(float32))
	}
	return 0
}

// 任意类型转string
func ToString(in interface{}) string {
	switch in.(type) {
	case string:
		return in.(string)
		break
	case int:
		return strconv.Itoa(in.(int))
		break
	case int64:
		return strconv.FormatInt(in.(int64), 10)
		break
	case float64:
		return strconv.FormatFloat(in.(float64), 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(in.(float32)), 'f', -1, 64)
	}
	return ""
}
