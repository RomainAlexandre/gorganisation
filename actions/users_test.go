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
	as.Equal(users[0].Name, "roro")
	as.Equal(users[0].Provider, "github")
	as.Equal(users[0].ProviderID, "githubID1")
}
