package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/svartvalp/topo-course-work/internal/pkg/cart"
	"github.com/svartvalp/topo-course-work/internal/pkg/errors"
)

type CartController struct {
	cartService *cart.Service
}

func NewCartController(cartService *cart.Service) *CartController {
	return &CartController{cartService: cartService}
}

func (cc CartController) GenerateSession(c *gin.Context) error {
	session, err := cc.cartService.GenerateSession(c.Request.Context())
	if err != nil {
		return err
	}
	c.JSON(200, gin.H{
		"session": session,
	})
	return nil
}

func (cc CartController) GetCart(c *gin.Context) error {
	session := c.Query("session")
	if session == "" {
		return errors.NewStatusError(400, "not specified session")
	}
	info, err := cc.cartService.GetCart(c.Request.Context(), session)
	if err != nil {
		return err
	}
	c.JSON(200, &info)
	return nil
}

func (cc CartController) AddToCart(c *gin.Context) error {
	session := c.Query("session")
	bookID := GetInt(c.Query("bookID"))
	if session == "" {
		return errors.NewStatusError(400, "not specified session")
	}
	if bookID == 0 {
		return errors.NewStatusError(400, "not specified book id")
	}
	err := cc.cartService.AddToCart(c.Request.Context(), session, int64(bookID))
	if err != nil {
		return err
	}
	c.Status(200)
	return nil
}

func (cc CartController) RemoveFromCart(c *gin.Context) error {
	session := c.Query("session")
	bookID := GetInt(c.Query("bookID"))
	if session == "" {
		return errors.NewStatusError(400, "not specified session")
	}
	if bookID == 0 {
		return errors.NewStatusError(400, "not specified book id")
	}
	err := cc.cartService.RemoveFromCart(c.Request.Context(), session, int64(bookID))
	if err != nil {
		return err
	}
	c.Status(200)
	return nil
}

func (cc CartController) GetCartCount(c *gin.Context) error {
	session := c.Query("session")
	if session == "" {
		return errors.NewStatusError(400, "not specified session")
	}
	count, err := cc.cartService.GetCartCount(c.Request.Context(), session)
	if err != nil {
		return err
	}
	c.JSON(200, gin.H{"count": count})
	return nil
}

func (cc CartController) ClearCart(c *gin.Context) error {
	session := c.Query("session")
	if session == "" {
		return errors.NewStatusError(400, "not specified session")
	}
	err := cc.cartService.ClearCart(c.Request.Context(), session)
	if err != nil {
		return err
	}
	c.Status(200)
	return nil
}
