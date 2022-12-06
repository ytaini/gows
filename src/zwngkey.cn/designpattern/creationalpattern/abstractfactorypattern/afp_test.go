/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-07 01:59:32
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-07 03:19:18
 */
package abstractfactorypattern

import (
	"fmt"
	"testing"
)

// 假设你要开发一个游戏，这个游戏包含了多个等级，每个等级都包含了不同的敌人、道具和背景。
// 在这种情况下，你可以使用抽象工厂模式来实现。

// 首先，定义一个抽象工厂接口 LevelFactory，它包含了创建不同类型敌人、道具和背景的方法。
type LevelFactory interface {
	CreateEnemy() Enemy
	CreateProp() Prop
	CreateBackground() Background
}

// 然后，定义一些接口，表示不同类型的敌人、道具和背景。
type Enemy interface {
	Attack()
}

type Prop interface {
	Use()
}

type Background interface {
	Draw()
}

// 接下来，定义一些具体的敌人、道具和背景类型。例如，下面是一个具体的敌人类型
type Goblin struct{}

func (g *Goblin) Attack() {
	// 实现Attack方法
	fmt.Println("goblin attack")
}

type Gun struct{}

func (g *Gun) Use() {
	// 实现Use方法
	fmt.Println("gun is used")
}

type RedBlue struct{}

func (r *RedBlue) Draw() {
	// 实现Draw方法
	fmt.Println("background red blue")
}

// 然后，定义一些具体的工厂类型，例如 Level1Factory，它实现了 LevelFactory 接口，并提供了创建等级 1 中的敌人、道具和背景的具体方法。
type Level1Factory struct{}

func (f *Level1Factory) CreateEnemy() Enemy {
	return &Goblin{}
}

func (f *Level1Factory) CreateProp() Prop {
	return &Gun{}
}

func (f *Level1Factory) CreateBackground() Background {
	return &RedBlue{}
}

// 再定义一些具体的敌人、道具和背景类型
type Zombie struct{}

func (g *Zombie) Attack() {
	// 实现Attack方法
	fmt.Println("Zombie attack")
}

type Bomb struct{}

func (g *Bomb) Use() {
	// 实现Use方法
	fmt.Println("Bomb is used")
}

type BlackGreen struct{}

func (r *BlackGreen) Draw() {
	// 实现Draw方法
	fmt.Println("background BlackGreen")
}

// 再定义一些具体的工厂类型，例如 Level2Factory，它实现了 LevelFactory 接口，并提供了创建等级 2 中的敌人、道具和背景的具体方法。
type Level2Factory struct{}

func (f *Level2Factory) CreateEnemy() Enemy {
	return &Zombie{}
}

func (f *Level2Factory) CreateProp() Prop {
	return &Bomb{}
}

func (f *Level2Factory) CreateBackground() Background {
	return &BlackGreen{}
}

// 这个例子中，我们定义了一个抽象工厂接口 LevelFactory，它包含了创建不同类型敌人、道具和背景的方法。
// 然后，定义了一些接口，表示不同类型的敌人、道具和背景。
// 接着我们定义了一些具体的敌人、道具和背景类型，例如 Goblin 敌人类型。
// 然后，定义了两个具体的工厂类型 Level1Factory 和 Level2Factory，它们实现了 LevelFactory 接口，
// 并提供了创建不同等级中的敌人、道具和背景的具体方法。
func Test(t *testing.T) {
	Run(&Level1Factory{}) //通过等级1的工厂,创建等级1对应的敌人,道具,背景
	Run(&Level2Factory{}) //通过等级2的工厂,创建等级2对应的敌人,道具,背景
}

func Run(factory LevelFactory) {
	enemy := factory.CreateEnemy()
	enemy.Attack()
	prop := factory.CreateProp()
	prop.Use()
	bg := factory.CreateBackground()
	bg.Draw()
}
