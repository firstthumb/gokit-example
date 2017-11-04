package responses

import(
  "github.com/firstthumb/usersrv/models"
)

type General struct {
	Response interface{} `json:"response"`
	Err      string      `json:"error"`
}

type Get struct {
	models.User
}

type Add struct {
	Success bool `json:"success"`
}

type Update struct {
	Success bool `json:"success"`
}

type Delete struct {
	Success bool `json:"success"`
}
