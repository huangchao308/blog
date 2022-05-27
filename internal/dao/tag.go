package dao

import (
	"github.com/huangchao308/blog/internal/model"
	"github.com/huangchao308/blog/pkg/app"
)

func (d *Dao) CountTag(name string, status int8) (int, error) {
	tag := model.BlogTag{TagName: name, TagStatus: status}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, status int8, page int, pageSize int) ([]*model.BlogTag, error) {
	tag := model.BlogTag{TagName: name, TagStatus: status}
	offset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, offset, pageSize)
}

func (d *Dao) CreateTag(name string, stauts int8) error {
	tag := model.BlogTag{TagName: name, TagStatus: stauts}
	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id int64, name string, status int8) error {
	tag := model.BlogTag{ID: id}
	return tag.Update(d.engine, map[string]interface{}{
		"tag_name":   name,
		"tag_status": status,
	})
}

func (d *Dao) DeleteTag(id int64) error {
	tag := model.BlogTag{ID: id}
	return tag.Delete(d.engine)
}

func (d *Dao) GetTag(id int64) (*model.BlogTag, error) {
	tag := model.BlogTag{ID: id}
	return &tag, tag.First(d.engine)
}
