package actions

import (
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr/v2"
)

var r *render.Engine

func init() {
	r = render.New(render.Options{
		DefaultContentType: "application/json",

		// Box containing all of the templates:
		TemplatesBox: packr.New("app:templates", "../public"),
		AssetsBox:    packr.New("app:assets", "../public"),
	})
}
