package requests

import(
  "github.com/firstthumb/usersrv/models"
)

type Get struct {
	ID int `json:"user_id"`
}

type Add struct {
	models.User
}

type Update struct {
	ID int
	models.User
}

type Delete struct {
	ID int
}
