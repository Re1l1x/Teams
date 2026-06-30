package services

import (
	"context"
	"errors"

	"github.com/Re1l1x/Teams/internal/models"
	"github.com/Re1l1x/Teams/internal/repository"
	"github.com/google/uuid"
)

type PersonService struct {
	personRepo *repository.PersonRepository
	groupRepo  *repository.GroupRepository
}

func NewPersonService(p *repository.PersonRepository, g *repository.GroupRepository) *PersonService {
	return &PersonService{
		personRepo: p,
		groupRepo:  g,
	}
}

func (s *PersonService) Create(ctx context.Context, person *models.Person) error {

	if person.FirstName == "" || person.LastName == "" {
		return errors.New("invalid person name")
	}

	if person.BirthYear <= 0 {
		return errors.New("invalid birth year")
	}

	_, err := s.groupRepo.GetByID(ctx, person.GroupID)
	if err != nil {
		return errors.New("group not found")
	}

	return s.personRepo.Create(ctx, person)
}

func (s *PersonService) GetByID(ctx context.Context, id uuid.UUID) (*models.Person, error) {
	return s.personRepo.GetByID(ctx, id)
}

func (s *PersonService) GetAll(ctx context.Context) ([]models.Person, error) {
	return s.personRepo.GetAll(ctx)
}

func (s *PersonService) Update(ctx context.Context, person *models.Person) error {

	if person.FirstName == "" || person.LastName == "" {
		return errors.New("invalid person name")
	}

	if person.BirthYear <= 0 {
		return errors.New("invalid birth year")
	}

	_, err := s.groupRepo.GetByID(ctx, person.GroupID)
	if err != nil {
		return errors.New("group not found")
	}

	return s.personRepo.Update(ctx, person)
}

func (s *PersonService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.personRepo.Delete(ctx, id)
}

func (s *PersonService) GetByGroup(ctx context.Context, groupID uuid.UUID) ([]models.Person, error) {
	return s.personRepo.GetByGroup(ctx, groupID)
}

func (s *PersonService) GetByGroups(ctx context.Context, groupIDs []uuid.UUID) ([]models.Person, error) {
	return s.personRepo.GetByGroups(ctx, groupIDs)
}
