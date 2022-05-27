package service

import (
	"github.com/huangchao308/blog/internal/model"
	"github.com/huangchao308/blog/pkg/app"
)

type CountTagRequest struct {
	Name   string `form:"name" binding:"required,min=2,max=100"`
	Status int8   `form:"status,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name   string `form:"name" binding:"max=100"`
	Status int8   `form:"status,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name   string `form:"name" binding:"required,min=2,max=100"`
	Status int8   `form:"status,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID     int64  `form:"id" binding:"required,gte=1"`
	Name   string `form:"name" binding:"required,min=2,max=100"`
	Status int8   `form:"status,default=1" binding:"oneof=0 1"`
}

type DeleteTagRequest struct {
	ID int64 `form:"id" binding:"required,gte=1"`
}

type GetTagRequest struct {
	ID int64 `path:"id" binding:"required,gte=1"`
}

func (s *Service) CountTag(req *CountTagRequest) (int, error) {
	return s.dao.CountTag(req.Name, req.Status)
}

func (s *Service) GetTagList(req *TagListRequest, pager *app.Pager) ([]*model.BlogTag, error) {
	return s.dao.GetTagList(req.Name, req.Status, pager.Page, pager.PageSize)
}

func (s *Service) CreateTag(req *CreateTagRequest) error {
	return s.dao.CreateTag(req.Name, req.Status)
}

func (s *Service) UpdateTag(req *UpdateTagRequest) error {
	return s.dao.UpdateTag(req.ID, req.Name, req.Status)
}

func (s *Service) DeleteTag(req *DeleteTagRequest) error {
	return s.dao.DeleteTag(req.ID)
}

func (s *Service) GetTag(id int64) (*model.BlogTag, error) {
	return s.dao.GetTag(id)
}
