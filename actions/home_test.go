package actions

func (as *ActionSuite) Test_HomeHandler() {
	res := as.JSON("/").Get()

	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), `<div id=\"root\"></div>`)
}
