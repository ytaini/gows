/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-07 01:02:32
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-07 01:03:54
 */
package sfp

import "testing"

func Test(t *testing.T) {
	factory := &ShapeFactory{}

	rectangle := factory.NewShape("rectangle")
	rectangle.Draw()

	circle := factory.NewShape("circle")
	circle.Draw()

	triangle := factory.NewShape("triangle")
	triangle.Draw()
}
