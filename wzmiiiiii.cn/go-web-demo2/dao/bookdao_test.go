// @author: wzmiiiiii
// @since: 2022/12/24 16:10:12
// @desc: TODO

package dao

import (
	"log"
	"testing"
)

func TestGetBooks(t *testing.T) {
	books, err := GetBooks()
	log.Println(err)
	for _, book := range books {
		log.Println(book)
	}
}

func TestDeleteBookByID(t *testing.T) {
	log.Println(DeleteBookByID(1222))
}
