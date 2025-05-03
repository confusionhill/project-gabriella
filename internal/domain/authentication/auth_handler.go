package authentication

import (
	"encoding/xml"
	"net/http"

	utils "com.github/confusionhill/df/private/server/internal/Utils"
	"com.github/confusionhill/df/private/server/internal/config"
	errorDto "com.github/confusionhill/df/private/server/internal/data/dto/error"
	gameDto "com.github/confusionhill/df/private/server/internal/data/dto/game"
	"com.github/confusionhill/df/private/server/internal/data/dto/ninja"
	"com.github/confusionhill/df/private/server/internal/data/entity/game"
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

func (h *Handler) RegisterAccountCharacterHandler(c echo.Context) error {
	return nil
}

type flash struct {
	Password string `xml:"strPassword"`
	Username string `xml:"strUsername"`
}

func (h *Handler) AuthenticateAccountHandler(c echo.Context) error {
	var payload ninja.NinjaDTO
	if err := xml.NewDecoder(c.Request().Body).Decode(&payload); err != nil {
		return c.String(http.StatusBadRequest, "Invalid XML")
	}
	decryptedContent, err := utils.DecryptNinja(&h.cfg.Server.DragonFable, payload.Content)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	var f flash
	if err = xml.Unmarshal([]byte(decryptedContent), &f); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	user, err := h.usecase.AuthenticateUserUsecase(c.Request().Context(), f.Username, f.Password)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	authUserResp := gameDto.NewAuthUserDTO(user)
	content, err := authUserResp.ToString(h.cfg)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	//return c.String(http.StatusOK, content)
	response := ninja.NinjaDTO{
		Content: content,
	}

	return c.XML(http.StatusOK, response)
}

func (h *Handler) RegisterAccountHandler(c echo.Context) error {
	strDOB := c.FormValue("strDOB")
	if strDOB == "" {
		return c.String(http.StatusBadRequest, "Birthdate is required")
	}
	strUserName := c.FormValue("strUserName")
	if strUserName == "" {
		return c.String(http.StatusBadRequest, "Username is required")
	}
	strPassword := c.FormValue("strPassword")
	if strPassword == "" {
		return c.String(http.StatusBadRequest, "Password is required")
	}
	strEmail := c.FormValue("strEmail")
	if strEmail == "" {
		return c.String(http.StatusBadRequest, "Email is required")
	}

	resp := errorDto.ErrorResponseDTO{
		Code: errorDto.INVALID_EMAIL,
	}
	err := h.usecase.RegisterUserUsecase(c.Request().Context(), &game.User{
		Username:  strUserName,
		Password:  strPassword,
		Email:     strEmail,
		BirthDate: strDOB,
	})
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, resp.ToStringValues())
}
