package models

type User struct {
  FirstName   string    `json:"firstname"`
  LastName    string    `json:"lastname"`
  Age         int       `json:"age"`
}
