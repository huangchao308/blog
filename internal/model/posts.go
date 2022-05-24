package model

import "time"

type BlogPosts struct {
	ID             int64     `json:"ID" gorm:"primary_key"` // ID
	CreatedBy      string    `json:"created_by"`            // 创建人
	CreatedOn      time.Time `json:"created_on"`            // 创建时间
	DeletedOn      time.Time `json:"deleted_on"`            // 删除时间
	IsDel          int8      `json:"is_del"`                // 是否删除, 0 未删除, 1 已删除
	PostAuthorId   int64     `json:"post_author_id"`        // 文章作者 ID
	PostAuthorName string    `json:"post_author_name"`      // 文章作者名字
	PostContent    string    `json:"post_content"`          // 文章内容
	PostDesc       string    `json:"post_desc"`             // 文章描述
	PostStatus     int8      `json:"post_status"`           // 文章状态 0: draft, 1: published, 2: private
	PostTitle      string    `json:"post_title"`            // 文章标题
	UpdatedBy      string    `json:"updated_by"`            // 更新人
	UpdatedOn      time.Time `json:"updated_on"`            // 更新时间
}
