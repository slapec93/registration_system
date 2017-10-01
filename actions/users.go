package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/slapec93/reg_sys/models"
	"github.com/slapec93/reg_sys/utils"
)

// UserParams ...
type UserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UsersCreate default implementation.
func UsersCreate(c buffalo.Context) error {
	userParam, err := utils.ParseRequestParams(c, UserParams{})
	if err != nil {
		return c.Render(500, r.JSON(err))
	}

	hashedPassword := utils.GeneratePassword(userParam)

	user := models.User{
		Username:       userParam.Username,
		HashedPassword: hashedPassword,
	}

	return c.Render(200, r.JSON(userParam))
}
