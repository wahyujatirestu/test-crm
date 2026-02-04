package dto

type CreateContactRequest struct {
	ContactType  string `json:"contact_type" binding:"required,oneof=email phone"`
	ContactValue string `json:"contact_value" binding:"required,min=5,max=100"`
}

type UpdateContactRequest struct {
	ContactValue string `json:"contact_value" binding:"required,min=5,max=100"`
	IsActive     *bool  `json:"is_active,omitempty"`
}

type ContactResponse struct {
	ID           int    `json:"contact_id"`
	ContactType  string `json:"contact_type"`
	ContactValue string `json:"contact_value"`
	IsActive     bool   `json:"is_active"`
}
