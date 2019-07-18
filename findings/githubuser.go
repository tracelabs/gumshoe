package findings

import (
	"fmt"

	"github.com/tracelabs/gumshoe"
)

// GithubType is the name of the github username finding struct
const GithubType = "GithubUser"

// GithubUser is a user of Github
type GithubUser struct {
	Name string
}

// Investigate a Github user
func (g *GithubUser) Investigate() []gumshoe.Finding {
	found := []gumshoe.Finding{}
	// TODO
	return found
}

// GetTypeName impl.
func (g *GithubUser) GetTypeName() string {
	return GithubType
}

// GetID impl.
func (g *GithubUser) GetID() string {
	return g.Name
}

// Display a Github user
func (g *GithubUser) Display() {
	fmt.Printf("[github username] %s\n", g.Name)
}
