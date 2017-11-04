package service

import(
  "errors"

  "github.com/firstthumb/usersrv/models"
  "github.com/firstthumb/usersrv/requests"
  "github.com/firstthumb/usersrv/responses"
)

type UserService interface {
  Get(req *requests.Get) (*responses.Get, error)
  Add(req *requests.Add) (*responses.Add, error)
  Delete(req *requests.Delete) (*responses.Delete, error)
  Update(req *requests.Update) (*responses.Update, error)
}

type userService struct {
  userMap       map[int]*models.User
  incrementID   int
}

func New() UserService {
  return &userService{make(map[int]*models.User), 1}
}

func (svc userService) Get(req *requests.Get) (*responses.Get, error) {
  if user, exists := svc.userMap[req.ID]; exists {
    return &responses.Get{*user}, nil
  } else {
    return nil, errors.New("No user found")
  }
}

func (svc userService) Add(req *requests.Add) (*responses.Add, error) {
  svc.userMap[svc.incrementID] = &req.User
  svc.incrementID++
  return &responses.Add{true}, nil
}

func (svc userService) Update(req *requests.Update) (*responses.Update, error) {
	if _, exists := svc.userMap[req.ID]; exists {
		svc.userMap[req.ID] = &req.User
	} else {
		return nil, errors.New("No such user.")
	}
	return &responses.Update{true}, nil
}

func (svc *userService) Delete(req *requests.Delete) (*responses.Delete, error) {
	if _, exists := svc.userMap[req.ID]; exists {
		delete(svc.userMap, req.ID)
	} else {
		return nil, errors.New("No such user.")
	}
	return &responses.Delete{true}, nil
}
