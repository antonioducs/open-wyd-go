package character

import (
	"context"
	"log/slog"

	"github.com/antonioducs/wyd/pkg/domain/entity"
	shared_gateway "github.com/antonioducs/wyd/pkg/domain/gateway"
	"github.com/antonioducs/wyd/timer-server/internal/domain/gateway"
)

type CreateCharacterUseCase struct {
	output              gateway.GameOutput
	characterRepository shared_gateway.CharacterRepository
	sessionRepository   gateway.SessionRepository
	logger              *slog.Logger
}

func NewCreateCharacterUseCase(
	output gateway.GameOutput,
	characterRepository shared_gateway.CharacterRepository,
	sessionRepository gateway.SessionRepository,
	logger *slog.Logger,
) *CreateCharacterUseCase {
	return &CreateCharacterUseCase{
		output:              output,
		sessionRepository:   sessionRepository,
		characterRepository: characterRepository,
		logger:              logger,
	}
}

type CreateCharacterInput struct {
	Context   context.Context
	SessionID uint32
	Name      string
	ClassID   uint8
	SlotID    uint8
}

func (u *CreateCharacterUseCase) Execute(input CreateCharacterInput) {
	session, ok := u.sessionRepository.Get(input.SessionID)
	if !ok {
		u.output.SendMessage(input.SessionID, "Sessão não encontrada")
		return
	}

	if len(session.Characters) >= 3 {
		u.output.SendMessage(input.SessionID, "Você já tem o máximo de personagens")
		return
	}

	hasCharacterInSlot := session.Characters.FindBySlot(input.SlotID) != nil
	if hasCharacterInSlot {
		u.output.SendMessage(input.SessionID, "Você já tem um personagem nesse slot")
		return
	}

	character, err := entity.NewCharacter(session.Account.ID, input.Name, input.ClassID, input.SlotID)
	if err != nil {
		u.output.SendMessage(input.SessionID, err.Error())
		return
	}

	err = u.characterRepository.Create(input.Context, character)
	if err != nil {
		u.logger.Error("error creating character", "error", err)
		u.output.SendMessage(input.SessionID, "Erro ao criar personagem")
		return
	}

	session.AppendCharacter(character)
	u.output.SendMessage(input.SessionID, "Personagem criado com sucesso")
	u.output.SendUpdateCharacterList(input.SessionID, session.Characters)
}
