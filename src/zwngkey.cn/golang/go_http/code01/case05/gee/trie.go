/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-29 02:31:29
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-29 02:34:44
 */
package gee

// 我们实现的动态路由具备以下两个功能:
// - 参数匹配:。例如 /p/:lang/doc，可以匹配 /p/c/doc 和 /p/go/doc。
// - 通配*。例如 /static/*filepath，可以匹配/static/fav.ico，也可以匹配/static/js/jQuery.js，这种模式常用于静态服务器，能够递归地匹配子路径。

type node struct {
	pattern  string  // 待匹配路由，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true
}
