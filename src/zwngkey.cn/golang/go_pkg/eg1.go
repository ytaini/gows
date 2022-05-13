package gopkg

/*
	一个包可以由若干Go源文件组成。
	 一个包的源文件须都处于同一个目录下。
	 一个目录（不包含子目录）下的所有源文件必须都处于同一个包中，即这些源文件开头的package pkgname语句必须一致。
	 所以，一个包对应着一个目录（不包含子目录），反之亦然。 对应着一个包的目录称为此包的目录。
	 一个包目录下的每个子目录对应的都是另外一个独立的包。

	对于Go官方工具链来说，一个引入路径中包含有 internal 目录名的包被视为一个特殊的包。
	 它只能被此internal目录的直接父目录（和此父目录的子目录）中的包所引入。
	 比如，包.../a/b/c/internal/d/e/f和.../a/b/c/internal只能被引入路径含有.../a/b/c前缀的包引入。


	当一个包中的某个文件引入了另外一个包，则我们说前者包依赖于后者包。

	Go不支持循环引用（依赖）。
	 如果一个包a依赖于包b，同时包b依赖于包c，
	 则包c中的源文件不能引入包a和包b，包b中的源文件也不能引入包a。

	和包依赖类似，一个模块也可能依赖于一些其它模块。
	 此模块的直接依赖模块和这些依赖模块的版本在此模块中的go.mod文件中指定。
	  模块循环依赖是允许的，但模块循环依赖这种情况在实践中很少见。

	我们称一个程序中含有main入口函数的名称为main的包为程序代码包(main包)（或者命令代码包），称其它代码包为库代码包。
		main包不能被其它代码包引入。一个程序只能有一个main包。

	包的目录的名称并不要求一定要和其对应的包的名称相同。
	 但是，库包目录的名称最好设为和其对应的包的名称相同。
	 因为一个包的引入路径中包含的是此包的目录名，但是此包的默认引入名为此包的名称。
	  如果两者不一致，会使人感到困惑。
*/

/*
	init函数
		在一个包中，甚至一个源文件中，可以声明若干名为init的函数。 这些init函数必须不带任何输入参数和返回结果。
		注意：我们不能声明,名为init的包级变量、常量或者类型。
		在程序运行时刻，在进入main入口函数之前，每个init函数在此包加载的时候将被（串行）执行并且只执行一遍。

	程序资源初始化顺序
		一个程序中所涉及到的所有的在运行时刻要用到的包的加载是串行执行的。
		 在一个程序启动时，每个包中总是在它所有依赖的包都加载完成之后才开始加载。
		  main包总是最后一个被加载的代码包。每个被用到的包会被而且仅会被加载一次。

		在加载一个包的过程中，所有声明在此包中的init函数将被串行调用并且仅调用执行一次。
		 一个包中声明的init函数的调用肯定晚于此包所依赖的包中声明的init函数。
		  所有的init函数都将在调用main入口函数之前被调用执行。

		在同一个源文件中声明的init函数将按从上到下的顺序被调用执行。
		 对于声明在同一个包中的两个不同源文件中的两个init函数，
		 Go语言白皮书推荐（但不强求）按照它们所处于的源文件的名称的词典序列（对英文来说，即字母顺序）来调用。
		  所以最好不要让声明在同一个包中的两个不同源文件中的两个init函数存在依赖关系。

		在加载一个包的时候，此包中声明的所有包级变量都将在此包中的任何一个init函数执行之前初始化完毕。

		在同一个包内，包级变量将尽量按照它们在代码中的出现顺序被初始化，但是一个包级变量的初始化肯定晚于它所依赖的其它包级变量。
*/
// 在下面的代码片段中，四个包级变量的初始化顺序依次为y、z、x、w。
var (
	w       = x
	x, y, z = f(), 123, g()
)

func f() int {
	return z + y
}

func g() int {
	return y / 2
}

/*
	一个完整引入声明语句形式的引入名importname可以是一个句点(.)。 这样的引入称为句点引入。
		使用被句点引入的包中的导出资源时，限定标识符的前缀必须省略。

	一个完整引入声明语句形式的引入名importname可以是一个空标识符(_)。 这样的引入称为匿名引入。
		一个包被匿名引入的目的主要是为了加载这个包，从而使得这个包中的资源得以初始化。
		被匿名引入的包中的init函数将被执行并且仅执行一遍。

	模块
		一个模块（module）为若干代码包的集合。当被下载至本地后，这些代码包处于同一个目录（此模块的根目录）下。 一个模块可以有很多版本.
*/