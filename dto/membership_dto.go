package dto

type CreateMembershipRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=100"`
	Password string `json:"password" binding:"required,min=6,max=50"`
	Address  string `json:"address" binding:"required,min=5"`
}

type UpdateMembershipRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=100"`
	Address  string `json:"address" binding:"required,min=5"`
	IsActive *bool  `json:"is_active,omitempty"`
}

type MembershipResponse struct {
	ID       int    `json:"membership_id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	IsActive bool   `json:"is_active"`
}

type MembershipDetailResponse struct {
	ID        int    `json:"membership_id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	IsActive  bool   `json:"is_active"`
	CreatedBy string `json:"created_by"`
}