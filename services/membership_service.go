package services

import (
	"context"
	"errors"

	"test_crm/dto"
	"test_crm/models"
	"test_crm/repository"
	"test_crm/utils"
)

type MembershipService interface {
	Create(ctx context.Context, req dto.CreateMembershipRequest) error
	GetAll(ctx context.Context) ([]dto.MembershipResponse, error)
	GetByID(ctx context.Context, id int) (*dto.MembershipResponse, error)
	Update(ctx context.Context, id int, req dto.UpdateMembershipRequest) error
	Delete(ctx context.Context, id int) error
	GetActiveWithContact(ctx context.Context) ([]dto.MembershipWithContactResponse, error)
}

type membershipService struct {
	repo repository.MembershipRepository
}

func NewMembershipService(r repository.MembershipRepository) MembershipService {
	return &membershipService{r}
}

func (s *membershipService) Create(ctx context.Context, req dto.CreateMembershipRequest) error {
	data := models.Membership{
		Name:      req.Name,
		Password: utils.HashMD5(req.Password),
		Address:  req.Address,
		IsActive: true,
		CreatedBy: "system",
	}
	return s.repo.Create(ctx, &data)
}

func (s *membershipService) GetAll(ctx context.Context) ([]dto.MembershipResponse, error) {
	rows, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]dto.MembershipResponse, 0)
	for _, m := range rows {
		resp = append(resp, dto.MembershipResponse{
			ID:       m.MembershipID,
			Name:     m.Name,
			Address:  m.Address,
			IsActive: m.IsActive,
		})
	}
	return resp, nil
}

func (s *membershipService) GetByID(ctx context.Context, id int) (*dto.MembershipResponse, error) {
	m, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("membership not found")
	}

	return &dto.MembershipResponse{
		ID:       m.MembershipID,
		Name:     m.Name,
		Address:  m.Address,
		IsActive: m.IsActive,
	}, nil
}

func (s *membershipService) Update(ctx context.Context, id int, req dto.UpdateMembershipRequest) error {
	if id <= 0 {
		return errors.New("invalid id")
	}

	active := true
	if req.IsActive != nil {
		active = *req.IsActive
	}

	data := models.Membership{
		MembershipID: id,
		Name:         req.Name,
		Address:      req.Address,
		IsActive:     active,
		UpdatedBy:    ptr("system"),
	}

	return s.repo.Update(ctx, &data)
}

func (s *membershipService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("invalid id")
	}
	return s.repo.Delete(ctx, id)
}

func (s *membershipService) GetActiveWithContact(ctx context.Context) ([]dto.MembershipWithContactResponse, error) {

	rows, err := s.repo.FindActiveWithContactRows(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	resultMap := make(map[int]*dto.MembershipWithContactResponse)

	for rows.Next() {
		var (
			membershipID int
			name, address string
			isActive bool

			contactID int
			contactType, contactValue string
			contactIsActive bool
		)


		if err := rows.Scan(
			&membershipID,
			&name,
			&address,
			&isActive,
			&contactID,
			&contactType,
			&contactValue,
			&contactIsActive,
		); err != nil {
			return nil, err
		}


		if _, ok := resultMap[membershipID]; !ok {
			resultMap[membershipID] = &dto.MembershipWithContactResponse{
				ID:       membershipID,
				Name:     name,
				Address:  address,
				IsActive: isActive,
				Contacts: []dto.ContactItemResponse{},
			}
		}

		resultMap[membershipID].Contacts = append(
			resultMap[membershipID].Contacts,
			dto.ContactItemResponse{
				ID:           contactID,
				ContactType:  contactType,
				ContactValue: contactValue,
				IsActive:     contactIsActive,
			},
		)

	}

	resp := make([]dto.MembershipWithContactResponse, 0, len(resultMap))
	for _, v := range resultMap {
		resp = append(resp, *v)
	}

	return resp, nil
}

func ptr(v string) *string { return &v }
