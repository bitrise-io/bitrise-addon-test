package loginserver

import (
	"net/http"

	"github.com/bitrise-io/bitrise-addon-test/addonprovisioner"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
)

// LoginServer ...
type LoginServer struct {
	e    *echo.Echo
	port string
}

// Start ...
func (ls *LoginServer) Start() error {
	return errors.WithStack(ls.e.Start(":" + ls.port))
}

// NewLoginServer ...
func NewLoginServer(port string, loginRequestGenFn func() (addonprovisioner.LoginRequestInfos, error)) *LoginServer {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		loginRequest, err := loginRequestGenFn()
		if err != nil {
			return errors.WithStack(err)
		}

		loginPage := `<head/>
<title>Add-on post</title>
<body>
	<form method="post" action="` + loginRequest.URL + `">
		<input type="hidden" name="timestamp" value="` + loginRequest.FormData.Timestamp + `"></input>
		<input type="hidden" name="token" value="` + loginRequest.FormData.Token + `"></input>
		<input type="hidden" name="app_slug" value="` + loginRequest.FormData.AppSlug + `"></input>
		<input type="submit" value="submit"></input>
	</form>
</body>
`
		return c.HTML(http.StatusOK, loginPage)
	})

	return &LoginServer{
		e:    e,
		port: port,
	}
}
