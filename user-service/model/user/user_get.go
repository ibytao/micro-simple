package user

import (
	"micro-simple/basic/db"
	proto "micro-simple/user-service/proto/user"

	log "github.com/micro/go-micro/v2/logger"
)

func (s *service) QueryUserByName(userName string) (ret *proto.User, err error) {
	queryString := `SELECT user_id, user_name, pwd FROM user WHERE user_name = ?`

	o := db.GetDB()

	ret = &proto.User{}

	err = o.QueryRow(queryString, userName).Scan(&ret.Id, &ret.Name, &ret.Pwd)
	if err != nil {
		log.Errorf("[QueryUserByName] 查询数据失败, err: %s", err)
		return
	}
	return
}
