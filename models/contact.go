package models

import "time"

type Contact struct {
	ContactID    int        `db:"contact_id"`
	MembershipID int        `db:"membership_id"`
	ContactType  string     `db:"contact_type"`
	ContactValue string     `db:"contact_value"`
	IsActive     bool       `db:"is_active"`
	CreatedDate  time.Time  `db:"created_date"`
	CreatedBy    string     `db:"created_by"`
	UpdatedDate  *time.Time `db:"updated_date"`
	UpdatedBy    *string    `db:"updated_by"`
}
