/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-09 00:08:56
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-09 00:08:56
 */
package main

type ConcreteBurgerBuilder2 struct {
	burger *Burger
}

func NewConcreteBurgerBuilder2() *ConcreteBurgerBuilder2 {
	return &ConcreteBurgerBuilder2{
		burger: new(Burger),
	}
}

func (b *ConcreteBurgerBuilder2) AddBun() {
	b.burger.bun = "bun2"
}

func (b *ConcreteBurgerBuilder2) AddPatty() {
	b.burger.patty = " patty2"
}

func (b *ConcreteBurgerBuilder2) AddSauce() {
	b.burger.sauce = "sauce2"
}

func (b *ConcreteBurgerBuilder2) AddToppings() {
	b.burger.toppings = "toppings2"
}

func (b *ConcreteBurgerBuilder2) GetBurger() *Burger {
	return b.burger
}
