package model

import "time"

type BlogPostsTag struct {
	CreatedBy string    `json:"created_by"` // 创建人
	CreatedOn time.Time `json:"created_on"` // 创建时间
	DeletedOn time.Time `json:"deleted_on"` // 删除时间
	ID        int64     `json:"ID"`         // ID
	IsDel     int8      `json:"is_del"`     // 是否删除, 0 未删除, 1 已删除
	PostId    int64     `json:"post_id"`    // 文章 ID
	TagId     int64     `json:"tag_id"`     // 标签 ID
	UpdatedBy string    `json:"updated_by"` // 更新人
	UpdatedOn time.Time `json:"updated_on"` // 更新时间
}
