package service

import (
	"context"
	"fmt"
	"log"

	"github.com/ccsunnyfd/practice/blog/ent"
)

// Service service.
type Service struct {
	dao *ent.Client
}

// New new a service and return.
func NewService(d *ent.Client) (s *Service, cf func(), err error) {
	s = &Service{
		dao: d,
	}
	cf = s.Close
	return
}

// CreateTag CreateTag
func (s *Service) CreateTag(ctx context.Context) (*ent.Tag, error) {
	t, err := s.dao.Tag.
		Create().
		SetName("tech").
		SetCreatedBy("Chen").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("tag was created: ", t)
	return t, nil
}

// Close close the resource.
func (s *Service) Close() {
}
