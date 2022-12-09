/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-09 00:08:10
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-09 00:08:11
 */
package main

type ConcreteBurgerBuilder1 struct {
	burger *Burger
}

func NewConcreteBurgerBuilder1() *ConcreteBurgerBuilder1 {
	return &ConcreteBurgerBuilder1{
		burger: new(Burger),
	}
}

func (b *ConcreteBurgerBuilder1) AddBun() {
	b.burger.bun = "bun1"
}

func (b *ConcreteBurgerBuilder1) AddPatty() {
	b.burger.patty = " patty1"
}

func (b *ConcreteBurgerBuilder1) AddSauce() {
	b.burger.sauce = "sauce1"
}

func (b *ConcreteBurgerBuilder1) AddToppings() {
	b.burger.toppings = "toppings1"
}

func (b *ConcreteBurgerBuilder1) GetBurger() *Burger {
	return b.burger
}
