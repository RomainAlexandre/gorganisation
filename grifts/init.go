package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/romainalexandre/gorganisation/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
