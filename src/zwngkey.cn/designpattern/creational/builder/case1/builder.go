/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-08 22:14:43
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-09 00:14:09
 */
package main

// 建造者模式是一种设计模式，它用于将对象的构建与它们的表示分开。
// 这意味着在创建一个复杂对象时，可以按步骤将它们构建起来，而不必在一开始就指定它们的所有细节。

// 建造者模式通常被用来创建复杂的对象。它的优点在于能够分步骤构建一个复杂的对象，
// - 并且每一步都能够提供默认值，从而避免了在一开始就必须指定所有细节的问题。
// 另一个优点是建造者模式能够使客户端与具体的构建器实现分离。这样，客户端就无需知道构建过程的细节，只需使用指定的构建器来构建对象即可。
// 建造者模式也有一些缺点。首先，由于它分步骤构建对象，所以它可能会产生额外的复杂性。其次，它还可能会导致构建过程的控制权被分散，从而难以控制构建过程。

// 通常，建造者模式由以下几个部分组成：
// Builder 接口：定义了构建各个部分所需的方法。
// ConcreteBuilder 类：实现了构建器接口，实际上负责构建各个部分。
// Director 类：负责构建一个使用 Builder 接口的对象。
// Product 类：被构造的复杂对象。

// 例如，您可以使用建造者模式创建一个类来表示汉堡，汉堡由面包、汉堡肉、酱料和佐料组成。并允许客户选择要添加哪些配料。
// 要实现这一点，您需要定义一个接口来表示汉堡构建器，并定义具体的构建器来实现这个接口：
type BurgerBuilder interface {
	AddBun()
	AddPatty()
	AddSauce()
	AddToppings()
	GetBurger() *Burger
}

// 可以使用建造者模式来更好地控制构建过程，例如通过定义一个 Director 类来管理和控制构建过程：
// 通过这种方式，可以在不暴露构建器细节的情况下，更好地控制构建过程。
type Director struct {
	builder BurgerBuilder
}

func NewDirector(builder BurgerBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) SetBuilder(builder BurgerBuilder) {
	d.builder = builder
}

// 指挥者指挥建造者做汉堡.
func (d *Director) Construct() {
	// 标准制作流程.
	// 不同的建造者做出来的汉堡不一样(放的配料不同). 但流程一致.
	d.builder.AddBun()
	d.builder.AddPatty()
	d.builder.AddSauce()
	d.builder.AddToppings()
}

// 返回需要创建的对象(即:汉堡)
func (d *Director) GetBurger() *Burger {
	return d.builder.GetBurger()
}
