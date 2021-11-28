package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/svartvalp/topo-course-work/internal/pkg/author"
	"github.com/svartvalp/topo-course-work/internal/pkg/errors"
)

type AuthorController struct {
	authorService *author.Service
}

func NewAuthorController(authorService *author.Service) *AuthorController {
	return &AuthorController{authorService: authorService}
}

func (cc *AuthorController) GetAuthor(c *gin.Context) error {
	id := GetInt(c.Param("id"))
	if id == 0 {
		return errors.NewStatusError(400, "id not specified")
	}
	a, err := cc.authorService.GetAuthor(c.Request.Context(), int64(id))
	if err != nil {
		return err
	}
	c.JSON(200, &a)
	return nil
}
