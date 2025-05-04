package authentication

import (
	"encoding/xml"
	"net/http"
	"time"

	"com.github/confusionhill/df/private/server/internal/config"
	errorDto "com.github/confusionhill/df/private/server/internal/data/dto/error"
	gameDto "com.github/confusionhill/df/private/server/internal/data/dto/game"
	"com.github/confusionhill/df/private/server/internal/data/dto/ninja"
	"com.github/confusionhill/df/private/server/internal/data/entity/game"
	"com.github/confusionhill/df/private/server/internal/utilities/fable"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	cfg     *config.Config
	usecase *Usecase
}

type flash struct {
	Password string `xml:"strPassword"`
	Username string `xml:"strUsername"`
	Token    string `xml:"strToken"`
	CharID   int64  `xml:"intCharID"`
}

func NewHandler(cfg *config.Config, usecase *Usecase) (*Handler, error) {
	return &Handler{
		cfg:     cfg,
		usecase: usecase,
	}, nil
}

func (h *Handler) DeleteAccountCharacterHandler(c echo.Context) error {
	var payload ninja.NinjaDTO
	if err := xml.NewDecoder(c.Request().Body).Decode(&payload); err != nil {
		return c.String(http.StatusBadRequest, "Invalid XML")
	}
	decryptedContent, err := fable.DecryptNinja(&h.cfg.Server.DragonFable, payload.Content)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	var f flash
	if err = xml.Unmarshal([]byte(decryptedContent), &f); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	h.usecase.DeleteCharacterUsecase(c.Request().Context(), f.CharID)
	return c.String(http.StatusOK, "Character deleted successfully")
}

func (h *Handler) RegisterAccountCharacterHandler(c echo.Context) error {
	var form gameDto.CreateCharacterRequestDTO
	if err := c.Bind(&form); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid form data",
		})
	}
	if err := h.usecase.RegisterCharacterUsecase(c.Request().Context(), &game.Character{
		UserId:       form.UserID,
		Name:         form.CharacterName,
		Level:        1,
		Experience:   0,
		Strength:     1,
		Dexterity:    1,
		Intelligence: 1,
		Luck:         1,
		Charisma:     1,
		Endurance:    1,
		Wisdom:       1,
		Gender:       form.Gender,
		Pronoun:      form.Pronoun,
		HairId:       form.HairID,
		ColorHair:    "#000000",
		ColorSkin:    "#000000",
		ColorBase:    "#000000",
		ColorTrim:    "#000000",
		ClassId:      form.ClassID,
		RaceId:       form.RaceID,
		BaseClassId:  form.ClassID,
	}); err != nil {
		c.Echo().Logger.Error(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "&code=0&reason=Character+created+Successfully%21&message=none&action=none")
}

func (h *Handler) AuthenticateAccountHandler(c echo.Context) error {
	var payload ninja.NinjaDTO
	if err := xml.NewDecoder(c.Request().Body).Decode(&payload); err != nil {
		return c.String(http.StatusBadRequest, "Invalid XML")
	}
	decryptedContent, err := fable.DecryptNinja(&h.cfg.Server.DragonFable, payload.Content)
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
	bday, err := time.Parse("1/2/2006", strDOB)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	err = h.usecase.RegisterUserUsecase(c.Request().Context(), &game.User{
		Username:  strUserName,
		Password:  strPassword,
		Email:     strEmail,
		Birthdate: bday,
	})
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, resp.ToStringValues())
}
