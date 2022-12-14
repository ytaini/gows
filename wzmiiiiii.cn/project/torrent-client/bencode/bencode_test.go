package bencode

import (
	"io"
	"os"
	"strconv"
	"testing"
)

func TestEncodeString(t *testing.T) {
	file, _ := os.OpenFile("./test1", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	data := "absflkjhsaflsdjfdafhafhahfaksljfhahf"
	cnt := EncodeString(file, data)
	lenData := strconv.Itoa(len(data))
	if cnt != len(data)+len(lenData)+1 {
		t.Errorf("Encode err!!")
	}
	_ = file.Close()
}

func TestDecodeString(t *testing.T) {
	file, _ := os.Open("./test1")
	str, err := DecodeString(file)
	if err != nil {
		if err != io.EOF {
			t.Errorf(err.Error())
		}
	}
	t.Log(str)
	_ = file.Close()
}

func TestEncodeInt(t *testing.T) {
	file, _ := os.OpenFile("./test1", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	data := 1248712
	cnt := EncodeInt(file, data)
	t.Log(cnt)
	_ = file.Close()
}

func TestDecodeInt(t *testing.T) {
	file, _ := os.Open("./test1")
	val, err := DecodeInt(file)
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Logf("%T,%d", val, val)
	_ = file.Close()
}

func TestBEncode(t *testing.T) {
	val1_ := make(map[string]*BObject)
	val1_["length"] = NewBObject(BINT, 351272960)
	val1_["name"] = NewBObject(BSTR, "debian-10.2.0-amd64-netinst.iso")
	val1_["piece length"] = NewBObject(BINT, 262144)
	val1_["pieces"] = NewBObject(BSTR, "askfhjghahkfahjfgahfhsa")

	val_ := make(map[string]*BObject)
	val_["announce"] = NewBObject(BSTR, "http://bttracker.debian.org:6969/announce")
	val_["info"] = NewBObject(BDICT, val1_)
	bo := NewBObject(BDICT, val_)

	file, _ := os.OpenFile("./test1", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	cnt := bo.BEncode(file)
	t.Log(cnt)
}
