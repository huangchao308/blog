use blog;
CREATE TABLE if not exists `blog_posts` (
    `ID` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `post_author_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '文章作者 ID',
    `post_author_name` varchar(100) NOT NULL DEFAULT '' COMMENT '文章作者名字',
    `post_content` longtext COMMENT '文章内容',
    `post_title` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文章标题',
    `post_status` tinyint(3) unsigned NOT NULL DEFAULT 0 COMMENT '文章状态 0: draft, 1: published, 2: private',
    `post_desc` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '文章描述',
    `created_on` datetime NOT NULL DEFAULT current_timestamp COMMENT '创建时间',
    `updated_on` datetime NOT NULL DEFAULT current_timestamp COMMENT '更新时间',
    `created_by` varchar(100) NOT NULL DEFAULT '' COMMENT '创建人',
    `updated_by` varchar(100) NOT NULL DEFAULT '' COMMENT '更新人',
    `deleted_on` datetime NOT NULL DEFAULT current_timestamp COMMENT '删除时间',
    `is_del` tinyint(3) unsigned DEFAULT 0 COMMENT '是否删除, 0 未删除, 1 已删除',
	PRIMARY KEY (`ID`),
	KEY (`post_author_id`, `is_del`, `post_status`, `created_on`),
	KEY (`post_author_id`, `is_del`, `deleted_on`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '博客文章';

CREATE TABLE if not exists `blog_tag` (
    `ID` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `tag_name` varchar(100) NOT NULL DEFAULT '' COMMENT '标签名称',
    `tag_status` tinyint(3) unsigned NOT NULL DEFAULT 1 COMMENT '标签状态 0: 禁用, 1: 启用',
    `created_on` datetime NOT NULL DEFAULT current_timestamp COMMENT '创建时间',
    `updated_on` datetime NOT NULL DEFAULT current_timestamp COMMENT '更新时间',
    `created_by` varchar(100) NOT NULL DEFAULT '' COMMENT '创建人',
    `updated_by` varchar(100) NOT NULL DEFAULT '' COMMENT '更新人',
    `deleted_on` datetime NOT NULL DEFAULT current_timestamp COMMENT '删除时间',
    `is_del` tinyint(3) unsigned DEFAULT 0 COMMENT '是否删除, 0 未删除, 1 已删除',
    PRIMARY KEY (`ID`),
    KEY (`tag_status`, `is_del`, `created_on`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '博客标签';

CREATE TABLE if not exists `blog_posts_tag` (
    `ID` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `post_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '文章 ID',
    `tag_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '标签 ID',
    `created_on` datetime NOT NULL DEFAULT current_timestamp COMMENT '创建时间',
    `updated_on` datetime NOT NULL DEFAULT current_timestamp COMMENT '更新时间',
    `created_by` varchar(100) NOT NULL DEFAULT '' COMMENT '创建人',
    `updated_by` varchar(100) NOT NULL DEFAULT '' COMMENT '更新人',
    `deleted_on` datetime NOT NULL DEFAULT current_timestamp COMMENT '删除时间',
    `is_del` tinyint(3) unsigned DEFAULT 0 COMMENT '是否删除, 0 未删除, 1 已删除',
    PRIMARY KEY (`ID`),
    KEY (`post_id`, `updated_on`),
    KEY (`tag_id`, `updated_on`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '博客文章标签关联';
