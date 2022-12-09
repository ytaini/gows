/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-09 04:07:57
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-09 04:07:57
 */
package strategy

// 定义策略接口
type Strategy interface {
	DoAlgorithm([]string) []string
}

// 定义按字母顺序排序的策略类
type ConcreteStrategyA struct{}

func (a *ConcreteStrategyA) DoAlgorithm(data []string) []string {
	// 按字母顺序排序
	return data
}

// 定义按长度排序的策略类
type ConcreteStrategyB struct{}

func (b *ConcreteStrategyB) DoAlgorithm(data []string) []string {
	// 按长度排序
	return data
}

// 定义按数字大小排序的策略类
type ConcreteStrategyC struct{}

func (c *ConcreteStrategyC) DoAlgorithm(data []string) []string {
	// 按数字大小排序
	return data
}

// 定义环境类
type Context struct {
	strategy Strategy
}

// 设置策略
func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

// 执行策略
func (c *Context) ExecuteStrategy(data []string) []string {
	return c.strategy.DoAlgorithm(data)
}
