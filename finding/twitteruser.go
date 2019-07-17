package finding

import "fmt"

const twitterType = "TwitterUser"

// TwitterUser is a user of Twitter
type TwitterUser struct {
	Name string
}

// GetID impl.
func (t *TwitterUser) GetID() string {
	return fmt.Sprintf("%s-%s", twitterType, t.Name)
}

// Investigate a Twitter user
func (t *TwitterUser) Investigate() []Finding {
	found := []Finding{}
	// TODO
	return found
}

// Display a Twitter user
func (t *TwitterUser) Display() {
	fmt.Printf("[twitter username] %s\n", t.Name)
}
