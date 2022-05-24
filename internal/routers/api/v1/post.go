package v1

import "github.com/gin-gonic/gin"

type Post struct{}

func NewPost() *Post {
	return &Post{}
}

func (p *Post) Get(c *gin.Context) {
	c.JSON(200, gin.H{"message": "GetPosts"})
}

func (p *Post) List(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ListPosts"})
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
