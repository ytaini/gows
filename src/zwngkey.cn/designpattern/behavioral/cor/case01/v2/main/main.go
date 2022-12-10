package main

import (
	"fmt"
	"strings"
)

func main() {
	// 假如我们需要对请求的数据进行过滤.

	// 数据
	info := NewInfo("张三", "大家好:),<script> 今天吃饭了吗? 欢迎访问 zwngkey.cn, 你现在还是996吗?")

	// 处理
	filterChain := &FilterChain{}
	filterChain.Add(&HTMLFilter{}).Add(&SensitiveFilter{})

	filterChain1 := &FilterChain{}
	filterChain1.Add(filterChain, &FaceFilter{}, &AddrFilter{})
	filterChain1.DoFilter(info)

	fmt.Println(info.Msg)
}

type FilterChain struct {
	filters []Filter
}

func (fc *FilterChain) Add(filters ...Filter) *FilterChain {
	fc.filters = append(fc.filters, filters...)
	return fc
}

func (fc *FilterChain) DoFilter(info *Info) bool {
	for _, filter := range fc.filters {
		if !filter.DoFilter(info) { // 只要某个过滤器返回false.剩下的过滤器就都不执行了.
			return false
		}
	}
	return true
}

type HTMLFilter struct{}

func (f *HTMLFilter) DoFilter(info *Info) bool {
	info.Msg = strings.Replace(info.Msg, "<", "[", -1)
	info.Msg = strings.Replace(info.Msg, ">", "]", -1)
	return true
}

type SensitiveFilter struct{}

func (f *SensitiveFilter) DoFilter(info *Info) bool {
	info.Msg = strings.Replace(info.Msg, "996", "995", -1)
	return false
}

type FaceFilter struct{}

func (f *FaceFilter) DoFilter(info *Info) bool {
	info.Msg = strings.Replace(info.Msg, ":)", "^-^", -1)
	return true
}

type AddrFilter struct{}

func (f *AddrFilter) DoFilter(info *Info) bool {
	info.Msg = strings.Replace(info.Msg, "zwngkey.cn", "wzmiiiiii.cn", -1)
	return true
}

type Filter interface {
	DoFilter(info *Info) bool
}

type Info struct {
	Msg  string
	Name string
}

func NewInfo(name, msg string) *Info {
	return &Info{msg, name}
}
