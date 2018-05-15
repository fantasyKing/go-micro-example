package service

import (
	"context"
	"fmt"

	greeter "github.com/fantasyKing/go-micro-example/proto"
)

// Hello ...
func (s *Service) Hello(ctx context.Context, req *greeter.Request) (res *greeter.Response, err error) {
	name := req.Name

	res.Msg = fmt.Sprintf("Hellp %s", name)

	return res, nil
}
