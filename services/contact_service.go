package services

import (
	"context"
	"errors"

	"test_crm/dto"
	"test_crm/models"
	"test_crm/repository"
)

type ContactService interface {
	Create(ctx context.Context, membershipID int, req dto.CreateContactRequest) error
	GetByMembership(ctx context.Context, membershipID int) ([]dto.ContactResponse, error)
	Update(ctx context.Context, id int, req dto.UpdateContactRequest) error
	Delete(ctx context.Context, id int) error
}

type contactService struct {
	repo repository.ContactRepository
}

func NewContactService(r repository.ContactRepository) ContactService {
	return &contactService{r}
}

func (s *contactService) Create(
	ctx context.Context,
	membershipID int,
	req dto.CreateContactRequest,
) error {

	if membershipID <= 0 {
		return errors.New("invalid membership id")
	}

	data := models.Contact{
		MembershipID: membershipID,
		ContactType:  req.ContactType,
		ContactValue: req.ContactValue,
		IsActive:     true,
		CreatedBy:    "system",
	}

	return s.repo.Create(ctx, &data)
}

func (s *contactService) GetByMembership(
	ctx context.Context,
	membershipID int,
) ([]dto.ContactResponse, error) {

	rows, err := s.repo.FindByMembershipID(ctx, membershipID)
	if err != nil {
		return nil, err
	}

	resp := make([]dto.ContactResponse, 0)
	for _, c := range rows {
		resp = append(resp, dto.ContactResponse{
			ID:           c.ContactID,
			ContactType:  c.ContactType,
			ContactValue: c.ContactValue,
			IsActive:     c.IsActive,
		})
	}

	return resp, nil
}

func (s *contactService) Update(ctx context.Context, id int, req dto.UpdateContactRequest) error {

	if id <= 0 {
		return errors.New("invalid contact id")
	}

	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if req.ContactValue != "" {
		existing.ContactValue = req.ContactValue
	}

	if req.IsActive != nil {
		existing.IsActive = *req.IsActive
	}

	existing.UpdatedBy = ptr("system")

	return s.repo.Update(ctx, existing)
}


func (s *contactService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("invalid contact id")
	}
	return s.repo.Delete(ctx, id)
}

