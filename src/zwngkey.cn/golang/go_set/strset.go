/*
 * @Author: zwngkey
 * @Date: 2022-05-04 00:49:12
 * @LastEditTime: 2022-05-04 01:04:46
 * @Description:	go实现set集合
 */
/*
	用映射来模拟集合（set）
		Go不支持内置集合（set）类型。但是，集合类型可以轻松使用映射类型来模拟。
			在实践中，我们常常使用映射类型map[K]struct{}来模拟一个元素类型为K的集合类型。
			 类型struct{}的尺寸为零，所以此映射类型的值中的元素不消耗内存。
*/
package goset

import (
	"reflect"
)

// 这里实现的是一个字符串集合，使用空结构体是为了节省内存
type Empty struct{}
type String map[string]Empty

func NewString(items ...string) (ss String) {
	ss = String{}
	ss.Insert(items...)
	return
}

// StringKeySet根据一个map的key来创建字符串集合
// 如何参数不是map类型会panic
func StringKeySet(theMap any) String {
	v := reflect.ValueOf(theMap)
	ret := String{}

	for _, keyValue := range v.MapKeys() {
		ret.Insert(keyValue.Interface().(string))
	}
	return ret
}

func (s String) Insert(items ...string) String {
	for _, item := range items {
		s[item] = Empty{}
	}
	return s
}

// Delete 删除集合所有元素
func (s *String) DeleteAll() String {
	*s = String{}
	return *s
}

// func (s String) DeleteAll() String {
// 	s = String{}
// 	return s
// }
// 通过返回值,set1 = set1.DeleteAll()

// Delete 删除集合指定元素
func (s String) Delete(items ...string) String {
	for _, item := range items {
		delete(s, item)
	}
	return s
}

// Has 返回元素是否在集合中
func (s String) Has(item string) bool {
	_, contained := s[item]
	return contained
}

// HasAll 判断切片中元素是否都在集合中
func (s String) HasAll(items ...string) bool {
	for _, item := range items {
		if !s.Has(item) {
			return false
		}
	}
	return true
}

// HasAny 返回集合中是否包含任意切片中元素
func (s String) HasAny(items ...string) bool {
	for _, item := range items {
		if s.Has(item) {
			return true
		}
	}
	return false
}
