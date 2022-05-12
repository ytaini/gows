/*
 * @Author: zwngkey
 * @Date: 2022-05-07 00:18:26
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-07 02:54:35
 * @Description: go 方法
 */
package gomethod

import (
	"fmt"
	"testing"
)

/*
	方法
		Go支持一些面向对象编程特性，方法是这些所支持的特性之一。

	方法声明
		在Go中，我们可以为类型T和*T显式地声明一个方法，其中类型T必须满足四个条件：
			T必须是一个定义类型；
			T必须和此方法声明定义在同一个代码包中；
			T不能是一个指针类型；
			T不能是一个接口类型。

		invalid receiver type Ptr (pointer or interface type)
		invalid receiver type *int (basic or unnamed type)

	类型T和*T称为它们各自的方法的属主类型（receiver type）。
		类型T被称作为类型T和*T声明的所有方法的属主基类型（receiver base type）

	注意：我们也可以为满足上列条件的类型T和*T的别名声明方法。 这样做的效果和直接为类型T和*T声明方法是一样的。

	如果我们为某个类型声明了一个方法，以后我们可以说此类型拥有此方法。

	从上面列出的条件，我们得知我们不能为下列类型（显式地）声明方法：
		内置基本类型。比如int和string。 因为这些类型声明在内置builtin标准包中，而我们不能在标准包中声明方法。
		接口类型。但是接口类型可以拥有方法。
		除了满足上面条件的形如*T的指针类型之外的非定义组合类型。

	一个方法声明和一个函数声明很相似，但是比函数声明多了一个额外的参数声明部分。
		 此额外的参数声明部分只能含有一个类型为此方法的属主类型的参数，此参数称为此方法声明的属主参数（receiver parameter）。
		  此属主参数声明必须包裹在一对小括号()之中。 此属主参数声明部分必须处于func关键字和方法名之间。

*/
// Age和int是两个不同的类型。我们不能为int和*int
// 类型声明方法，但是可以为Age和*Age类型声明方法。
type Age int

func (age Age) LargerThan(a Age) bool {
	return age > a
}
func (age *Age) Increase() {
	*age++ // 如果age是一个空指针，则此行将产生一个恐慌。
}

// 为自定义的函数类型FilterFunc声明方法。
type FilterFunc func(in int) bool

func (ff FilterFunc) Filte(in int) bool {
	return ff(in)
}

// 为自定义的映射类型StringSet声明方法。
type StringSet map[string]struct{}

func (ss StringSet) Has(key string) bool {
	_, present := ss[key] // 永不会产生恐慌，即使ss为nil。
	return present
}
func (ss StringSet) Add(key string) {
	ss[key] = struct{}{}
}
func (ss StringSet) Remove(key string) {
	delete(ss, key)
}

// 为自定义的结构体类型Book和它的指针类型*Book声明方法。

type Book struct {
	pages int
}

func (b Book) Pages() int {
	return b.pages
}

func (b *Book) SetPages(pages int) {
	b.pages = pages
}

/*
	从上面的例子可以看出，我们可以为各种种类（kind）的类型声明方法，而不仅仅是结构体类型。

	在很多其它面向对象的编程语言中，属主参数名总是为隐式声明的this或者self。这样的名称不推荐在Go编程中使用。

	指针类型的属主参数称为指针类型属主，非指针类型的属主参数称为值类型属主。

	方法名可以是空标识符_。一个类型可以拥有若干名为空标识符的方法，但是这些方法无法被调用。
		只有导出的方法才可以在其它代码包中调用。


	每个方法对应着一个隐式声明的函数
		对每个方法声明，编译器将自动隐式声明一个相对应的函数。
		  比如对于上一节的例子中为类型Book和*Book声明的两个方法，编译器将自动声明下面的两个函数：
			func Book.Pages(b Book) int {
				return b.pages // 此函数体和Book类型的Pages方法体一样
			}

			func (*Book).SetPages(b *Book, pages int) {
				b.pages = pages // 此函数体和*Book类型的SetPages方法体一样
			}
		在上面的两个隐式函数声明中，它们各自对应的方法声明的属主参数声明被插入到了普通参数声明的第一位。
			 它们的函数体和各自对应的显式方法的方法体是一样的。

		两个隐式函数名Book.Pages和(*Book).SetPages都是aType.MethodName这种形式的。
			 我们不能显式声明名称为这种形式的函数，因为这种形式不属于合法标识符。这样的函数只能由编译器隐式声明。
			  但是我们可以在代码中调用这些隐式声明的函数： 如下所示
*/
func TestEg11(t *testing.T) {
	var book Book
	// 调用隐式声明的函数
	Book.Pages(book)
	(*Book).SetPages(&book, 123)
}

/*
	事实上，在隐式声明上述两个函数的同时，编译器也将改写这两个函数对应的显式方法（至少，我们可以这样认为），
		让这两个方法在体内直接调用这两个隐式函数：
			func (b Book) Pages() int {
				return Book.Pages(b)
			}

			func (b *Book) SetPages(pages int) {
				(*Book).SetPages(b, pages)
			}
*/
/*
	为指针类型属主隐式声明的方法
		对每一个为值类型属主T声明的方法，一个相应的同名方法将自动隐式地为其对应的指针类型属主*T而声明
			以上面的为类型Book声明的Pages方法为例，一个同名方法将自动为类型*Book而声明：
			// 注意：这不是合法的Go语法。这里这样表示只是为了解释目的。
			// 它表明表达式(&aBook).Pages将被估值为aBook.Pages
			func (b *Book) Pages = (*b).Pages

		当我们为一个非指针类型显式声明一个方法的时候，事实上两个方法被声明了。
			一个方法是为非指针类型显式声明的，另一个是为指针类型隐式声明的。

		每一个方法对应着一个编译器隐式声明的函数。 所以对于刚提到的隐式方法，编译器也将隐式声明一个相应的函数：
			func (*Book).Pages(b *Book) int {
				return Book.Pages(*b)
			}

		换句话说，对于每一个为值类型属主显式声明的方法，同时将有一个隐式方法和两个隐式函数被自动声明。
*/
/*
	方法原型（method prototype）和方法集（method set）
		一个方法原型可以看作是一个不带func关键字的函数原型。
			我们可以把每个方法声明看作是由一个func关键字、一个属主参数声明部分、一个方法原型和一个方法体组成。

		例如:上面的例子中的Pages和SetPages的原型如下：
			Pages() int
			SetPages(pages int)

		每个类型都有个方法集。一个非接口类型的方法集由所有为它声明的（不管是显式的还是隐式的）方法的原型组成。

		比如，在上面的例子中，Book类型的方法集为：
			Pages() int
		而*Book类型的方法集为：
			Pages() int
			SetPages(pages int)

		对于一个方法集，如果其中的每个方法原型都处于另一个方法集中，则我们说前者方法集为后者（即另一个）方法集的子集，后者为前者的超集。
			如果两个方法集互为子集（或超集），则这两个方法集必等价。

		给定一个类型T，假设它既不是一个指针类型也不是一个接口类型，因为上一节中提到的原因，类型T的方法集总是类型*T的方法集的子集。
			 比如，在上面的例子中，Book类型的方法集为*Book类型的方法集的子集

		请注意：不同代码包中的同名非导出方法将总被认为是不同名的。

		方法集在Go中的多态特性中扮演着重要的角色。

		下列类型的方法集总为空：
			内置基本类型；
			定义的指针类型；
			基类型为指针类型或者接口类型的指针类型；
			非定义的数组/切片/映射/函数/通道类型。

*/
/*
	方法值和方法调用
		方法事实上是特殊的函数。方法也常被称为成员函数。
		 当一个类型拥有一个方法，则此类型的每个值将拥有一个不可修改的函数类型的成员（类似于结构体的字段）。
		  此成员的名称为此方法名，它的类型和此方法的声明中不包括属主部分的函数声明的类型一致。

		一个方法调用其实是调用了一个值的成员函数。假设一个值v有一个名为m的方法，则此方法可以用选择器语法形式v.m来表示。
*/
func TestEg12(t *testing.T) {
	var book Book

	fmt.Printf("%T \n", book.Pages)       // func() int
	fmt.Printf("%T \n", (&book).SetPages) // func(int)
	// &book值有一个隐式方法Pages。
	fmt.Printf("%T \n", (&book).Pages) // func() int

	// 调用这三个方法。
	(&book).SetPages(123)
	book.SetPages(123)           // 等价于上一行
	fmt.Println(book.Pages())    // 123
	fmt.Println((&book).Pages()) // 123
}

/*
	上例中的(&book).SetPages(123)一行为什么可以被简化为book.SetPages(123)呢？ 毕竟，类型Book并不拥有一个SetPages方法。
		这可以看作是Go中为了让代码看上去更简洁而特别设计的语法糖。此语法糖只对可寻址的值类型的属主有效。

	编译器会隐式地将book.SetPages(123)改写为(&book).SetPages(123)。
		但另一方面，我们应该总是认为aBookExpression.SetPages是一个合法的选择器（从语法层面讲），
		即使表达式aBookExpression被估值为一个不可寻址的Book值（在这种情况下，aBookExpression.SetPages是一个无效但合法的选择器）。

	如上面刚提到的，当为一个类型声明了一个方法后，每个此类型的值将拥有一个和此方法同名的成员函数。
		此类型的零值也不例外，不论此类型的零值是否用nil来表示。
*/
func (age *Age) IsNil() bool {
	return age == nil
}

func TestEg13(t *testing.T) {
	_ = StringSet(nil).Has     // 不会产生恐慌
	_ = (*Age)(nil).IsNil      // 不会产生恐慌
	_ = ((*Age)(nil)).Increase // 不会产生恐慌

	_ = StringSet(nil).Has("key") // 不会产生恐慌
	_ = ((*Age)(nil)).IsNil()     // 不会产生恐慌

	// 下面这行将产生一个恐慌，但是此恐慌不是在调用方法的时
	// 候产生的，而是在此方法体内解引用空指针的时候产生的。
	((*Age)(nil)).Increase()
}

/*
	属主参数的传参是一个值复制过程
		和普通参数传参一样，属主参数的传参也是一个值复制过程。 所以，在方法体内对属主参数的直接部分的修改将不会反映到方法体外。
*/
type Books []Book

func (books Books) Modify() {
	// 对属主参数的间接部分的修改将反映到方法之外。
	books[0].pages = 500
	// 对属主参数的直接部分的修改不会反映到方法之外。
	books = append(books, Book{789})
}

func TestEg14(t *testing.T) {
	var books = Books{{123}, {456}}
	books.Modify()
	fmt.Println(books) // [{500} {456}]
}

/*
	如果将上例中Modify方法中的两行代码次序调换，那么此方法中的两处修改都不能反映到此方法之外.
		原因是append函数调用将开辟一块新的内存来存储它返回的结果切片的元素。
			而此结果切片的前两个元素是属主参数切片的元素的副本。对此副本所做的修改不会反映到Modify方法之外。
*/

/*
	方法值的正规化
		在编译阶段，编译器将正规化各个方法值表达式。简而言之，正规化就是将方法值表达式中的隐式取地址和解引用操作均转换为显式操作。

		假设值v的类型为T，并且v.m是一个合法的方法值表达式，
			如果m是一个为类型*T显式声明的方法，那么编译器将把它正规化(&v).m；
			如果m是一个为类型T显式声明的方法，那么v.m已经是一个正规化的方法值表达式。
		假设值p的类型为*T，并且p.m是一个合法的方法值表达式，
			如果m是一个为类型T显式声明的方法，那么编译器将把它正规化(*p).m；
			如果m是一个为类型*T显式声明的方法，那么p.m已经是一个正规化的方法值表达式。


	方法值的估值
		假设v.m是一个已经正规化的方法值表达式，在运行时刻，当v.m被估值的时候，
			属主实参v的估值结果的一个副本将被存储下来以供后面调用此方法值的时候使用。

		以下面的代码为例：
			1.b.Pages是一个已经正规化的方法值表达式。 在运行时刻对其进行估值时，属主实参b的一个副本将被存储下来。
				此副本等于b的当前值：Book{pages: 123}，此后对b值的修改不影响此副本值。 这就是为什么调用f1()打印出123。

			2.在编译时刻，方法值表达式p.Pages将被正规化为(*p).Pages。 在运行时刻，属主实参*p被估值为当前的b值，
				也就是Book{pages: 123}。 这就是为什么调用f2()也打印出123。

			3.p.Pages2是一个已经正规化的方法值表达式。 在运行时刻对其进行估值时，属主实参p的一个副本将被存储下来，
				此副本的值为b值的地址。 当b被修改后，此修改可以通过对此地址值解引用而反映出来，这就是为什么调用g1()打印出789。

			4.在编译时刻，方法值表达式b.Pages2将被正规化为(&b).Pages2。 在运行时刻，
				属主实参&b的估值结果的一个副本将被存储下来，此副本的值为b值的地址。 这就是为什么调用g2()也打印出789。
*/

func (b *Book) Pages2() int {
	return (*b).Pages()
}

// func (b Book) Pages() int {
// 	return b.pages
// }
func TestEg15(t *testing.T) {

	var b = Book{pages: 123}
	var p = &b
	/*
		1.b.Pages是一个已经正规化的方法值表达式。 在运行时刻对其进行估值时，属主实参b的一个副本将被存储下来。
			此副本等于b的当前值：Book{pages: 123}，此后对b值的修改不影响此副本值。 这就是为什么调用f1()打印出123。

	*/
	var f1 = b.Pages
	/*
		2.在编译时刻，方法值表达式p.Pages将被正规化为(*p).Pages。 在运行时刻，属主实参*p被估值为当前的b值，
			也就是Book{pages: 123}。 这就是为什么调用f2()也打印出123。
	*/
	var f2 = p.Pages
	/*
		3.p.Pages2是一个已经正规化的方法值表达式。 在运行时刻对其进行估值时，属主实参p的一个副本将被存储下来，
			此副本的值为b值的地址。 当b被修改后，此修改可以通过对此地址值解引用而反映出来，这就是为什么调用g1()打印出789。

	*/
	var g1 = p.Pages2
	/*
		4.在编译时刻，方法值表达式b.Pages2将被正规化为(&b).Pages2。 在运行时刻，
			属主实参&b的估值结果的一个副本将被存储下来，此副本的值为b值的地址。 这就是为什么调用g2()也打印出789。
	*/
	var g2 = b.Pages2
	b.pages = 789
	fmt.Println(f1()) // 123
	fmt.Println(f2()) // 123
	fmt.Println(g1()) // 789
	fmt.Println(g2()) // 789

}

/*
	如何决定一个方法声明使用值类型属主还是指针类型属主？
		有时候我们必须在某些方法声明中使用指针类型属主。

		事实上，我们总可以在方法声明中使用指针类型属主而不会产生任何逻辑问题。
			我们仅仅是为了程序效率考虑有时候才会在函数声明中使用值类型属主。

		对于值类型属主还是指针类型属主都可以接受的方法声明，下面列出了一些考虑因素：
			1.太多的指针可能会增加垃圾回收器的负担。
			2.如果一个值类型的尺寸太大，那么属主参数在传参的时候的复制成本将不可忽略。 指针类型都是小尺寸类型。
			3.在并发场合下，同时调用值类型属主和指针类型属主方法比较易于产生数据竞争。
			4.sync标准库包中的类型的值不应该被复制，所以如果一个结构体类型内嵌了这些类型，
				则不应该为这个结构体类型声明值类型属主的方法。

		如果实在拿不定主意在一个方法声明中应该使用值类型属主还是指针类型属主，那么请使用指针类型属主。
*/
