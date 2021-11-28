package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/svartvalp/topo-course-work/internal/pkg/book"
	"github.com/svartvalp/topo-course-work/internal/pkg/errors"
)

type BookController struct {
	bookService *book.Service
}

func NewBookController(bookService *book.Service) *BookController {
	return &BookController{bookService: bookService}
}

func (cc BookController) ListBooks(c *gin.Context) error {
	page := GetInt(c.Query("page"))
	size := GetInt(c.Query("size"))
	if page == 0 || size == 0 {
		return errors.NewStatusError(400, "page or size not defined")
	}
	name := c.Query("name")
	session := c.Query("session")
	req := book.ListBooksRequest{
		Page: page,
		Size: size,
	}
	req.Filter = &book.ListBooksFilter{
		Name:    name,
		Session: session,
	}
	books, err := cc.bookService.ListBooks(c.Request.Context(), req)
	if err != nil {
		return err
	}
	c.JSON(200, &books)
	return nil
}
