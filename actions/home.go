package actions

import (
	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	// var message string
	// if currentUser, ok := c.Data()["currentUser"].(*models.User); ok {
	// 	message = "Welcome to Bufallo: " + currentUser.Name
	// } else {
	// 	message = "Welcome to Bufallo! Please Log in"
	// }
	// return c.Render(200, r.JSON(map[string]string{"message": message}))
	return c.Render(200, r.HTML("index.html"))
}
