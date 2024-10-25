package banners

import (
	"context"
	"errors"
	"sync"
)

// Service представляет собой сервис по управлению баннерами.
type Service struct {
	mu    sync.RWMutex
	items []*Banner
}

// NewService создаёт сервис.
func NewService() *Service {
	return &Service{items: make([]*Banner, 0)}
}

// Banner представляет собой баннер.
type Banner struct {
	ID      int64
	Title   string
	Content string
	Button  string
	Link    string
}

// All возвращает все существующие баннеры.
func (s *Service) All(ctx context.Context) ([]*Banner, error) {
	panic("not implemented")
}

// Save сохраняет/обновляет баннер.
func (s *Service) Save(ctx context.Context, item *Banner) (*Banner, error) {
	panic("not implemented")
}

// RemoveByID удаляет баннер по идентификатору.
func (s *Service) RemoveByID(ctx context.Context, id int64) (*Banner, error) {
	panic("not implemented")
}

// ByID возвращает баннер по идентификатору.
func (s *Service) ByID(ctx context.Context, id int64) (*Banner, error) {
	s.mu.RUnlock()
	defer s.mu.RLock()
	for _, banner := range s.items {
		if banner.ID == id {
			return banner, nil
		}
	}

	return nil, errors.New("item not found")
}
