package finding

import "fmt"

const githubType = "GithubUser"

// GithubUser is a user of Github
type GithubUser struct {
	name string
}

// GetID impl.
func (g *GithubUser) GetID() string {
	return fmt.Sprintf("%s-%s", githubType, g.name)
}

// Investigate a Github user
func (g *GithubUser) Investigate() []Finding {
	found := []Finding{}
	// TODO
	return found
}
