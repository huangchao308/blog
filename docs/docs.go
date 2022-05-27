// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/tag/{id}": {
            "get": {
                "description": "获取单个标签",
                "produces": [
                    "application/json"
                ],
                "summary": "获取单个标签",
                "parameters": [
                    {
                        "type": "string",
                        "description": "标签ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/model.BlogTag"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.BlogTag": {
            "type": "object",
            "properties": {
                "ID": {
                    "description": "ID",
                    "type": "integer"
                },
                "created_by": {
                    "description": "创建人",
                    "type": "string"
                },
                "created_on": {
                    "description": "创建时间",
                    "type": "string"
                },
                "deleted_on": {
                    "description": "删除时间",
                    "type": "string"
                },
                "is_del": {
                    "description": "是否删除, 0 未删除, 1 已删除",
                    "type": "integer"
                },
                "tag_name": {
                    "description": "标签名称",
                    "type": "string"
                },
                "tag_status": {
                    "description": "标签状态 0: 禁用, 1: 启用",
                    "type": "integer"
                },
                "updated_by": {
                    "description": "更新人",
                    "type": "string"
                },
                "updated_on": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}