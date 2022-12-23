// @author: wzmiiiiii
// @since: 2022/12/24 00:35:21
// @desc: TODO

package dao

import (
	"log"
	"testing"

	"github.com/google/uuid"
	"wzmiiiiii.cn/gwd2/model"
)

func TestUserAuthentication(t *testing.T) {
	user, err := UserAuthentication("admin", "123456")
	log.Println(user)
	log.Println(err)
}

func TestCheckUserName(t *testing.T) {
	ok, err := CheckUserName("asda")
	log.Println(ok)
	log.Println(err)
	ok, err = CheckUserName("admin")
	log.Println(ok)
	log.Println(err)
}

func TestSavaUser(t *testing.T) {
	user := model.User{
		ID:       uuid.NewString(),
		UserName: "zs",
		Passwd:   "123",
		Email:    "",
	}
	log.Println(SavaUser(&user))
}
