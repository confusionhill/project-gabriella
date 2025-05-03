package session

import (
	"com.github/confusionhill/df/private/server/internal/config"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	cfg     *config.Config
	usecase *Usecase
}

func NewHandler(cfg *config.Config, usecase *Usecase) (*Handler, error) {
	return &Handler{
		cfg:     cfg,
		usecase: usecase,
	}, nil
}

func (h *Handler) GameSettingsHandler(c echo.Context) error {
	result := h.usecase.GameSettingsUsecase(c.Request().Context())
	return c.String(200, result.ToStringValues())
}
