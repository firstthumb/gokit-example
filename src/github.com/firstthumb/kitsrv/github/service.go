package github

import(
  "context"

  "fmt"
  "github.com/google/go-github/github"
)

type GithubService interface {
  GetRepositories(context.Context, string) (string, error)
}

type githubService struct{
  client  *github.Client
}

type Repository struct {
  ID    *int      `json:"id,omitempty"`
  Name  *string   `json:"name,omitempty"`
}

func New() GithubService {
  client := github.NewClient(nil)
  return githubService{client: client}
}

func (g githubService) GetRepositories(_ context.Context, s string) (string, error) {
  repos, _, err := g.client.Repositories.List(context.Background(), s, nil)
  if err != nil {
    fmt.Println("Error : ", err)
  }
  for _, repo := range repos {
    fmt.Printf("Repo => %s", *repo.Name)
  }
  return "", nil
}
