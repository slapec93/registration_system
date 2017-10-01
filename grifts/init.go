package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/slapec93/reg_sys/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
