package guard

import "github.com/zxdstyle/icarus/server/requests"

type sessionGuard struct {
}

func (s sessionGuard) Check(req requests.Request) error {
	//TODO implement me
	panic("implement me")
}

func (s sessionGuard) ID(req requests.Request) uint {
	//TODO implement me
	panic("implement me")
}

func (s sessionGuard) LoginUsingID(id uint) (any, error) {
	//TODO implement me
	panic("implement me")
}
