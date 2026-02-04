package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"test_crm/models"
)

type ContactRepository interface {
	Create(ctx context.Context, c *models.Contact) error
	FindByMembershipID(ctx context.Context, membershipID int) ([]models.Contact, error)
	FindByID(ctx context.Context, id int) (*models.Contact, error)
	Update(ctx context.Context, c *models.Contact) error
	Delete(ctx context.Context, id int) error
}

type contactRepo struct {
	db *sqlx.DB
}

func NewContactRepository(db *sqlx.DB) ContactRepository {
	return &contactRepo{db}
}

func (r *contactRepo) Create(ctx context.Context, c *models.Contact) error {
	query := `
	INSERT INTO contact (membership_id, contact_type, contact_value, is_active, created_by)
	VALUES ($1,$2,$3,true,$4)`
	_, err := r.db.ExecContext(ctx, query,
		c.MembershipID, c.ContactType, c.ContactValue, c.CreatedBy)
	return err
}

func (r *contactRepo) FindByMembershipID(ctx context.Context, membershipID int) ([]models.Contact, error) {
	var rows []models.Contact
	query := `
	SELECT * FROM contact
	WHERE membership_id=$1
	ORDER BY contact_id DESC`
	err := r.db.SelectContext(ctx, &rows, query, membershipID)
	return rows, err
}

func (r *contactRepo) FindByID(ctx context.Context, id int) (*models.Contact, error) {
	var c models.Contact
	query := `SELECT * FROM contact WHERE contact_id=$1`
	err := r.db.GetContext(ctx, &c, query, id)
	return &c, err
}

func (r *contactRepo) Update(ctx context.Context, c *models.Contact) error {
	query := `
	UPDATE contact
	SET contact_value=$1, is_active=$2, updated_date=NOW(), updated_by=$3
	WHERE contact_id=$4`
	_, err := r.db.ExecContext(ctx, query,
		c.ContactValue, c.IsActive, c.UpdatedBy, c.ContactID)
	return err
}

func (r *contactRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM contact WHERE contact_id=$1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
