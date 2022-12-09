/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-09 00:16:08
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-09 00:16:08
 */
package build

type (
	Profession string
	HairColor  string
	HairType   string
	Armor      string
	Weapon     string
)

type Hero struct {
	profession Profession
	name       string
	hairType   HairType
	hairColor  HairColor
	armor      Armor
	weapon     Weapon
}

// 问题: 构造函数参数的数量很快就会失控，并且可能会变得难以理解参数的排列。
// 另外，如果将来想添加更多选项，此参数列表可能会继续增长。
func NewHero(profession Profession, name string, hairType HairType, hairColor HairColor, armor Armor, weapon Weapon) *Hero {
	return &Hero{
		profession,
		name,
		hairType,
		hairColor,
		armor,
		weapon,
	}
}
