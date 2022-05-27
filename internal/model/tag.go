package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type BlogTag struct {
	CreatedBy string    `json:"created_by"` // 创建人
	CreatedOn time.Time `json:"created_on"` // 创建时间
	DeletedOn time.Time `json:"deleted_on"` // 删除时间
	ID        int64     `json:"ID"`         // ID
	IsDel     int8      `json:"is_del"`     // 是否删除, 0 未删除, 1 已删除
	TagName   string    `json:"tag_name"`   // 标签名称
	TagStatus int8      `json:"tag_status"` // 标签状态 0: 禁用, 1: 启用
	UpdatedBy string    `json:"updated_by"` // 更新人
	UpdatedOn time.Time `json:"updated_on"` // 更新时间
}

func (t *BlogTag) TableName() string {
	return "blog_tag"
}

func (t *BlogTag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.TagName != "" {
		db = db.Where("tag_name = ?", t.TagName)
	}
	db = db.Where("tag_status = ?", t.TagStatus)
	err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (t *BlogTag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t *BlogTag) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&t).Where("id = ?", t.ID).Updates(values).Error
}

func (t *BlogTag) List(db *gorm.DB, offset int, limit int) ([]*BlogTag, error) {
	var tags []*BlogTag
	var err error
	if offset > 0 && limit > 0 {
		db = db.Offset(offset).Limit(limit)
	}
	if t.TagName != "" {
		db = db.Where("tag_name = ?", t.TagName)
	}
	db = db.Where("tag_status = ?", t.TagStatus)
	err = db.Where("is_del = ?", 0).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (t *BlogTag) Delete(db *gorm.DB) error {
	return db.Where("id = ?", t.ID).Delete(&t).Error
}

func (t *BlogTag) First(db *gorm.DB) error {
	return db.First(&t).Error
}
