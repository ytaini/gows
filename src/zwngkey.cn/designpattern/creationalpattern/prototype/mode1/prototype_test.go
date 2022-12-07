/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-08 00:43:51
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-08 01:00:25
 */
package mode1

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

// initData 生成测试数据
func CreateBidRequest() *BidRequest {
	str := `{"id":"MM7dIXz4H05qtmViqnY5dW","imp":[{"id":"1","tagid":"3979722720","bidfloor":0.01}],"device":{"ua":"Mozilla/5.0 (Linux; Android 10; SM-G960N Build/QP1A.190711.020; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/92.0.4515.115 Mobile Safari/537.36 (Mobile; afma-sdk-a-v212621039.212621039.0)","ip":"192.168.1.0","geo":{"lat":0,"lon":0,"country":"KOR","region":"KR-11","city":"Seoul"},"make":"samsung","model":"sm-g960n","os":"android","osv":"10"}}`
	ans := new(BidRequest)
	json.Unmarshal([]byte(str), &ans)
	return ans
}

func Test(t *testing.T) {
	src := CreateBidRequest()
	copy1 := src.Clone().(*BidRequest)

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

func Test2(t *testing.T) {
	src := CreateBidRequest()

	pm := NewPrototypeManager()
	pm.Set("br1", src)

	copy1 := pm.Get("br1").(*BidRequest)

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

type BidRequest struct {
	ID     string  `json:"id"`
	Imps   []*Imp  `json:"imp"`
	Device *Device `json:"device"`
}

func (b *BidRequest) Clone() Clonable {
	imps := make([]*Imp, len(b.Imps))
	for i := range imps {
		imps[i] = b.Imps[i].Clone().(*Imp)
	}
	return &BidRequest{
		b.ID,
		imps,
		b.Device.Clone().(*Device),
	}
}

// Imp: imp对象
type Imp struct {
	ID       string  `json:"id"`
	Tagid    string  `json:"tagid"`
	Bidfloor float64 `json:"bidfloor"`
}

func (i *Imp) Clone() Clonable {
	return &Imp{
		i.ID,
		i.Tagid,
		i.Bidfloor,
	}
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

func (d *Device) Clone() Clonable {
	return &Device{
		d.Ua,
		d.IP,
		d.Geo.Clone().(*Geo),
		d.Make,
		d.Model,
		d.Os,
		d.Osv,
	}
}

// Geo: 地理位置信息
type Geo struct {
	Lat     int    `json:"lat"`
	Lon     int    `json:"lon"`
	Country string `json:"country"`
	Region  string `json:"region"`
	City    string `json:"city"`
}

func (g *Geo) Clone() Clonable {
	return &Geo{
		g.Lat,
		g.Lat,
		g.Country,
		g.Region,
		g.City,
	}
}
