/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-30 17:52:43
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-30 18:38:09
 * @Description:

 */
package problem

import (
	"fmt"
	"testing"
)

func Test10(t *testing.T) {
	// 116.61607692307692
	exp := "12.813 + (2 - 3.55)*4+10/5.2 + 8*13.51"
	res, err := CalcExp(exp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("res: ", res)
}
