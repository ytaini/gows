/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-08 23:05:42
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-09 00:16:25
 */
package build

// 合理的替代方案是使用Builder模式
type HeroBuilder struct {
	hero *Hero
}

func NewHeroBuilder(profession Profession, name string) *HeroBuilder {
	// 默认值
	hairType := "中分"
	hairColor := "黑色"
	armor := "盔甲1"
	weapon := "枪"

	hero := &Hero{
		profession: profession,
		name:       name,
		hairType:   HairType(hairType),
		hairColor:  HairColor(hairColor),
		armor:      Armor(armor),
		weapon:     Weapon(weapon),
	}
	return &HeroBuilder{
		hero: hero,
	}
}

func (b *HeroBuilder) WithHairType(hairType HairType) *HeroBuilder {
	b.hero.hairType = hairType
	return b
}

func (b *HeroBuilder) WithHairColor(hairColor HairColor) *HeroBuilder {
	b.hero.hairColor = hairColor
	return b
}

func (b *HeroBuilder) Build() *Hero {
	return b.hero
}
