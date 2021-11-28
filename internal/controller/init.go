package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/svartvalp/topo-course-work/internal/pkg/author"
	"github.com/svartvalp/topo-course-work/internal/pkg/book"
	"github.com/svartvalp/topo-course-work/internal/pkg/cart"
)

func InitRoutes(
	r *gin.Engine,
	bookService *book.Service,
	authorService *author.Service,
	cartService *cart.Service,
) {
	bookController := NewBookController(bookService)
	authorController := NewAuthorController(authorService)
	cartController := NewCartController(cartService)

	r.GET("/books", HandleError(bookController.ListBooks))
	r.GET("/author/:id", HandleError(authorController.GetAuthor))
	r.POST("/session/generate", HandleError(cartController.GenerateSession))
	r.GET("/cart", HandleError(cartController.GetCart))
	r.GET("/cart/count", HandleError(cartController.GetCartCount))
	r.POST("/cart/add", HandleError(cartController.AddToCart))
	r.POST("/cart/remove", HandleError(cartController.RemoveFromCart))
	r.POST("/cart/clear", HandleError(cartController.ClearCart))
}
