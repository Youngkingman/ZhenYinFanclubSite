package users

import (
	"basic/memutils/mysql"
	"fmt"
)

type UserInfo struct {
	Uid      int    `json:"uid"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUserInfo(uid int) (userinfo UserInfo) {
	db := mysql.GetDB()
	err := db.Get(&userinfo, mysql.Prefix("select * from #__userinfo where uid = ? "), uid)
	if err != nil {
		fmt.Println(err)
	}
	return
}
