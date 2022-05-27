package service

import (
	"context"

	"github.com/huangchao308/blog/global"
	"github.com/huangchao308/blog/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) *Service {
	svc := &Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
