package bencode

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

var (
	ErrTyp     = errors.New("wrong type")
	ErrResolve = errors.New("format err")
)

type BType uint8

const (
	BSTR BType = iota
	BINT
	BLIST
	BDICT
)

type BValue any

type BObject struct {
	type_ BType
	val_  BValue
}

func NewBObject(type_ BType, val_ BValue) *BObject {
	return &BObject{type_: type_, val_: val_}
}

func (o *BObject) Str() (string, error) {
	if o.type_ != BSTR {
		return "", ErrTyp
	}
	return o.val_.(string), nil
}

func (o *BObject) Int() (int, error) {
	if o.type_ != BINT {
		return 0, ErrTyp
	}
	return o.val_.(int), nil
}

func (o *BObject) List() ([]*BObject, error) {
	if o.type_ != BLIST {
		return nil, ErrTyp
	}
	return o.val_.([]*BObject), nil
}

func (o *BObject) Dict() (map[string]*BObject, error) {
	if o.type_ != BDICT {
		return nil, ErrTyp
	}
	return o.val_.(map[string]*BObject), nil
}

func (o *BObject) BEncode(w io.Writer) int {
	bw, ok := w.(*bufio.Writer)
	if !ok {
		bw = bufio.NewWriter(w)
	}
	wLen := 0
	switch o.type_ {
	case BSTR:
		str, _ := o.Str()
		wLen += EncodeString(bw, str)
	case BINT:
		val, _ := o.Int()
		wLen += EncodeInt(bw, val)
	case BLIST:
		bw.WriteByte('l')
		list, _ := o.List()
		for _, elem := range list {
			wLen += elem.BEncode(bw)
		}
		bw.WriteByte('e')
		wLen += 2
	case BDICT:
		bw.WriteByte('d')
		dict, _ := o.Dict()
		for k, v := range dict {
			wLen += EncodeString(bw, k)
			wLen += v.BEncode(bw)
		}
		bw.WriteByte('e')
		wLen += 2
	}
	bw.Flush()
	return wLen
}

// EncodeString 将 go的string编码为BObject对象
func EncodeString(w io.Writer, val string) int {
	strLen := len(val)
	var bufStr string
	bufStr += strconv.Itoa(strLen)
	bufStr += ":"
	bufStr += val
	cnt, err := w.Write([]byte(bufStr))
	if err != nil {
		panic(err)
	}
	return cnt
}

func DecodeString(r io.Reader) (string, error) {
	br := bufio.NewReader(r)
	strLen, err := br.ReadString(':')
	if err != nil {
		if err == io.EOF {
			if strings.Contains(strLen, ":") {
				return "", nil
			}
			return "", ErrResolve
		}
		return "", err
	}
	strLen = strLen[:len(strLen)-1]
	length, err := strconv.Atoi(strLen)
	if err != nil {
		return "", err
	}
	buf := make([]byte, length)
	_, err = io.ReadFull(br, buf)
	return string(buf), nil
}

// EncodeInt ...
func EncodeInt(w io.Writer, val int) int {
	bufStr := ""
	bufStr += "i"
	bufStr += strconv.Itoa(val)
	bufStr += "e"
	cnt, err := w.Write([]byte(bufStr))
	if err != nil {
		return 0
	}
	return cnt
}

func DecodeInt(r io.Reader) (int, error) {
	br := bufio.NewReader(r)
	b, err := br.ReadByte()
	if b != 'i' {
		return 0, err
	}
	intStr, err := br.ReadString('e')
	if err != nil {
		return 0, err
	}
	intStr = intStr[:len(intStr)-1]
	val, err := strconv.Atoi(intStr)
	if err != nil {
		return 0, err
	}
	return val, nil
}
