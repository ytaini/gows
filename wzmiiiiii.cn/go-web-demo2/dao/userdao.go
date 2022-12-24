// @author: wzmiiiiii
// @since: 2022/12/24 00:18:57
// @desc: TODO

package dao

import (
	"database/sql"

	"wzmiiiiii.cn/gwd2/common"
	"wzmiiiiii.cn/gwd2/model"
)

// UserAuthentication 身份认证
func UserAuthentication(username, passwd string) (user *model.User, err error) {
	user = &model.User{}
	sqlStr := `select * from users where username = ? and passwd = ?`
	err = common.Db.Get(user, sqlStr, username, passwd)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return
}

func CheckUserName(username string) (ok bool, err error) {
	user := model.User{}
	sqlStr := `select * from users where username = ?`
	if err = common.Db.Get(&user, sqlStr, username); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func SavaUser(user *model.User) (err error) {
	sqlStr := `insert into users(id, username, passwd, email) values (:id,:username,:passwd,:email)`
	_, err = common.Db.NamedExec(sqlStr, user)
	return
}
