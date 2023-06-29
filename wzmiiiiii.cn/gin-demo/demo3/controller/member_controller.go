// @author: wzmiiiiii
// @since: 2022/12/26 00:53:55
// @desc: TODO

package controller

import (
	"net/http"

	"wzmiiiiii.cn/gind/demo3/service"

	"wzmiiiiii.cn/gind/demo3/tool"

	"github.com/gin-gonic/gin"
)

type MemberController struct{}

func (mc *MemberController) Router(r *gin.RouterGroup) {
	r.GET("/sendsms", mc.sendSmsCode)
}

// http://localhost:8080/api/sendsms?phone=电话号码
func (mc *MemberController) sendSmsCode(c *gin.Context) {
	phone, exist := c.GetQuery("phone")
	if !exist {
		c.JSON(http.StatusOK, tool.NewResponse(tool.Fail, nil, "缺少查询参数: phone"))
		return
	}
	ms := &service.MemberService{}
	if err := ms.SendSms(phone); err == nil {
		c.JSON(http.StatusOK, tool.NewResponse(tool.Success, nil, "发送成功"))
	} else {
		c.JSON(http.StatusOK, tool.NewResponse(tool.Fail, nil, "发送失败"))
	}
}
