// @author: wzmiiiiii
// @since: 2022/12/26 01:21:28
// @desc: TODO

package service

import (
	"fmt"
	"math/rand"
	"time"

	"wzmiiiiii.cn/gind/demo3/tool"
)

type MemberService struct{}

func (ms *MemberService) SendSms(phone string) bool {
	cfg := tool.GetConfig()

	// 生成验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000000))

	// 调用阿里云sdk 完成发送
	_ = cfg
	_ = code

	// 接收返回结果,并判断发送状态
	return false
}
