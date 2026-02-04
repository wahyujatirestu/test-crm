package dto

type ContactItemResponse struct {
	ID           int    `json:"contact_id"`
	ContactType  string `json:"contact_type"`
	ContactValue string `json:"contact_value"`
	IsActive     bool   `json:"is_active"`
}


type MembershipWithContactResponse struct {
	ID       int                   `json:"membership_id"`
	Name     string                `json:"name"`
	Address  string                `json:"address"`
	IsActive bool                  `json:"is_active"`
	Contacts []ContactItemResponse `json:"contacts"`
}
