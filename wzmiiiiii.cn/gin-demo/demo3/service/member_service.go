// @author: wzmiiiiii
// @since: 2022/12/26 01:21:28
// @desc: TODO

package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"

	"wzmiiiiii.cn/gind/demo3/tool"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
)

type MemberService struct{}

func (ms *MemberService) SendSms(phone string) error {
	// 生成验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000000))

	// 调用阿里云sdk 完成发送
	return sendSms(phone, code)
}

func sendSms(phone, code string) error {
	cfg := tool.GetConfig()
	config := sdk.NewConfig()
	credential := credentials.NewAccessKeyCredential(cfg.Sms.AppKey, cfg.Sms.AppSecret)
	client, err := dysmsapi.NewClientWithOptions("cn-hangzhou", config, credential)
	if err != nil {
		log.Println("短信发送失败")
		return err
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = cfg.Sms.SignName
	request.TemplateCode = cfg.Sms.TemplateCode
	request.PhoneNumbers = phone
	par, err := json.Marshal(map[string]any{
		"code": code,
	})
	request.TemplateParam = string(par)
	response, err := client.SendSms(request)
	if err != nil {
		log.Println("短信发送失败")
		return err
	}
	if response.Code != "OK" {
		return errors.New("短信发送失败")
	}
	return nil
}
