package actions

import (
	"encoding/json"

	"github.com/romainalexandre/gorganisation/models"
)

func (as *ActionSuite) Test_UsersResource_List() {
	as.LoadFixture("lots of users")
	res := as.JSON("/users").Get()

	as.Equal(200, res.Code)
	body := res.Body.String()

	users := []models.User{}

	err := json.Unmarshal([]byte(body), &users)
	as.NoError(err)
	as.Equal(len(users), 1)
	as.Equal(users[0].Email, "roro@test.com")
}

// func (as *ActionSuite) Test_UsersResource_Show() {
// 	as.Fail("Not Implemented!")
// }

func (as *ActionSuite) Test_UsersResource_Create() {
	count, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(0, count)

	u := &models.User{
		Email:                "mark@example.com",
		Password:             "password",
		PasswordConfirmation: "password",
	}

	res := as.JSON("/users").Post(u)
	as.Equal(201, res.Code)

	count, err = as.DB.Count("users")
	as.NoError(err)
	as.Equal(1, count)
}

// func (as *ActionSuite) Test_UsersResource_Update() {
// 	as.Fail("Not Implemented!")
// }

// func (as *ActionSuite) Test_UsersResource_Destroy() {
// 	as.Fail("Not Implemented!")
// }
