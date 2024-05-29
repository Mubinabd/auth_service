package repo

import (
	pbu "server/genproto/user"
)

type UserStoreI interface {
	CreateUser(*pbu.RegisterReq) (*pbu.UserRes, error)
}
