package findings

import (
	"fmt"

	"github.com/tracelabs/gumshoe"
)

// FacebookType is the name of the facebook username finding struct
const FacebookType = "FacebookUser"

// FacebookUser is a user of Facebook
type FacebookUser struct {
	Name string
}

// Investigate a Facebook user
func (f *FacebookUser) Investigate() []gumshoe.Finding {
	found := []gumshoe.Finding{}
	// TODO
	return found
}

// GetTypeName impl.
func (f *FacebookUser) GetTypeName() string {
	return FacebookType
}

// GetID impl.
func (f *FacebookUser) GetID() string {
	return f.Name
}

// Display a Facebook user
func (f *FacebookUser) Display() {
	fmt.Printf("[facebook username] %s\n", f.Name)
}
