package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/svartvalp/topo-course-work/internal/controller"
	authorpkg "github.com/svartvalp/topo-course-work/internal/pkg/author"
	bookpkg "github.com/svartvalp/topo-course-work/internal/pkg/book"
	cartpkg "github.com/svartvalp/topo-course-work/internal/pkg/cart"
	"github.com/svartvalp/topo-course-work/internal/repository"
	authorrepo "github.com/svartvalp/topo-course-work/internal/repository/author"
	bookrepo "github.com/svartvalp/topo-course-work/internal/repository/book"
	cartrepo "github.com/svartvalp/topo-course-work/internal/repository/cart"
)

func main() {
	ctx := context.Background()
	db, err := repository.NewDB(ctx)
	if err != nil {
		log.Fatal(err)
	}
	cartRepo := cartrepo.New(db)
	cartService := cartpkg.New(cartRepo)

	bookRepo := bookrepo.NewRepository(db)
	bookService := bookpkg.New(bookRepo, cartService)

	authorRepo := authorrepo.New(db)
	authorService := authorpkg.New(authorRepo)

	r := gin.Default()
	controller.InitRoutes(r, bookService, authorService, cartService)
	_ = r.Run()
}
