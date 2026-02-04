package models

import "time"

type Membership struct {
	MembershipID int        `db:"membership_id"`
	Name         string     `db:"name"`
	Password     string     `db:"password"`
	Address      string     `db:"address"`
	IsActive     bool       `db:"is_active"`
	CreatedDate  time.Time  `db:"created_date"`
	CreatedBy    string     `db:"created_by"`
	UpdatedDate  *time.Time `db:"updated_date"`
	UpdatedBy    *string    `db:"updated_by"`
}
