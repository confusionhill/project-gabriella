package game

import (
	"encoding/xml"
	"net/http"

	"com.github/confusionhill/df/private/server/internal/config"
	"com.github/confusionhill/df/private/server/internal/data/dto/ninja"
	"com.github/confusionhill/df/private/server/internal/utilities/fable"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	cfg *config.Config
}

func NewHandler(cfg *config.Config) (*Handler, error) {
	return &Handler{
		cfg: cfg,
	}, nil
}

func (h *Handler) LoadCharacterHandler(c echo.Context) error {
	var payload ninja.NinjaDTO
	if err := xml.NewDecoder(c.Request().Body).Decode(&payload); err != nil {
		return c.String(http.StatusBadRequest, "Invalid XML")
	}
	decryptedContent, err := fable.DecryptNinja(&h.cfg.Server.DragonFable, payload.Content)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, decryptedContent)
}
