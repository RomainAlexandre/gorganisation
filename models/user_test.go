package models_test

import (
	"github.com/gobuffalo/suite"
	"github.com/romainalexandre/gorganisation/models"
)

type ModelSuite struct {
	*suite.Model
}

func (ms *ModelSuite) Test_User_Create() {
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	u := &models.User{
		Name:       "Roro",
		Provider:   "github",
		ProviderID: "gitbubID1",
	}

	verrs, err := ms.DB.ValidateAndCreate(u)
	ms.NoError(err)
	ms.False(verrs.HasAny())

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)
}

func (ms *ModelSuite) Test_User_Create_UserExists() {
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	u := &models.User{
		Name:       "Roro",
		Provider:   "github",
		ProviderID: "gitbubID1",
	}

	verrs, err := ms.DB.ValidateAndCreate(u)
	ms.NoError(err)
	ms.False(verrs.HasAny())

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)

	u = &models.User{
		Name:       "Roro",
		Provider:   "github",
		ProviderID: "gitbubID1",
	}
	verrs, err = ms.DB.ValidateAndCreate(u)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)
}
