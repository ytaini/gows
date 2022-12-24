// @author: wzmiiiiii
// @since: 2022/12/24 15:57:36
// @desc: TODO

package dao

import (
	"wzmiiiiii.cn/gwd2/common"
	"wzmiiiiii.cn/gwd2/model"
)

func GetBooks() ([]*model.Book, error) {
	sqlStr := `select * from books`
	rows, err := common.Db.Queryx(sqlStr)
	if err != nil {
		return nil, err
	}
	books := make([]*model.Book, 0)
	for rows.Next() {
		var book model.Book
		err := rows.StructScan(&book)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	return books, nil
}

func SaveBook(book *model.Book) (err error) {
	sqlStr := `insert into books(title, author, price, sales, stock,img_path) values (:title,:author,:price,:sales,:stock,:img_path) `
	_, err = common.Db.NamedExec(sqlStr, book)
	return err
}

func DeleteBookByID(id int) (err error) {
	sqlStr := `delete from books where id = ?`
	_, err = common.Db.Exec(sqlStr, id)
	return
}
