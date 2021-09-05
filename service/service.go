package service

import (
	"context"
	"formllm/dao"
	"formllm/model"
)

type Service struct {
	Dao dao.DBInterface
}

func (s *Service) CreateUser(ctx context.Context, user *model.User) (id int64, err error) {
	tx := s.Dao.BeginTransaction(ctx)
	defer tx.Commit()
	_, err = s.Dao.InsertUser(tx, user)
	return
}
