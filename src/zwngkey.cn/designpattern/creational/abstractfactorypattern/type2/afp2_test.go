/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-07 01:59:32
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-07 22:01:03
 */
package type2

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
type ConcreteFactory struct {
	Level LevelType
}

type LevelType int

const (
	Level1 = iota + 1
	Level2
	Level3
)

func NewFactory(level LevelType) LevelFactory {
	// 从配置文件中读取Level的值.
	return &ConcreteFactory{Level: level}
}

func (f *ConcreteFactory) CreateEnemy() Enemy {
	switch f.Level {
	case 1:
		return &Goblin{}
	case 2:
		return &Zombie{}
	default:
		return nil
	}
}

func (f *ConcreteFactory) CreateProp() Prop {
	switch f.Level {
	case 1:
		return &Gun{}
	case 2:
		return &Bomb{}
	default:
		return nil
	}
}

func (f *ConcreteFactory) CreateBackground() Background {
	switch f.Level {
	case 1:
		return &RedBlue{}
	case 2:
		return &BlackGreen{}
	default:
		return nil
	}
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

// 这个例子中，我们定义了一个抽象工厂接口 LevelFactory，它包含了创建不同类型敌人、道具和背景的方法。
// 然后，定义了一些接口，表示不同类型的敌人、道具和背景。
// 接着我们定义了一些具体的敌人、道具和背景类型，例如 Goblin 敌人类型。
func Test(t *testing.T) {
	factory := NewFactory(Level1)
	enemy1 := factory.CreateEnemy()
	enemy1.Attack()
	prop1 := factory.CreateProp()
	prop1.Use()
	bg1 := factory.CreateBackground()
	bg1.Draw()

	factory = NewFactory(Level2)
	enemy2 := factory.CreateEnemy()
	enemy2.Attack()
	prop2 := factory.CreateProp()
	prop2.Use()
	bg2 := factory.CreateBackground()
	bg2.Draw()
}
