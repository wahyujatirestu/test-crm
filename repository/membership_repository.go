package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"test_crm/models"
)

type MembershipRepository interface {
	Create(ctx context.Context, m *models.Membership) error
	FindAll(ctx context.Context) ([]models.Membership, error)
	FindByID(ctx context.Context, id int) (*models.Membership, error)
	Update(ctx context.Context, m *models.Membership) error
	Delete(ctx context.Context, id int) error
	FindByContactValue(ctx context.Context, value string) (*models.Membership, error)
	FindActiveWithContactRows(ctx context.Context) (*sqlx.Rows, error)
}

type membershipRepo struct {
	db *sqlx.DB
}

func NewMembershipRepository(db *sqlx.DB) MembershipRepository {
	return &membershipRepo{db}
}

func (r *membershipRepo) Create(ctx context.Context, m *models.Membership) error {
	query := `
	INSERT INTO membership (name,password,address,is_active,created_by)
	VALUES ($1,$2,$3,true,$4)`
	_, err := r.db.ExecContext(ctx, query,
		m.Name, m.Password, m.Address, m.CreatedBy)
	return err
}

func (r *membershipRepo) FindAll(ctx context.Context) ([]models.Membership, error) {
	var result []models.Membership
	query := `SELECT * FROM membership ORDER BY membership_id DESC`
	err := r.db.SelectContext(ctx, &result, query)
	return result, err
}

func (r *membershipRepo) FindByID(ctx context.Context, id int) (*models.Membership, error) {
	var m models.Membership
	query := `SELECT * FROM membership WHERE membership_id=$1`
	err := r.db.GetContext(ctx, &m, query, id)
	return &m, err
}

func (r *membershipRepo) Update(ctx context.Context, m *models.Membership) error {
	query := `
	UPDATE membership
	SET name=$1, address=$2, is_active=$3, updated_date=NOW(), updated_by=$4
	WHERE membership_id=$5`
	_, err := r.db.ExecContext(ctx, query,
		m.Name, m.Address, m.IsActive, m.UpdatedBy, m.MembershipID)
	return err
}

func (r *membershipRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM membership WHERE membership_id=$1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *membershipRepo) FindByContactValue(ctx context.Context, value string) (*models.Membership, error) {
	query := `
	SELECT m.*
	FROM membership m
	JOIN contact c ON c.membership_id = m.membership_id
	WHERE c.contact_value = $1
		AND c.is_active = true
		AND m.is_active = true
	LIMIT 1`

	var m models.Membership
	err := r.db.GetContext(ctx, &m, query, value)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *membershipRepo) FindActiveWithContactRows(
	ctx context.Context,
) (*sqlx.Rows, error) {

	query := `
	SELECT
	m.membership_id,
	m.name,
	m.address,
	m.is_active,
	c.contact_id,
	c.contact_type,
	c.contact_value,
	c.is_active AS contact_is_active
	FROM membership m
	JOIN contact c ON c.membership_id = m.membership_id
	WHERE m.is_active = true
	AND c.is_active = true
	ORDER BY m.membership_id;
	`

	return r.db.QueryxContext(ctx, query)
}