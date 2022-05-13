/*
 * @Author: zwngkey
 * @Date: 2022-05-13 21:51:12
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-13 22:01:55
 * @Description:
 */
package gopkgcryptorand

import (
	"crypto/rand"
	"fmt"
	"testing"
)

/*
	rand.Read函数:	func rand.Read(b []byte) (n int, err error)
		生成随机数写入b.
*/
func Test(t *testing.T) {
	values := make([]byte, 16)
	rand.Read(values)
	fmt.Println(values)
}
