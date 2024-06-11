package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`    // status code
	Message string      `json:"message"` // thông báo lỗi
	Data    interface{} `json:"data"`    // dữ liệu return
}

// success response

func SuccessRespone(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: Msg[code],
		Data:    data,
	},
	)
}

func ErrorRespone(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: Msg[code],
		Data:    nil,
	},
	)
}
