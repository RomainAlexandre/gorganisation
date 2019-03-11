package actions

import (
	"fmt"
	"os"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"github.com/pkg/errors"
	"github.com/romainalexandre/gorganisation/models"
)

func init() {
	gothic.Store = App().SessionStore

	goth.UseProviders(
		github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/github/callback")),
	)
}

// AuthCallback saves in the session the user id after a provider callback
func AuthCallback(c buffalo.Context) error {
	gothUser, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Error(401, err)
	}

	tx := c.Value("tx").(*pop.Connection)

	q := tx.Where("provider = ? and provider_id = ?", gothUser.Provider, gothUser.UserID)
	exists, err := q.Exists("users")

	if err != nil {
		return c.Error(422, err)
	}

	u := &models.User{}
	if exists {
		err = q.First(u)

		if err != nil {
			return errors.WithStack(err)
		}
	} else {
		u.Name = gothUser.Name
		u.Provider = gothUser.Provider
		u.ProviderID = gothUser.UserID

		verr, err := tx.ValidateAndSave(u)

		if err != nil {
			return errors.WithStack(err)
		}

		if verr.HasAny() {
			return c.Render(422, r.JSON(verr))
		}
	}

	c.Session().Set("current_user_id", u.ID)

	err = c.Session().Save()
	if err != nil {
		return errors.WithStack(err)
	}

	return c.Redirect(302, "/")
}

// AuthDestroy destroys the session of the current user, thus log him out
func AuthDestroy(c buffalo.Context) error {
	c.Session().Clear()
	err := c.Session().Save()
	if err != nil {
		return errors.WithStack(err)
	}

	return c.Render(200, r.JSON("You have been logged out"))
}

// SetCurrentUser sets the current user in the context
func SetCurrentUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if userID := c.Session().Get("current_user_id"); userID != nil {
			user := &models.User{}
			tx := c.Value("tx").(*pop.Connection)
			err := tx.Find(user, userID)

			if err != nil {
				return errors.WithStack(err)
			}
			c.Set("currentUser", user)
		}
		return next(c)
	}
}

// Authorize verifies that the current user have access to the resource
func Authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if userID := c.Session().Get("current_user_id"); userID == nil {
			c.Render(401, r.JSON("You need to log in to access this resource"))
		}
		return next(c)
	}
}
