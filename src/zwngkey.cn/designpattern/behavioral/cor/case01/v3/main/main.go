package main

import "fmt"

func main() {
	// 假如我们需要对请求的数据进行过滤.

	// 数据
	req := &Request{str: "Request->"}
	res := &Response{str: "Response->"}

	// 处理
	filterChain := &FilterChain{}
	filterChain.Add(&HTMLFilter{}, &SensitiveFilter{})
	filterChain.DoFilter(req, res, filterChain)

	fmt.Println(req)
	fmt.Println(res)
	//输出:
	//htmlFilter 处理请求...
	//SensitiveFilter 处理请求...
	//SensitiveFilter 处理响应...
	//htmlFilter 处理响应...
	//&{Request->htmlFilter handle requestSensitiveFilter handle request}
	//&{Response->SensitiveFilter handle responsehtmlFilter handle response}
}

type FilterChain struct {
	filters []Filter
	index   int
}

func (fc *FilterChain) Add(filters ...Filter) *FilterChain {
	fc.filters = append(fc.filters, filters...)
	return fc
}

func (fc *FilterChain) DoFilter(req *Request, res *Response, chain *FilterChain) bool {
	if len(fc.filters) == fc.index {
		return true
	}
	filter := fc.filters[fc.index]
	fc.index++
	return filter.DoFilter(req, res, chain)
}

type HTMLFilter struct{}

func (f *HTMLFilter) DoFilter(req *Request, res *Response, chain *FilterChain) bool {
	fmt.Println("htmlFilter 处理请求...")
	req.str = req.str + "htmlFilter handle request"
	if !chain.DoFilter(req, res, chain) {
		return false
	}
	fmt.Println("htmlFilter 处理响应...")
	res.str = res.str + "htmlFilter handle response"
	return true
}

type SensitiveFilter struct{}

func (f *SensitiveFilter) DoFilter(req *Request, res *Response, chain *FilterChain) bool {
	fmt.Println("SensitiveFilter 处理请求...")
	req.str = req.str + "SensitiveFilter handle request"
	if !chain.DoFilter(req, res, chain) {
		return false
	}
	fmt.Println("SensitiveFilter 处理响应...")
	res.str = res.str + "SensitiveFilter handle response"
	return true
}

type Filter interface {
	DoFilter(req *Request, res *Response, chain *FilterChain) bool
}

type Request struct {
	str string
}

type Response struct {
	str string
}
