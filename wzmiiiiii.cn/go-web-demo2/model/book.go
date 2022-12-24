// @author: wzmiiiiii
// @since: 2022/12/24 15:47:42
// @desc: TODO

package model

type Book struct {
	ID      int     `db:"id"`
	Title   string  `db:"title"`
	Author  string  `db:"author"`
	Price   float64 `db:"price"`
	Sales   int     `db:"sales"`
	Stock   int     `db:"stock"`
	ImgPath string  `db:"img_path"`
}
