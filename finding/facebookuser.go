package finding

import "fmt"

const facebookType = "FacebookUser"

// FacebookUser is a user of Facebook
type FacebookUser struct {
	name string
}

// GetID impl.
func (f *FacebookUser) GetID() string {
	return fmt.Sprintf("%s-%s", facebookType, f.name)
}

// Investigate a Facebook user
func (f *FacebookUser) Investigate() []Finding {
	found := []Finding{}
	// TODO
	return found
}
