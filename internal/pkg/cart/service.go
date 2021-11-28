package cart

import (
	"context"

	"github.com/hashicorp/go-uuid"
	"github.com/svartvalp/topo-course-work/internal/pkg/errors"
)

type Service struct {
	cartRepo cartAdapter
}

type cartAdapter interface {
	GetBySession(ctx context.Context, session string) (Cart, error)
	GetCartBooks(ctx context.Context, id int64) ([]BookWithCount, error)
	DeleteCartBooks(ctx context.Context, id int64) error
	GetBookInCart(ctx context.Context, cartID int64, bookID int64) (BookInCart, error)
	DeleteBookFromCart(ctx context.Context, cartID int64, bookID int64) error
	CreateOrUpdateBookInCart(ctx context.Context, book BookInCart) error
	Create(ctx context.Context, cart Cart) (Cart, error)
}

func New(
	cartRepo cartAdapter) *Service {
	return &Service{cartRepo: cartRepo}
}

func (s *Service) GenerateSession(ctx context.Context) (string, error) {
	return uuid.GenerateUUID()
}

func (s *Service) GetCart(ctx context.Context, session string) (Info, error) {
	cart, err := s.cartRepo.GetBySession(ctx, session)
	if err != nil {
		if errors.Status(err) == 404 {
			return Info{
				SessionID: session,
				Books:     make([]BookWithCount, 0),
			}, nil
		}
		return Info{}, err
	}
	books, err := s.cartRepo.GetCartBooks(ctx, cart.ID)
	if err != nil {
		return Info{}, err
	}
	return Info{
		SessionID: cart.SessionID,
		Books:     books,
	}, nil
}

func (s Service) GetCartCount(ctx context.Context, session string) (int, error) {
	cart, err := s.cartRepo.GetBySession(ctx, session)
	if err != nil && errors.Status(err) != 404 {
		return 0, err
	}
	if err != nil {
		return 0, nil
	}
	books, err := s.cartRepo.GetCartBooks(ctx, cart.ID)
	if err != nil {
		return 0, err
	}
	count := 0
	for _, b := range books {
		count += b.Count
	}
	return count, nil
}

func (s *Service) AddToCart(ctx context.Context, session string, bookID int64) error {
	cart, err := s.cartRepo.GetBySession(ctx, session)
	if err != nil && errors.Status(err) != 404 {
		return err
	}
	if err != nil {
		cart, err = s.cartRepo.Create(ctx, Cart{SessionID: session})
		if err != nil {
			return err
		}
	}

	bookInCart, err := s.cartRepo.GetBookInCart(ctx, cart.ID, bookID)
	if err != nil && errors.Status(err) != 404 {
		return err
	}
	if err != nil {
		bookInCart = BookInCart{
			CartID: cart.ID,
			BookID: bookID,
			Count:  1,
		}
	} else {
		bookInCart.Count += 1
	}

	err = s.cartRepo.CreateOrUpdateBookInCart(ctx, bookInCart)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) RemoveFromCart(ctx context.Context, session string, bookID int64) error {
	cart, err := s.cartRepo.GetBySession(ctx, session)
	if err != nil {
		return err
	}
	bookInCart, err := s.cartRepo.GetBookInCart(ctx, cart.ID, bookID)
	if err != nil {
		return err
	}
	if bookInCart.Count == 1 {
		err = s.cartRepo.DeleteBookFromCart(ctx, cart.ID, bookID)
		if err != nil {
			return err
		}
	} else {
		bookInCart.Count -= 1
		err = s.cartRepo.CreateOrUpdateBookInCart(ctx, bookInCart)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s Service) ClearCart(ctx context.Context, session string) error {
	cart, err := s.cartRepo.GetBySession(ctx, session)
	if err != nil {
		return err
	}
	err = s.cartRepo.DeleteCartBooks(ctx, cart.ID)
	if err != nil {
		return err
	}
	return nil
}
