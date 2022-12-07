/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-07 23:18:09
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-08 01:00:44
 */
package prototype

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

// 深度复制可以基于reflect包的反射机制完成, 但是全部重头手写的话会很繁琐.
// 最简单的方式是基于序列化和反序列化来实现对象的深度复制:
func DeepCopyByGob(dst, src any) error {
	var buf bytes.Buffer
	// Gob的底层也是基于reflect包完成的.
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

// 在内存中序列化,反序列化对象实体 来完成对象实体的深拷贝这个是大部分语言实现对象深拷贝的惯用法。
// Go语言中所有赋值操作都是值传递，如果结构中不含指针，则直接赋值就是深度拷贝；
// 如果结构中含有指针（包括自定义指针，以及切片，map等使用了指针的内置类型），
// 则数据源和拷贝之间对应指针会共同指向同一块内存，这时深度拷贝需要特别处理。
// 目前，有三种方法，一是用gob序列化成字节序列再反序列化生成克隆对象；
// 二是先转换成json字节序列，再解析字节序列生成克隆对象；
// 三是针对具体情况，定制化拷贝。
// 前两种方法虽然比较通用但是因为使用了reflex反射，性能比定制化拷贝要低出2个数量级，所以在性能要求较高的情况下应该尽量避免使用前两者。

// 案例

// BidRequest 请求信息
type BidRequest struct {
	ID     string  `json:"id"`
	Imps   []*Imp  `json:"imp"`
	Device *Device `json:"device"`
}

// Imp: imp对象
type Imp struct {
	ID       string  `json:"id"`
	Tagid    string  `json:"tagid"`
	Bidfloor float64 `json:"bidfloor"`
}

// Device: 设备信息
type Device struct {
	Ua    string `json:"ua"`
	IP    string `json:"ip"`
	Geo   *Geo   `json:"geo"`
	Make  string `json:"make"`
	Model string `json:"model"`
	Os    string `json:"os"`
	Osv   string `json:"osv"`
}

// Geo: 地理位置信息
type Geo struct {
	Lat     int    `json:"lat"`
	Lon     int    `json:"lon"`
	Country string `json:"country"`
	Region  string `json:"region"`
	City    string `json:"city"`
}

// DeepCopyByJson 利用json进行深拷贝
func DeepCopyByJson(dst, src any) error {
	if tmp, err := json.Marshal(src); err != nil {
		return err
	} else {
		err = json.Unmarshal(tmp, dst)
		return err
	}
}

// 通过自定义进行深拷贝
func DeepCopyByCustom(dst, src *BidRequest) {
	dst.ID = src.ID
	dst.Device = &Device{
		Ua: src.Device.Ua,
		IP: src.Device.IP,
		Geo: &Geo{
			Lat:     src.Device.Geo.Lat,
			Lon:     src.Device.Geo.Lon,
			Country: src.Device.Geo.Country,
			Region:  src.Device.Geo.Region,
			City:    src.Device.Geo.City,
		},
		Make:  src.Device.Make,
		Model: src.Device.Model,
		Os:    src.Device.Os,
		Osv:   src.Device.Osv,
	}
	dst.Imps = make([]*Imp, len(src.Imps))
	for index, imp := range src.Imps {
		dst.Imps[index] = &Imp{
			imp.ID,
			imp.Tagid,
			imp.Bidfloor,
		}
	}
}

// CreateBidRequest 生成测试数据
func CreateBidRequest() *BidRequest {
	str := `{"id":"MM7dIXz4H05qtmViqnY5dW","imp":[{"id":"1","tagid":"3979722720","bidfloor":0.01}],"device":{"ua":"Mozilla/5.0 (Linux; Android 10; SM-G960N Build/QP1A.190711.020; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/92.0.4515.115 Mobile Safari/537.36 (Mobile; afma-sdk-a-v212621039.212621039.0)","ip":"192.168.1.0","geo":{"lat":0,"lon":0,"country":"KOR","region":"KR-11","city":"Seoul"},"make":"samsung","model":"sm-g960n","os":"android","osv":"10"}}`
	ans := new(BidRequest)
	json.Unmarshal([]byte(str), &ans)
	return ans
}

// BenchmarkDeepCopy_Gob 压测深拷贝 -gob
func BenchmarkDeepCopy_Gob(b *testing.B) {
	src := CreateBidRequest()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepCopyByGob(new(BidRequest), src)
	}
}

// BenchmarkDeepCopy_Json 压测深拷贝 -json
func BenchmarkDeepCopy_Json(b *testing.B) {
	src := CreateBidRequest()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepCopyByJson(new(BidRequest), src)
	}
}

// 压测深拷贝 -custom
func BenchmarkDeepCopy_custom(b *testing.B) {
	src := CreateBidRequest()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepCopyByCustom(new(BidRequest), src)
	}
}

// goos: darwin
// goarch: arm64
// pkg: zwngkey.cn/designpattern/creationalpattern/prototype
// BenchmarkDeepCopy_Gob-8            47876             22303 ns/op           16294 B/op        394 allocs/op
// BenchmarkDeepCopy_Json-8          252796              4668 ns/op            1488 B/op         27 allocs/op
// BenchmarkDeepCopy_custom-8      12458702                95.57 ns/op          232 B/op          4 allocs/op

func TestCopyIsOk(t *testing.T) {
	src := CreateBidRequest()
	//1.gob
	copy1 := new(BidRequest)
	DeepCopyByGob(copy1, src)
	fmt.Println(src)
	fmt.Printf("%p\n", src.Device)
	fmt.Printf("%p\n", src.Device.Geo)
	fmt.Printf("%p %p\n", src.Imps, src.Imps[0])
	fmt.Println(copy1)
	fmt.Printf("%p\n", copy1.Device)
	fmt.Printf("%p\n", copy1.Device.Geo)
	fmt.Printf("%p %p\n", copy1.Imps, copy1.Imps[0])
	fmt.Println(reflect.DeepEqual(src, copy1)) //true
}

// 测试拷贝是否ok
func TestCopyIsOk1(t *testing.T) {
	src := CreateBidRequest()

	//2.json
	copy1 := new(BidRequest)
	DeepCopyByJson(copy1, src)
	fmt.Println(src)
	fmt.Printf("%p\n", src.Device)
	fmt.Printf("%p\n", src.Device.Geo)
	fmt.Printf("%p %p\n", src.Imps, src.Imps[0])
	fmt.Println(copy1)
	fmt.Printf("%p\n", copy1.Device)
	fmt.Printf("%p\n", copy1.Device.Geo)
	fmt.Printf("%p %p\n", copy1.Imps, copy1.Imps[0])
	fmt.Println(reflect.DeepEqual(src, copy1)) //true

}

func TestCopyIsOk2(t *testing.T) {
	src := CreateBidRequest()

	//3.custom
	copy1 := new(BidRequest)
	DeepCopyByCustom(copy1, src)
	fmt.Println(src)
	fmt.Printf("%p\n", src.Device)
	fmt.Printf("%p\n", src.Device.Geo)
	fmt.Printf("%p %p\n", src.Imps, src.Imps[0])
	fmt.Println(copy1)
	fmt.Printf("%p\n", copy1.Device)
	fmt.Printf("%p\n", copy1.Device.Geo)
	fmt.Printf("%p %p\n", copy1.Imps, copy1.Imps[0])
	fmt.Println(reflect.DeepEqual(src, copy1)) //true
}
