package services

import (
	"context"
	"errors"

	"github.com/Re1l1x/Teams/internal/models"
	"github.com/Re1l1x/Teams/internal/repository"
	"github.com/google/uuid"
)

type GroupService struct {
	groupRepo  *repository.GroupRepository
	personRepo *repository.PersonRepository
}

func NewGroupService(g *repository.GroupRepository, p *repository.PersonRepository) *GroupService {
	return &GroupService{
		groupRepo:  g,
		personRepo: p,
	}
}

func (s *GroupService) Create(ctx context.Context, group *models.Group) error {
	if group.Name == "" {
		return errors.New("group name is empty")
	}

	if group.ParentID != nil {
		_, err := s.groupRepo.GetByID(ctx, *group.ParentID)
		if err != nil {
			return errors.New("parent group not found")
		}
	}

	return s.groupRepo.Create(ctx, group)
}

func (s *GroupService) GetByID(ctx context.Context, id uuid.UUID) (*models.Group, error) {
	return s.groupRepo.GetByID(ctx, id)
}

func (s *GroupService) GetAll(ctx context.Context) ([]models.Group, error) {
	return s.groupRepo.GetAll(ctx)
}

func (s *GroupService) Update(ctx context.Context, group *models.Group) error {
	if group.Name == "" {
		return errors.New("group name is empty")
	}

	if group.ParentID != nil {
		if *group.ParentID == group.ID {
			return errors.New("group cannot be parent of itself")
		}

		_, err := s.groupRepo.GetByID(ctx, *group.ParentID)
		if err != nil {
			return errors.New("parent group not found")
		}
	}

	return s.groupRepo.Update(ctx, group)
}

func (s *GroupService) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := s.groupRepo.GetByID(ctx, id)
	if err != nil {
		return errors.New("group not found")
	}

	return s.groupRepo.Delete(ctx, id)
}

func (s *GroupService) GetPeople(ctx context.Context, groupID uuid.UUID) ([]models.Person, error) {
	return s.groupRepo.GetPeopleByGroup(ctx, groupID)
}

func (s *GroupService) GetPeopleRecursive(ctx context.Context, groupID uuid.UUID) ([]models.Person, error) {
	return s.groupRepo.GetPeopleRecursive(ctx, groupID)
}

func (s *GroupService) CountPeople(ctx context.Context, groupID uuid.UUID) (int, error) {
	return s.groupRepo.CountPeople(ctx, groupID)
}

func (s *GroupService) CountPeopleRecursive(ctx context.Context, groupID uuid.UUID) (int, error) {
	return s.groupRepo.CountPeopleRecursive(ctx, groupID)
}
