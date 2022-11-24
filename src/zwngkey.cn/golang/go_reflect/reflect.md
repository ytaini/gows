[TOC]

# reflect api
reflect包实现了run-time(运行时)反射,允许程序操作任意类型的对象。典型的用法是使用静态类型interface{}获取一个值，然后调用`TypeOf`提取其动态类型信息，该方法返回一个`Type`类型。

调用`ValueOf`会 `returns a Value representing the run-time data`,`(Value)Zero`接受一个类型，并返回表示该类型的零值。

## type Kind

Kind代表Type类型值表示的具体分类。零值表示非法分类。

```go
type Kind uint

const (
    Invalid Kind = iota
    Bool
    Int
    Int8
    Int16
    Int32
    Int64
    Uint
    Uint8
    Uint16
    Uint32
    Uint64
    Uintptr
    Float32
    Float64
    Complex64
    Complex128
    Array
    Chan
    Func
    Interface
    Map
    Ptr
    Slice
    String
    Struct
    UnsafePointer
)
```

<br>
<br>

## type Type interface



```
> Type 是 Go类型的表示形式.
> 并非所有方法都适用于所有类型。
> 如果有限制会在每个方法的文档中都有说明。
> 在调用特定种类的方法之前,使用Kind()找出类型的种类.
> 调用不适合该种类的方法会导致运行时崩溃。
> Type值具有可比性，例如使用 == 运算符， 因此，它们可以用作映射键。
> 如果两个 Type 值表示相同的类型，则它们相等。
```

```go
func ArrayOf(length int, elem Type) Type
func ChanOf(dir ChanDir, t Type) Type
func FuncOf(in, out []Type, variadic bool) Type
func MapOf(key, elem Type) Type
func PointerTo(t Type) Type
func PtrTo(t Type) Type
func SliceOf(t Type) Type
func StructOf(fields []StructField) Type
func TypeOf(i any) Type
```
```go
type Type interface {
    // 适用于所有类型的方法。

    // Align 返回在内存中分配此类型值时,会对齐的字节数. 返回 rtype.align
	Align() int

    // FieldAlign 返回当该类型作为结构体的字段时，会对齐的字节数. 返回 rtype.fieldAlign
	FieldAlign() int

    // 返回该类型的方法集中的第i个方法，i不在[0, NumMethod())范围内时，将导致panic.
    // 对非接口类型T或*T，返回值的Type字段和Func字段描述了一个函数，该函数的第一个参数是接收者(receiver),并且只能访问导出的方法。
    // 对接口类型，返回值的Type字段描述方法的签名，Func字段为nil
    // Methods 按字典顺序排序。
	Method(int) Method

    // 根据方法名返回该类型的方法集中的方法，以及一个指示是否找到该方法的布尔值。
    // 对于非接口类型T或*T，返回值的Type字段和Func字段描述了一个函数，该函数的第一个参数是接收者(receiver)。
    // 对接口类型，返回值的Type字段是方法的签名，没有接收者(receiver),Func字段为nil
    // 内含隐藏或非导出方法
	MethodByName(string) (Method, bool)

    // 返回该类型的方法集中导出方法的数目
    // 匿名字段的方法会被计算；主体类型的方法会屏蔽匿名字段的同名方法；
    // 匿名字段导致的歧义方法会滤除
    // 对于非接口类型，它返回导出方法的数量。
    // 对于接口类型，它返回导出和未导出方法的数量。
	NumMethod() int

    // 返回已定义类型在包中的类型名称。 
    // 对于未定义(未命名)的类型，它返回空字符串。
	Name() string

	// PkgPath返回类型的包路径，即明确指定包的import路径，如"encoding/base64"
    // 如果类型为内建类型(string, error)或未命名类型(*T, struct{}, []int)，会返回""
	PkgPath() string

	 // 返回要保存一个该类型的值需要多少字节；类似unsafe.Sizeof
	Size() uintptr

    // 返回该类型的字符串表示。该字符串可能会使用短包名（如用base64代替"encoding/base64"）
    // 也不保证每个类型的字符串表示不同。如果要比较两个类型是否相等，请直接用Type类型比较。
	String() string

    // Kind返回该类型的特定种类。
	Kind() Kind

	// 如果该类型实现了u代表的接口，会返回真
	Implements(u Type) bool

    // 如果该类型的值可以直接赋值给u代表的类型，返回真
	AssignableTo(u Type) bool

    // 如该类型的值可以转换为u代表的类型，返回真
    // 即使ConvertibleTo返回true，转换仍然可能出错。
    // 例如，[]T可转换为*[N]T， 但如果长度小于N，转换就会出错。
    //  arr := [...]int{1, 2, 2}
    //  s := arr[:]
    //  arr1 := (*[4]int)(s)
    //  fmt.Println(arr1)
	ConvertibleTo(u Type) bool

    // 该类型的值可比较。返回true
    // 即使Comparable返回true，比较仍然可能出错
    // 例如，接口类型的值是可比的,但是如果它们的动态类型不可比较，比较就会出错。，
	Comparable() bool




    // 下面的方法只适用于某些类型，具体取决于Kind。
    // 每种Kind允许的方法如下:
    //	Int*, Uint*, Float*, Complex*: Bits
	//	Array: Elem, Len
	//	Chan: ChanDir, Elem
	//	Func: In, NumIn, Out, NumOut, IsVariadic.
	//	Map: Key, Elem
	//	Pointer: Elem
	//	Slice: Elem
	//	Struct: Field, FieldByIndex, FieldByName, FieldByNameFunc, NumField

    // Bits以位数为单位返回该类型的大小。
    // 如果该类型的Kind不是Int、Uint、Float或Complex，会panic
	Bits() int

    // 返回一个channel类型的方向，
    // 如果该类型的Kind不是Chan类型将会panic
	ChanDir() ChanDir

	
    // 如果函数类型的最后一个输入参数是"..."形式的参数，IsVariadic返回真
    // 如果是，t.In(t.NumIn() - 1)返回参数的隐式的实际类型[]T（声明类型的切片）
    // 具体来说，如果t表示func(x int, y…float64),然后
	//	t.NumIn() == 2
	//	t.In(0) is the reflect.Type for "int"
	//	t.In(1) is the reflect.Type for "[]float64"
	//	t.IsVariadic() == true
    
    // 如果该类型的Kind不是Func类型将panic
	IsVariadic() bool

    // 返回该类型的元素类型
    // 如果该类型的Kind不是Array、Chan、Map、Ptr或Slice，会panic
	Elem() Type

	// 返回struct类型的第i个字段的类型，
    // 如果该类型的Kind不是Struct或者i不在[0, NumField())内将会panic
	Field(i int) StructField


    // type Person struct {
    //     Addr
    // }
    // type Addr struct {
    //     city string
    // }
    // sf1 := reflect.TypeOf(p).FieldByIndex([]int{0, 0})
	// fmt.Println(sf1.Name) //city
	// fmt.Println(sf1.Type) //string
    // 返回与index sequence对应的嵌套字段.
    // 相当于对每个索引i依次调用Field。
    // 如果该类型的Kind不是结构体将会panic
	FieldByIndex(index []int) StructField

    // FieldByName返回具有给定名称的struct字段,以及一个布尔值，表示是否找到了该字段。
    // 如果该类型的Kind不是结构体将会panic
	FieldByName(name string) (StructField, bool)

    // 返回该类型第一个字段名满足函数match的字段，布尔值说明是否找到
    // 如果该类型的Kind不是结构体将会panic
	FieldByNameFunc(match func(string) bool) (StructField, bool)

    // 返回func类型的第i个参数的类型
	// It panics if the type's Kind is not Func.
	// It panics if i is not in the range [0, NumIn()).
	In(i int) Type

    // 返回map类型的键的类型 
	// It panics if the type's Kind is not Map.
	Key() Type

	// 返回array类型的长度
	// It panics if the type's Kind is not Array.
	Len() int

	// 返回struct类型的字段数（匿名字段算作一个字段）
	// It panics if the type's Kind is not Struct.
	NumField() int

	// 返回func类型的参数个数
	// It panics if the type's Kind is not Func.
	NumIn() int

	// 返回func类型的返回值个数.
	// It panics if the type's Kind is not Func.
	NumOut() int

	// 返回func类型的第i个返回值的类型
	// It panics if the type's Kind is not Func.
	// It panics if i is not in the range [0, NumOut()).
	Out(i int) Type

	common() *rtype
	uncommon() *uncommonType
}
```

<br>
<br>

## type Value struct 

反射包中 Value 的类型与 Type 不同，它被声明成了结构体。这个结构体没有对外暴露的字段，但是提供了获取或者写入数据的方法：
```go
> Value 是 Go value 的反射接口
> 并非所有方法都适用于所有类型。
> 如果有限制会在每个方法的文档中都有说明。
> 在调用特定种类的方法之前,使用Kind()找出类型的种类.
> 调用不适合该种类的方法会导致运行时崩溃。


> The zero Value represents no value.
> 其 IsValid方法返回 false, 其 Kind方法返回 Invalid,
> 其 String方法返回"<invalid Value>", 其他所有方法引发panic.

> 大多数函数和方法永远不会返回 an invalid value。
> 如果有，其文档会明确说明条件。

> 一个Value可以被多个goroutines并发使用,前提是底层 Go 值可以并发用于等效的直接操作。

> 要比较两个Value，请比较Interface方法的结果。
> 对两个Value使用==不会比较它们所表示的底层值。
```

```go
func Append(s Value, x ...Value) Value
func AppendSlice(s, t Value) Value
func Indirect(v Value) Value
func MakeChan(typ Type, buffer int) Value
func MakeFunc(typ Type, fn func(args []Value) (results []Value)) Value
func MakeMap(typ Type) Value
func MakeMapWithSize(typ Type, n int) Value
func MakeSlice(typ Type, len, cap int) Value
func New(typ Type) Value
func NewAt(typ Type, p unsafe.Pointer) Value
func Select(cases []SelectCase) (chosen int, recv Value, recvOK bool)
func ValueOf(i any) Value
func Zero(typ Type) Value
```

```go
type Value struct {
    // typ 保存由 Value 表示的值的类型。
    // typ 是未导出的，从外部调不到 Type 接口的方法
    typ *rtype

    // 数据形式的指针值
    ptr unsafe.Pointer

    // 保存元数据
    flag
}
```

Value 的方法
``` go
// Addr返回一个指针值，表示v的地址。如果CanAddr()返回false，就会发生错误。
// Addr通常用于获取指向struct字段或slice元素的指针，以便调用需要指针接收器的方法。
func (v Value) Addr() Value

// Bool返回v的底层值.
// 如果v的Kind不是Bool类型,它会panic.
func (v Value) Bool() bool

// Bytes返回v的底层值
// 如果v的底层值不是一个字节切片或一个可寻址的字节数组，它会panic.
func (v Value) Bytes() []byte

// Call调用函数v并传入参数。
// 例如，如果 len（in） == 3，v.Call（in） 表示 Go 调用 v（in[0]， in[1]， in[2]）
// 如果v的Kind不是Func类型,它会panic.
// 它将返回值包装成[]Value返回.
// 在GO中，每个输入参数都必须可分配给函数的相应输入参数的类型。
// 如果 v 是可变参数函数，则 Call 会自行创建可变参数切片参数，并复制相应的值。
func (v Value) Call(in []Value) []Value


// CallSlice 调用带有可变参数的函数 v，将切片 in[len（in）-1] 分配给 v 的可变参数。 
// 例如，如果 len（in） == 3，v.CallSlice（in） 表示 Go 调用 v（in[0]， in[1]， in[2]...）。 
// 如果 v 的 Kind 不是 Func 或者 v 没有带可变参数，CallSlice 会panic。 
// 它将返回值包装成[]Value返回。 
// 在GO中，每个输入参数都必须可分配给函数的相应输入参数的类型。
func (v Value) CallSlice(in []Value) []Value


// 如果可以使用Addr()获取v的地址.返回true
// 此类值v称为可寻址值。如果v是切片的元素、可寻址数组的元素、可寻址结构的字段或取消引用指针的结果，则v是可寻址的。
// 如果 CanAddr 返回 false，则调用 Addr() 将出现恐慌。
func (v Value) CanAddr() bool


// 如果可以使用Complex(),返回true.
func (v Value) CanComplex() bool


// v是否可以转换为类型t。
// 如果v.CanConvert(t)返回true，那么v.convert(t)就不会发生panic.
func (v Value) CanConvert(t Type) bool


// v是否可以使用Float().
func (v Value) CanFloat() bool


// v是否可以使用Int().
func (v Value) CanInt() bool


// v是否可以使用Interface().
func (v Value) CanInterface() bool


// v 的值是否可以修改
// 只有当值v是可寻址的，并且不是通过使用未导出的struct字段获得时，才能更改该值。
// 如果CanSet返回false，调用Set或任何特定类型的setter(例如SetBool, SetInt)将出现panic。
func (v Value) CanSet() bool


// v是否可以使用Uint().
func (v Value) CanUint() bool

//  Cap returns v's capacity.
//  It panics if v's Kind is not Array, Chan, Slice or 指向数组的指针
func (v Value) Cap() int


// Close closes the channel v. 
// It panics if v's Kind is not Chan.
func (v Value) Close()



// Complex returns v's underlying value, as a complex128. 
// It panics if v's Kind is not Complex64 or Complex128
func (v Value) Complex() complex128


// 返回转换为类型 t 的值 v。
// 如果通常的 Go 转换规则不允许将值 v 转换为类型 t，或者如果将 v 转换为类型 t 发生恐慌，则 Convert 恐慌。
func (v Value) Convert(t Type) Value


// Elem返回接口v包含的值或指针v指向的值。
// It panics if v's Kind is not Interface or Pointer.
// It returns the zero Value if v is nil.
func (v Value) Elem() Value


// Field返回结构体v的第i个字段。
// It panics if v's Kind is not Struct or i is out of range. 
func (v Value) Field(i int) Value


// FieldByIndex返回与index对应的嵌套字段
// It panics if evaluation requires stepping through a nil pointer or a field that is not a struct.
func (v Value) FieldByIndex(index []int) Value
func (v Value) FieldByIndexErr(index []int) (Value, error)


// FieldByName返回给定名称的struct字段
// It returns the zero Value if no field was found. 
// It panics if v's Kind is not struct.
func (v Value) FieldByName(name string) Value

// 返回满足match函数的名称的struct字段
// It panics if v's Kind is not struct. 
// It returns the zero Value if no field was found.
func (v Value) FieldByNameFunc(match func(string) bool) Value


// Float returns v's underlying value, as a float64. 
// It panics if v's Kind is not Float32 or Float64
func (v Value) Float() float64


// Int returns v's underlying value, as an int64. 
// It panics if v's Kind is not Int, Int8, Int16, Int32, or Int64.
func (v Value) Int() int64


// Index returns v's i'th element. 
// It panics if v's Kind is not Array, Slice, or String or i is out of range.
func (v Value) Index(i int) Value


// Interface returns v's current value as an interface{}. 
// It is equivalent to: var i interface{} = (v's underlying value)
// It panics if the Value was obtained by accessing unexported struct fields.
func (v Value) Interface() (i any)


// 判断v 是否 等于 nil
// v must be a chan, func, interface, map, pointer, or slice value;  if it is not, IsNil panics. 
// 请注意，IsNil 并不总是等同于 Go 中与 nil 的常规比较。  
// 例如，如果 v 是通过使用未初始化的接口变量 i 调用 ValueOf 创建的，则 i==nil 将为真，但 v.IsNil 将崩溃，因为 v 将是零值。
func (v Value) IsNil() bool



// IsValid reports whether v represents a value. 
// It returns false if v is the zero Value.
// If IsValid returns false, 则除String() 之外的所有其他方法会panic。
// Most functions and methods never return an invalid Value. 
// If one does, its documentation states the conditions explicitly.
func (v Value) IsValid() bool


// IsZero reports whether v 是否为其类型的零值
// It panics if the argument(v) is invalid.
func (v Value) IsZero() bool


// Kind returns v's Kind. 
// If v is the zero Value (IsValid() returns false), Kind returns Invalid.
func (v Value) Kind() Kind


// Len returns v's length. 
// It panics if v's Kind is not Array, Chan, Map, Slice, String, or 指向数组的指针
func (v Value) Len() int


// MapIndex returns the value associated with key in the map v. 
// It panics if v's Kind is not Map. 
// It returns the zero Value if key is not found in the map or if v represents a nil map. 
// As in Go, the key's value must be assignable to the map's key type.
func (v Value) MapIndex(key Value) Value


// MapKeys returns a slice containing all the keys present in the map, 但没有指定顺序。
// It panics if v's Kind is not Map. 
// It returns an empty slice if v represents a nil map.
func (v Value) MapKeys() []Value


// MapRange returns a range iterator for a map. 
// It panics if v's Kind is not Map.
// Call Next() to advance(推进) the iterator;
// Call Key()/Value() to access each entry.
// Next returns false when the iterator is exhausted(耗尽)
// MapRange与range语句遵循相同的迭代语义。
// Example:
//  iter := reflect.ValueOf(m).MapRange()
//  for iter.Next() {
// 	    k := iter.Key()
// 	    v := iter.Value()
// 	    ...
//  }
func (v Value) MapRange() *MapIter


// Method returns a function value corresponding to v's i'th method. 
// The arguments to a Call on the returned function should not include a receiver; 
// the returned function will always use v as the receiver. 
// Method panics if i is out of range or if v is a nil interface value.
func (v Value) Method(i int) Value


// MethodByName 返回 一个 function Value, 对应于具有给定名称的 v 的方法
// 注意: 对返回函数的调用的参数不应包含接收者(receiver);
// the returned function will always use v as the receiver.  
// It returns the zero Value if no method was found.
func (v Value) MethodByName(name string) Value


// NumField returns the number of fields in the struct v. 
// It panics if v's Kind is not Struct.
func (v Value) NumField() int


// NumMethod 返回v的方法集的数目。
// 对于非接口类型，它返回导出方法的数量。
// 对于接口类型，它返回导出和未导出方法的数量。
func (v Value) NumMethod() int


// OverflowFloat reports float64 x 是否不能用 v 的类型表示。
// It panics if v's Kind is not Float32 or Float64.
func (v Value) OverflowFloat(x float64) bool


// It panics if v's Kind is not Complex64 or Complex128.
func (v Value) OverflowComplex(x complex128) bool


// OverflowInt reports int64 x 是否不能用 v 的类型表示。 
// It panics if v's Kind is not Int, Int8, Int16, Int32, or Int64.
func (v Value) OverflowInt(x int64) bool


// OverflowUint reports uint64 x 是否不能用 v 的类型表示。
// It panics if v's Kind is not Uint, Uintptr, Uint8, Uint16, Uint32, or Uint64.
func (v Value) OverflowUint(x uint64) bool


func (v Value) Pointer() uintptr


// Recv receives and returns a value from the channel v. 
// It panics if v's Kind is not Chan. 
// The receive blocks until a value is ready. 
// The boolean value ok is true if the value x corresponds to a send on the channel, false if it is a zero value received because the channel is closed.
func (v Value) Recv() (x Value, ok bool)


// Send sends x on the channel v. 
// It panics if v's kind is not Chan or if x的类型与v的元素类型不同。
// As in Go, x's value must be assignable to the channel's element type.
func (v Value) Send(x Value)


// Set assigns x to the value v. 
// It panics if CanSet returns false. 
// As in Go, x's value must be assignable to v's type.
func (v Value) Set(x Value)


// SetBool sets v's underlying value. 
// It panics if v's Kind is not Bool or if CanSet() is false.
func (v Value) SetBool(x bool)


// SetBytes sets v's underlying value. 
// It panics if v's underlying value is not a slice of bytes.
func (v Value) SetBytes(x []byte)


// SetCap sets v's capacity to n. 
// It panics if v's Kind is not Slice or if n 小于slice长度或大于slice容量。
func (v Value) SetCap(n int)


// SetComplex sets v's underlying value to x. 
// It panics if v's Kind is not Complex64 or Complex128, or if CanSet() is false.
func (v Value) SetComplex(x complex128)


// SetFloat sets v's underlying value to x. 
// It panics if v's Kind is not Float32 or Float64, or if CanSet() is false.
func (v Value) SetFloat(x float64)


// SetInt sets v's underlying value to x. 
// It panics if v's Kind is not Int, Int8, Int16, Int32, or Int64, or if CanSet() is false.
func (v Value) SetInt(x int64)

func (v Value) SetIterKey(iter *MapIter)

func (v Value) SetIterValue(iter *MapIter)


// SetLen sets v's length to n. 
// It panics if v's Kind is not Slice or if n 为负数或大于切片的容量。
func (v Value) SetLen(n int)


// SetMapIndex将map v中与key相关联的元素设置为elem。
// It panics if v's Kind is not Map.
// if v holds(保存) a nil map, SetMapIndex will panic. 
// If elem is the zero Value, SetMapIndex deletes the key from the map. 
// As in Go, key's elem must be assignable to the map's key type, and elem's value must be assignable to the map's elem type.
func (v Value) SetMapIndex(key, elem Value)


// SetString sets v's underlying value to x. 
// It panics if v's Kind is not String or if CanSet() is false.
func (v Value) SetString(x string)


// SetUint sets v's underlying value to x. 
// It panics if v's Kind is not Uint, Uintptr, Uint8, Uint16, Uint32, or Uint64, or if CanSet() is false.
func (v Value) SetUint(x uint64)


// Slice returns v[i:j]. 
// It panics if v's Kind is not Array, Slice or String, or if v is an unaddressable array, or if the indexes are out of bounds.
func (v Value) Slice(i, j int) Value


// Slice3 is the 3-index form of the slice operation: it returns v[i:j:k].
// 如果 v 的Kind不是数组或切片，或者 v 是不可寻址的数组，或者索引超出界限，它会panic。
func (v Value) Slice3(i, j, k int) Value


// String returns the string v's underlying value, as a string
// String is a special case because of Go's String method convention.  
// Unlike the other getters, it does not panic if v's Kind is not String. 
// 相反，它会返回一个形式为"<T value>"的字符串，其中T是v的类型。
// fmt 包会特别处理Value值.它不会隐式调用它们的 String 方法，而是打印它们持有的具体值。
func (v Value) String() string

// TryRecv 尝试从通道 v 接收值，但不会阻塞。
// 如果接受到值,则x为接受到的值,ok为true.
// 如果接收无法在不阻塞的情况下完成，则 x 为零值，ok 为假。
// 如果通道已关闭，则 x 是通道元素类型的零值，ok 为假。
// It panics if v's Kind is not Chan.
func (v Value) TryRecv() (x Value, ok bool)


// TrySend 尝试在通道 v 上发送 x，但不会阻塞。 它报告是否已发送值
// It panics if v's Kind is not Chan.
// As in Go, x's value must be assignable to the channel's element type.
func (v Value) TrySend(x Value) bool


// Type returns v's type.
func (v Value) Type() Type


// Uint returns v's underlying value, as a uint64. 
// It panics if v's Kind is not Uint, Uintptr, Uint8, Uint16, Uint32, or Uint64.
func (v Value) Uint() uint64    
```


<br>
<br>


## type StructTag string

```go
func (tag StructTag) Get(key string) string


func (tag StructTag) Lookup(key string) (value string, ok bool)
```

<br>
<br>

## type Method struct

```go
type Method struct {
	// Name is the method name.
	Name string

	// PkgPath is the package path that qualifies a lower case (unexported)
	// method name. It is empty for upper case (exported) method names.
	// The combination of PkgPath and Name uniquely identifies a method
	// in a method set.
	// See https://golang.org/ref/spec#Uniqueness_of_identifiers
	PkgPath string

	Type  Type  // method type
	Func  Value // func with receiver as first argument
	Index int   // index for Type.Method
}
```


```go
// IsExported reports whether the method is exported.
func (m Method) IsExported() bool
```

<br>
<br>


## type StructField struct
```go
type StructField struct {
	// Name is the field name.
	Name string

	// PkgPath is the package path that qualifies a lower case (unexported)
	// field name. It is empty for upper case (exported) field names.
	// See https://golang.org/ref/spec#Uniqueness_of_identifiers
	PkgPath string

	Type      Type      // field type
	Tag       StructTag // field tag string
	Offset    uintptr   // offset within struct, in bytes
	Index     []int     // index sequence for Type.FieldByIndex
	Anonymous bool      // is an embedded field
}

// VisibleFields 返回 t 中的所有可见字段，该字段必须是struct类型。
// 如果可以通过FieldByName调用直接访问字段，则该字段被定义为可见的.
// 返回的字段包括匿名结构成员中的字段和未导出的字段。
// 它们的顺序与结构体中的顺序相同，匿名字段紧接着的是提升字段。
// 对于返回切片的每个元素e，可以通过调用v.Fieldbyindex(e.Index)从类型为t的值v中取得对应的字段。
func VisibleFields(t Type) []StructField

// IsExported reports whether the field is exported.
func (f StructField) IsExported() bool
```



## type MapIter struct

```go
// Key returns the key of iter's current map entry.
func (iter *MapIter) Key() Value


// Value returns the value of iter's current map entry.
func (iter *MapIter) Value() Value


// Next advances the map iterator and reports whether there is another entry. 
// It returns false when iter is exhausted; 
// subsequent(后续) calls to Key, Value, or Next will panic.
func (iter *MapIter) Next() bool


// Reset modifies iter to iterate over v. 
// It panics if v's Kind is not Map and v is not the zero Value. 
// Reset(Value{}) causes iter to not to refer to any map, which may allow the previously iterated-over map to be garbage collected.
func (iter *MapIter) Reset(v Value)
```
