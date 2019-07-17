package finding

import "fmt"

const githubType = "GithubUser"

// GithubUser is a user of Github
type GithubUser struct {
	Name string
}

// GetID impl.
func (g *GithubUser) GetID() string {
	return fmt.Sprintf("%s-%s", githubType, g.Name)
}

// Investigate a Github user
func (g *GithubUser) Investigate() []Finding {
	found := []Finding{}
	// TODO
	return found
}

// Display a Github user
func (g *GithubUser) Display() {
	fmt.Printf("[github username] %s\n", g.Name)
}
