package router

import (
	"github.com/labstack/echo/v4"
)

type PagesRouter struct {
}

func NewPagesRouter() (*PagesRouter, error) {
	return &PagesRouter{}, nil
}

func (r *PagesRouter) Setup(e *echo.Echo) {
	//MARK: HOME PAGE
	e.GET("", func(c echo.Context) error {
		return c.File("public/pages/index.html")
	})
	//MARK: GAME PAGE
	e.GET("/game", func(c echo.Context) error {
		return c.File("public/pages/game.html")
	})
	//MARK: REGISTER PAGE
	e.GET("/register", func(c echo.Context) error {
		return c.File("public/pages/signup.html")
	})

	// e.GET("/server-emulator/server.php/DFversion.asp", func(c echo.Context) error {
	// 	response := `&gamemovie=game15_9_42-patched.swf&signUpMessage=Welcome+to+the+world+of+DragonFable%21%0A%0APlease+sign+up+to+play%21&server=server-emulator%2Fserver.php%2F&gamefilesPath=cdn%2Fgamefiles%2F&gameVersion=Build+15.9.42+alpha&online=true&end=here`
	// 	return c.String(http.StatusOK, response)
	// })
}
