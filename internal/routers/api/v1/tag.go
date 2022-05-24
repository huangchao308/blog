package v1

import "github.com/gin-gonic/gin"

type Tag struct{}

func NewTag() *Tag {
	return &Tag{}
}

func (t *Tag) Get(c *gin.Context) {
	c.JSON(200, gin.H{"message": "GetTags"})
}

func (t *Tag) List(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ListTags"})
}

func (t *Tag) Create(c *gin.Context) {
	c.JSON(200, gin.H{"message": "CreateTags"})
}

func (t *Tag) Update(c *gin.Context) {
	c.JSON(200, gin.H{"message": "UpdateTags"})
}

func (t *Tag) Delete(c *gin.Context) {
	c.JSON(200, gin.H{"message": "DeleteTags"})
}
