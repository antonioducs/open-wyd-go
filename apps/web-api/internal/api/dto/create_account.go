package dto

type CreateAccountRequest struct {
	Username string `json:"username" validate:"required,min=3" example:"antonio_hero"`
	Password string `json:"password" validate:"required" example:"senha_super_secreta"`
	Email    string `json:"email" validate:"required,email" example:"antonio@wyd.com"`
}

type CreateAccountResponse struct {
	ID       uint32 `json:"id" example:"105"`
	Username string `json:"username" example:"antonio_hero"`
}
