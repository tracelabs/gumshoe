package finding

import "fmt"

const twitterType = "TwitterUser"

// TwitterUser is a user of Twitter
type TwitterUser struct {
	name string
}

// GetID impl.
func (t *TwitterUser) GetID() string {
	return fmt.Sprintf("%s-%s", twitterType, t.name)
}

// Investigate a Twitter user
func (t *TwitterUser) Investigate() []Finding {
	found := []Finding{}
	// TODO
	return found
}
