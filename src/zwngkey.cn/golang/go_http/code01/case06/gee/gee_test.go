/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-29 18:46:09
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-29 19:13:01
 */
package gee

import (
	"fmt"
	"testing"
)

func TestNestedGroup(t *testing.T) {
	r := New()
	v1 := r.Group("/v1")
	v2 := v1.Group("/v2")
	v3 := v2.Group("/v3")

	if v2.prefix != "/v1/v2" {
		t.Fatal("v2 prefix should be /v1/v2")
	}
	if v3.prefix != "/v1/v2/v3" {
		t.Fatal("v2 prefix should be /v1/v2")
	}
	for _, v := range r.groups {
		if v.parent != nil {
			fmt.Println(v.prefix, v.parent.prefix)
		} else {
			fmt.Println(v.prefix, v.parent)
		}
	}
}
