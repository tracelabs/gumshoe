package findings

import (
	"fmt"

	"github.com/tracelabs/gumshoe"
)

// TwitterType is the name of the twitter username finding struct
const TwitterType = "TwitterUser"

// TwitterUser is a user of Twitter
type TwitterUser struct {
	Name string
}

// Investigate a Twitter user
func (t *TwitterUser) Investigate() []gumshoe.Finding {
	found := []gumshoe.Finding{}
	// TODO
	return found
}

// GetTypeName impl.
func (t *TwitterUser) GetTypeName() string {
	return TwitterType
}

// GetID impl.
func (t *TwitterUser) GetID() string {
	return t.Name
}

// Display a Twitter user
func (t *TwitterUser) Display() {
	fmt.Printf("[twitter username] %s\n", t.Name)
}
