package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/huangchao308/blog/pkg/app"
	"github.com/huangchao308/blog/pkg/errcode"
)

type Post struct{}

func NewPost() *Post {
	return &Post{}
}

func (p *Post) Get(c *gin.Context) {
	app.NewResponse(c).ToErrResponse(errcode.NotFound)
}

func (p *Post) List(c *gin.Context) {
	app.NewResponse(c).ToErrResponse(errcode.NotFound)
}

func (p *Post) Create(c *gin.Context) {
	c.JSON(200, gin.H{"message": "CreatePosts"})
}

func (p *Post) Update(c *gin.Context) {
	c.JSON(200, gin.H{"message": "UpdatePosts"})
}

func (p *Post) Delete(c *gin.Context) {
	c.JSON(200, gin.H{"message": "DeletePosts"})
}
