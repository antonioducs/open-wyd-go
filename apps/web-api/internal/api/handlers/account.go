package handlers

import (
	"net/http"

	"github.com/antonioducs/wyd/web-api/internal/api/dto"
	"github.com/antonioducs/wyd/web-api/internal/application"
	"github.com/antonioducs/wyd/web-api/internal/domain/usecase"
	"github.com/antonioducs/wyd/web-api/internal/utils/web"
)

type AccountHandler struct {
	usecases *application.UseCaseContainer
}

func NewAccountHandler(usecases *application.UseCaseContainer) *AccountHandler {
	return &AccountHandler{
		usecases: usecases,
	}
}

// HandleCreate creates a new user account
// @Summary      Create Account
// @Description  Registers a new user in the database and returns the created ID.
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        request body dto.CreateAccountRequest true "Account data"
// @Success      201  {object}  dto.CreateAccountResponse
// @Router       /accounts [post]
func (h *AccountHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {

	req, err := web.Decode[dto.CreateAccountRequest](r)
	if err != nil {
		web.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.usecases.CreateAccount.Execute(r.Context(), usecase.CreateAccountInput{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	})
	if err != nil {
		web.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	web.RespondJSON(w, http.StatusCreated, dto.CreateAccountResponse{
		ID:       resp.ID,
		Username: resp.Username,
	})
}
