package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/huangchao308/blog/global"
	"github.com/huangchao308/blog/internal/service"
	"github.com/huangchao308/blog/pkg/app"
	"github.com/huangchao308/blog/pkg/convert"
	"github.com/huangchao308/blog/pkg/errcode"
)

type Tag struct{}

func NewTag() *Tag {
	return &Tag{}
}

// @Summary 获取单个标签
// @Description 获取单个标签
// @Produce json
// @Param id path string true "标签ID"
// @Success 200 {object} model.BlogTag "ok"
// @Router /tag/{id} [get]
func (t *Tag) Get(c *gin.Context) {
	idParam := c.Param("id")
	id := convert.StrTo(idParam).MustInt()
	response := app.NewResponse(c)
	svr := service.New(c)
	tag, err := svr.GetTag(int64(id))
	if err != nil {
		global.Logger.Errorf("svr.GetTag err: %v", err)
		response.ToErrResponse(errcode.ErrorGetTagFail)
		return
	}
	response.ToResponse(tag)
}

func (t *Tag) List(c *gin.Context) {
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrResponse(errRsp)
		return
	}
	response.ToResponseList(nil, 0)
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
